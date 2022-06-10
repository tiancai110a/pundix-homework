package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// GetDeposit gets the deposit of a specific depositor on a specific proposal
func (keeper Keeper) GetDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) (deposit types.Deposit, found bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.DepositKey(proposalID, depositorAddr))
	if bz == nil {
		return deposit, false
	}

	keeper.cdc.MustUnmarshal(bz, &deposit)

	return deposit, true
}

// SetDeposit sets a Deposit to the gov store
func (keeper Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshal(&deposit)
	depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
	if err != nil {
		panic(err)
	}

	store.Set(types.DepositKey(deposit.ProposalId, depositor), bz)
}

// GetAllDeposits returns all the deposits from the store
func (keeper Keeper) GetAllDeposits(ctx sdk.Context) (deposits types.Deposits) {
	keeper.IterateAllDeposits(ctx, func(deposit types.Deposit) bool {
		deposits = append(deposits, deposit)
		return false
	})

	return
}

// GetDeposits returns all the deposits from a proposal
func (keeper Keeper) GetDeposits(ctx sdk.Context, proposalID uint64) (deposits types.Deposits) {
	keeper.IterateDeposits(ctx, proposalID, func(deposit types.Deposit) bool {
		deposits = append(deposits, deposit)
		return false
	})

	return
}

// DeleteDeposits deletes all the deposits on a specific proposal without refunding them
func (keeper Keeper) DeleteDeposits(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)

	keeper.IterateDeposits(ctx, proposalID, func(deposit types.Deposit) bool {
		err := keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, deposit.Amount)
		if err != nil {
			panic(err)
		}

		depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
		if err != nil {
			panic(err)
		}
		store.Delete(types.DepositKey(proposalID, depositor))
		return false
	})
}

// IterateAllDeposits iterates over the all the stored deposits and performs a callback function
func (keeper Keeper) IterateAllDeposits(ctx sdk.Context, cb func(deposit types.Deposit) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositsKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deposit types.Deposit

		keeper.cdc.MustUnmarshal(iterator.Value(), &deposit)

		if cb(deposit) {
			break
		}
	}
}

// IterateDeposits iterates over the all the proposals deposits and performs a callback function
func (keeper Keeper) IterateDeposits(ctx sdk.Context, proposalID uint64, cb func(deposit types.Deposit) (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositsKey(proposalID))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deposit types.Deposit

		keeper.cdc.MustUnmarshal(iterator.Value(), &deposit)

		if cb(deposit) {
			break
		}
	}
}

// AddDeposit adds or updates a deposit of a specific depositor on a specific proposal
// Activates voting period when appropriate
func (keeper Keeper) AddDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress, depositAmount sdk.Coins) (bool, error) {
	// Checks to see if proposal exists
	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		return false, sdkerrors.Wrapf(types.ErrUnknownProposal, "%d", proposalID)
	}

	// Check if proposal is still depositable
	if (proposal.Status != types.StatusDepositPeriod) && (proposal.Status != types.StatusVotingPeriod) {
		return false, sdkerrors.Wrapf(types.ErrInactiveProposal, "%d", proposalID)
	}

	// update the governance module's account coins pool
	err := keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, depositorAddr, types.ModuleName, depositAmount)
	if err != nil {
		return false, err
	}

	// first deposit
	first := proposal.TotalDeposit.IsZero()
	// Update proposal
	proposal.TotalDeposit = proposal.TotalDeposit.Add(depositAmount...)
	keeper.SetProposal(ctx, proposal)

	// Check if deposit has provided sufficient total funds to transition the proposal into the voting period
	activatedVotingPeriod := false

	var minDeposit sdk.Coins

	if keeper.SupportEGFProposal(ctx) && (types.CommunityPoolSpendByRouter == proposal.ProposalRoute() &&
		types.ProposalTypeCommunityPoolSpend == proposal.ProposalType()) {
		cpsp, ok := proposal.GetContent().(*types2.CommunityPoolSpendProposal)
		if !ok {
			return false, sdkerrors.Wrapf(types.ErrInvalidProposalType, "%d", proposalID)
		}
		minDeposit = keeper.SupportEGFProposalTotalDeposit(ctx, first, cpsp.Amount)
	} else {
		minDeposit = keeper.GetDepositParams(ctx).MinDeposit
	}
	if proposal.Status == types.StatusDepositPeriod && proposal.TotalDeposit.IsAllGTE(minDeposit) {
		keeper.ActivateVotingPeriod(ctx, proposal)
		activatedVotingPeriod = true
	}
	// Add or update deposit object
	deposit, found := keeper.GetDeposit(ctx, proposalID, depositorAddr)

	if found {
		deposit.Amount = deposit.Amount.Add(depositAmount...)
	} else {
		deposit = types.NewDeposit(proposalID, depositorAddr, depositAmount)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProposalDeposit,
			sdk.NewAttribute(sdk.AttributeKeyAmount, depositAmount.String()),
			sdk.NewAttribute(types.AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
		),
	)

	keeper.SetDeposit(ctx, deposit)

	return activatedVotingPeriod, nil
}

// RefundDeposits refunds and deletes all the deposits on a specific proposal
func (keeper Keeper) RefundDeposits(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)

	keeper.IterateDeposits(ctx, proposalID, func(deposit types.Deposit) bool {
		depositor, err := sdk.AccAddressFromBech32(deposit.Depositor)
		if err != nil {
			panic(err)
		}

		err = keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositor, deposit.Amount)
		if err != nil {
			panic(err)
		}

		store.Delete(types.DepositKey(proposalID, depositor))
		return false
	})
}

func (keeper Keeper) SupportEGFProposal(ctx sdk.Context) bool {
	// to reach the specified height
	if types.GetEGFProposalSupportBlock() == 0 || ctx.BlockHeight() < types.GetEGFProposalSupportBlock() {
		return false
	}
	if !keeper.hasInitEGFDepositParams(ctx) {
		// initialization egf parameters
		keeper.SetEGFDepositParams(ctx, types.DefaultEGFDepositParams())
	}
	return true
}

func (keeper Keeper) SupportEGFProposalTotalDeposit(ctx sdk.Context, first bool, claimCoin sdk.Coins) sdk.Coins {
	// minimum collateral amount for initializing EGF proposals
	egfDepositParams := keeper.GetEGFDepositParams(ctx)
	if claimCoin.IsAllLTE(egfDepositParams.DepositProposalThreshold) && first {
		return egfDepositParams.InitialDeposit
	}
	initialDeposit := egfDepositParams.InitialDeposit
	for _, coin := range claimCoin {
		amount := coin.Amount.ToDec().Mul(egfDepositParams.ClaimRatio).TruncateInt()
		initialDeposit = initialDeposit.Add(sdk.NewCoin(coin.Denom, amount))
	}
	return initialDeposit
}