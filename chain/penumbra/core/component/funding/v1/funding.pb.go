// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: penumbra/core/component/funding/v1/funding.proto

package fundingv1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Funding component configuration data.
type FundingParameters struct {
}

func (m *FundingParameters) Reset()         { *m = FundingParameters{} }
func (m *FundingParameters) String() string { return proto.CompactTextString(m) }
func (*FundingParameters) ProtoMessage()    {}
func (*FundingParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_586f7399b4e8a469, []int{0}
}
func (m *FundingParameters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundingParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundingParameters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundingParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingParameters.Merge(m, src)
}
func (m *FundingParameters) XXX_Size() int {
	return m.Size()
}
func (m *FundingParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingParameters.DiscardUnknown(m)
}

var xxx_messageInfo_FundingParameters proto.InternalMessageInfo

// Genesis data for the funding component.
type GenesisContent struct {
	FundingParams *FundingParameters `protobuf:"bytes,1,opt,name=funding_params,json=fundingParams,proto3" json:"funding_params,omitempty"`
}

func (m *GenesisContent) Reset()         { *m = GenesisContent{} }
func (m *GenesisContent) String() string { return proto.CompactTextString(m) }
func (*GenesisContent) ProtoMessage()    {}
func (*GenesisContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_586f7399b4e8a469, []int{1}
}
func (m *GenesisContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisContent.Merge(m, src)
}
func (m *GenesisContent) XXX_Size() int {
	return m.Size()
}
func (m *GenesisContent) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisContent.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisContent proto.InternalMessageInfo

func (m *GenesisContent) GetFundingParams() *FundingParameters {
	if m != nil {
		return m.FundingParams
	}
	return nil
}

func init() {
	proto.RegisterType((*FundingParameters)(nil), "penumbra.core.component.funding.v1.FundingParameters")
	proto.RegisterType((*GenesisContent)(nil), "penumbra.core.component.funding.v1.GenesisContent")
}

func init() {
	proto.RegisterFile("penumbra/core/component/funding/v1/funding.proto", fileDescriptor_586f7399b4e8a469)
}

var fileDescriptor_586f7399b4e8a469 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x9b, 0xf0, 0xf2, 0x0e, 0x51, 0x0b, 0xd6, 0xc5, 0x29, 0x48, 0x86, 0xe2, 0xe2, 0x9d,
	0x51, 0x04, 0x89, 0x5b, 0x03, 0xed, 0x24, 0x04, 0x87, 0x0e, 0x12, 0x90, 0x6b, 0xfa, 0xb4, 0x0d,
	0x98, 0xe7, 0xc2, 0xdd, 0x93, 0x7c, 0x0e, 0x3f, 0x83, 0xa3, 0x9f, 0x44, 0x9c, 0x3a, 0x3a, 0x6a,
	0xba, 0xf9, 0x29, 0xe4, 0x6a, 0xae, 0x0a, 0x0e, 0x75, 0x49, 0x9e, 0x27, 0xfc, 0x7e, 0xfc, 0xff,
	0x39, 0xce, 0x3b, 0x2d, 0x01, 0xab, 0x62, 0xa2, 0x04, 0xcf, 0xa4, 0x02, 0x9e, 0xc9, 0xa2, 0x94,
	0x08, 0x48, 0x7c, 0x56, 0xe1, 0x34, 0xc7, 0x39, 0xaf, 0x43, 0x3b, 0xb2, 0x52, 0x49, 0x92, 0xbd,
	0xc0, 0x1a, 0xcc, 0x18, 0x6c, 0x63, 0x30, 0x8b, 0xd5, 0x61, 0x70, 0xe0, 0xed, 0x0f, 0xbf, 0xb6,
	0x44, 0x28, 0x51, 0x00, 0x81, 0xd2, 0x01, 0x7a, 0xdd, 0x11, 0x20, 0xe8, 0x5c, 0xc7, 0x12, 0x09,
	0x90, 0x7a, 0xa9, 0xd7, 0x6d, 0xa5, 0xbb, 0xd2, 0x70, 0xfa, 0xd0, 0x39, 0x72, 0x8e, 0x77, 0xce,
	0x2e, 0xd8, 0xf6, 0x0c, 0xf6, 0x2b, 0xe0, 0x66, 0x6f, 0xf6, 0xe3, 0x93, 0x1e, 0xbc, 0xbb, 0xcf,
	0x8d, 0xef, 0x2c, 0x1b, 0xdf, 0x79, 0x6b, 0x7c, 0xe7, 0x61, 0xe5, 0x77, 0x96, 0x2b, 0xbf, 0xf3,
	0xba, 0xf2, 0x3b, 0x5e, 0x3f, 0x93, 0xc5, 0x1f, 0x32, 0x06, 0xbb, 0x36, 0xc4, 0xfc, 0x79, 0xe2,
	0xdc, 0x4e, 0xe7, 0x39, 0x2d, 0xaa, 0x89, 0x81, 0xb9, 0x26, 0x25, 0x70, 0x0e, 0xf7, 0xb2, 0x86,
	0x93, 0x1a, 0x90, 0x2a, 0x05, 0x9a, 0xe7, 0x48, 0xa0, 0xb2, 0x85, 0x30, 0x6f, 0x4d, 0xbc, 0xbe,
	0xe4, 0xeb, 0x85, 0x6f, 0x3f, 0xe5, 0xab, 0x76, 0xac, 0xc3, 0x47, 0xf7, 0x5f, 0x12, 0xc7, 0xc3,
	0x27, 0x37, 0x48, 0x6c, 0xc3, 0xd8, 0x34, 0x8c, 0x37, 0x0d, 0xdb, 0x56, 0x6c, 0x1c, 0xbe, 0x7c,
	0x43, 0xa9, 0x81, 0xd2, 0x0d, 0x94, 0xb6, 0x50, 0x3a, 0x0e, 0x1b, 0x97, 0x6d, 0x87, 0xd2, 0x51,
	0x32, 0xb8, 0x06, 0x12, 0x53, 0x41, 0xe2, 0xc3, 0xed, 0x5b, 0x21, 0x8a, 0x8c, 0x61, 0x9e, 0xad,
	0x12, 0x45, 0xad, 0x13, 0x45, 0xe3, 0x70, 0xf2, 0x7f, 0x7d, 0x27, 0xce, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x09, 0x80, 0x86, 0x4d, 0x47, 0x02, 0x00, 0x00,
}

func (m *FundingParameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundingParameters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FundingParameters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GenesisContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FundingParams != nil {
		{
			size, err := m.FundingParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFunding(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFunding(dAtA []byte, offset int, v uint64) int {
	offset -= sovFunding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FundingParameters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GenesisContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FundingParams != nil {
		l = m.FundingParams.Size()
		n += 1 + l + sovFunding(uint64(l))
	}
	return n
}

func sovFunding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFunding(x uint64) (n int) {
	return sovFunding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FundingParameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunding
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
			return fmt.Errorf("proto: FundingParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundingParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFunding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunding
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
func (m *GenesisContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunding
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
			return fmt.Errorf("proto: GenesisContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunding
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
				return ErrInvalidLengthFunding
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFunding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FundingParams == nil {
				m.FundingParams = &FundingParameters{}
			}
			if err := m.FundingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFunding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFunding
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
func skipFunding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFunding
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
					return 0, ErrIntOverflowFunding
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
					return 0, ErrIntOverflowFunding
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
				return 0, ErrInvalidLengthFunding
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFunding
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFunding
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFunding        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFunding          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFunding = fmt.Errorf("proto: unexpected end of group")
)