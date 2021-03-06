// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: other/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GasPriceRequest struct {
}

func (m *GasPriceRequest) Reset()         { *m = GasPriceRequest{} }
func (m *GasPriceRequest) String() string { return proto.CompactTextString(m) }
func (*GasPriceRequest) ProtoMessage()    {}
func (*GasPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2add8ac0113356f, []int{0}
}
func (m *GasPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPriceRequest.Merge(m, src)
}
func (m *GasPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *GasPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GasPriceRequest proto.InternalMessageInfo

type GasPriceResponse struct {
	GasPrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=gas_prices,json=gasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"gas_prices" yaml:"gas_prices"`
}

func (m *GasPriceResponse) Reset()         { *m = GasPriceResponse{} }
func (m *GasPriceResponse) String() string { return proto.CompactTextString(m) }
func (*GasPriceResponse) ProtoMessage()    {}
func (*GasPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2add8ac0113356f, []int{1}
}
func (m *GasPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPriceResponse.Merge(m, src)
}
func (m *GasPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *GasPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GasPriceResponse proto.InternalMessageInfo

func (m *GasPriceResponse) GetGasPrices() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.GasPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*GasPriceRequest)(nil), "fx.other.GasPriceRequest")
	proto.RegisterType((*GasPriceResponse)(nil), "fx.other.GasPriceResponse")
}

func init() { proto.RegisterFile("other/query.proto", fileDescriptor_b2add8ac0113356f) }

var fileDescriptor_b2add8ac0113356f = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x93, 0x7b, 0xb9, 0x97, 0x1a, 0x17, 0xda, 0xa8, 0x60, 0xa3, 0x4c, 0x25, 0xab, 0x22,
	0x74, 0x86, 0xd6, 0x9d, 0xcb, 0xaa, 0xb8, 0xd5, 0x2e, 0x5d, 0x28, 0x93, 0x38, 0x9d, 0x06, 0xdb,
	0x39, 0x69, 0xce, 0xa4, 0xa4, 0x2b, 0xc1, 0x27, 0x10, 0xfa, 0x16, 0x3e, 0x49, 0x97, 0x05, 0x37,
	0xae, 0xaa, 0xb4, 0x3e, 0x81, 0x4f, 0x20, 0xf9, 0x53, 0x2b, 0xe2, 0x2a, 0xe1, 0xfb, 0x98, 0x8f,
	0xdf, 0xf9, 0x59, 0x65, 0xd0, 0x5d, 0x11, 0xb1, 0x41, 0x2c, 0xa2, 0x11, 0x0d, 0x23, 0xd0, 0x60,
	0x97, 0x3a, 0x09, 0xcd, 0x52, 0x67, 0x5b, 0x82, 0x84, 0x2c, 0x64, 0xe9, 0x5f, 0xde, 0x3b, 0xfb,
	0x12, 0x40, 0xf6, 0x04, 0xe3, 0x61, 0xc0, 0xb8, 0x52, 0xa0, 0xb9, 0x0e, 0x40, 0x61, 0xd1, 0x12,
	0x1f, 0xb0, 0x0f, 0xc8, 0x3c, 0x8e, 0x82, 0x0d, 0x1b, 0x9e, 0xd0, 0xbc, 0xc1, 0x7c, 0x08, 0x54,
	0xde, 0xbb, 0x65, 0x6b, 0xe3, 0x9c, 0xe3, 0x45, 0x14, 0xf8, 0xa2, 0x2d, 0x06, 0xb1, 0x40, 0xed,
	0x8e, 0x4d, 0x6b, 0x73, 0x95, 0x61, 0x08, 0x0a, 0x85, 0x7d, 0x6f, 0x59, 0x92, 0xe3, 0x4d, 0x98,
	0x86, 0xb8, 0x6b, 0x1e, 0xfc, 0xad, 0xad, 0x37, 0x2b, 0x34, 0x1f, 0xa7, 0xe9, 0x38, 0x2d, 0xc6,
	0xe9, 0x09, 0x04, 0xaa, 0x75, 0x36, 0x99, 0x55, 0x8d, 0x8f, 0x59, 0xb5, 0x3c, 0xe2, 0xfd, 0xde,
	0xb1, 0xbb, 0x7a, 0xea, 0x3e, 0xbd, 0x56, 0x6b, 0x32, 0xd0, 0xdd, 0xd8, 0xa3, 0x3e, 0xf4, 0x59,
	0x81, 0x97, 0x7f, 0xea, 0x78, 0x7b, 0xc7, 0xf4, 0x28, 0x14, 0x98, 0xad, 0x60, 0x7b, 0x4d, 0x16,
	0x1c, 0xd8, 0x94, 0xd6, 0xbf, 0xcb, 0xd4, 0x8a, 0x7d, 0x6d, 0x95, 0x96, 0x74, 0x76, 0x85, 0x2e,
	0xe5, 0xd0, 0x1f, 0x57, 0x38, 0xce, 0x6f, 0x55, 0x7e, 0x8c, 0xbb, 0xf7, 0xf0, 0xfc, 0x3e, 0xfe,
	0xb3, 0x63, 0x6f, 0xb1, 0x5c, 0xf7, 0xb0, 0xc1, 0xbe, 0x08, 0x5b, 0xa7, 0x93, 0x39, 0x31, 0xa7,
	0x73, 0x62, 0xbe, 0xcd, 0x89, 0xf9, 0xb8, 0x20, 0xc6, 0x74, 0x41, 0x8c, 0x97, 0x05, 0x31, 0xae,
	0x0e, 0xbf, 0x71, 0x77, 0x62, 0xe5, 0xa7, 0x9e, 0x13, 0xd6, 0x49, 0xea, 0x3e, 0x44, 0x82, 0x25,
	0xc5, 0x58, 0xc6, 0xef, 0xfd, 0xcf, 0xf4, 0x1e, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x45,
	0xf8, 0xdb, 0xd1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	GasPrice(ctx context.Context, in *GasPriceRequest, opts ...grpc.CallOption) (*GasPriceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GasPrice(ctx context.Context, in *GasPriceRequest, opts ...grpc.CallOption) (*GasPriceResponse, error) {
	out := new(GasPriceResponse)
	err := c.cc.Invoke(ctx, "/fx.other.Query/GasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GasPrice(context.Context, *GasPriceRequest) (*GasPriceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GasPrice(ctx context.Context, req *GasPriceRequest) (*GasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GasPrice not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GasPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.other.Query/GasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GasPrice(ctx, req.(*GasPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fx.other.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GasPrice",
			Handler:    _Query_GasPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "other/query.proto",
}

func (m *GasPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GasPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasPrices) > 0 {
		for iNdEx := len(m.GasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GasPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GasPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GasPrices) > 0 {
		for _, e := range m.GasPrices {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GasPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GasPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GasPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GasPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrices = append(m.GasPrices, types.Coin{})
			if err := m.GasPrices[len(m.GasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
