// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ambinet/inflation/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the inflation module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// amount of past periods, based on the epochs per period param
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// inflation epoch identifier
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// number of epochs after which inflation is recalculated
	EpochsPerPeriod int64 `protobuf:"varint,4,opt,name=epochs_per_period,json=epochsPerPeriod,proto3" json:"epochs_per_period,omitempty"`
	// number of epochs that have passed while inflation is disabled
	SkippedEpochs uint64 `protobuf:"varint,5,opt,name=skipped_epochs,json=skippedEpochs,proto3" json:"skipped_epochs,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb8eee530db1235, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *GenesisState) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *GenesisState) GetEpochsPerPeriod() int64 {
	if m != nil {
		return m.EpochsPerPeriod
	}
	return 0
}

func (m *GenesisState) GetSkippedEpochs() uint64 {
	if m != nil {
		return m.SkippedEpochs
	}
	return 0
}

// Params holds parameters for the inflation module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// variables to calculate exponential inflation
	ExponentialCalculation ExponentialCalculation `protobuf:"bytes,2,opt,name=exponential_calculation,json=exponentialCalculation,proto3" json:"exponential_calculation"`
	// inflation distribution of the minted denom
	InflationDistribution InflationDistribution `protobuf:"bytes,3,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
	// parameter to enable inflation and halt increasing the skipped_epochs
	EnableInflation bool `protobuf:"varint,4,opt,name=enable_inflation,json=enableInflation,proto3" json:"enable_inflation,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb8eee530db1235, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetExponentialCalculation() ExponentialCalculation {
	if m != nil {
		return m.ExponentialCalculation
	}
	return ExponentialCalculation{}
}

func (m *Params) GetInflationDistribution() InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return InflationDistribution{}
}

func (m *Params) GetEnableInflation() bool {
	if m != nil {
		return m.EnableInflation
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ambinet.inflation.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "ambinet.inflation.v1.Params")
}

func init() { proto.RegisterFile("ambinet/inflation/v1/genesis.proto", fileDescriptor_1cb8eee530db1235) }

var fileDescriptor_1cb8eee530db1235 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xd7, 0x52, 0x51, 0x0f, 0x18, 0x58, 0x50, 0xaa, 0x4a, 0x84, 0xa8, 0x12, 0x52, 0x37,
	0x89, 0x44, 0xdb, 0x2e, 0x9c, 0xc7, 0xa6, 0x69, 0xb7, 0x2a, 0xdc, 0xb8, 0x58, 0x6e, 0xf2, 0x9a,
	0x5a, 0x24, 0xb6, 0x65, 0xbb, 0xd1, 0xf8, 0x17, 0xfc, 0x09, 0xfe, 0xcb, 0x8e, 0x3b, 0x72, 0x9a,
	0x50, 0xfb, 0x47, 0x50, 0xec, 0x90, 0x4e, 0x5a, 0x6e, 0xf6, 0xf7, 0x7d, 0xef, 0xfb, 0xde, 0x7b,
	0x7a, 0x38, 0x84, 0xaa, 0x94, 0x26, 0xe6, 0x62, 0x55, 0x30, 0xcb, 0xa5, 0x88, 0xab, 0xd3, 0x38,
	0x07, 0x01, 0x86, 0x9b, 0x48, 0x69, 0x69, 0x25, 0x21, 0x4e, 0x11, 0xb5, 0x8a, 0xa8, 0x3a, 0x9d,
	0xbe, 0xcd, 0x65, 0x2e, 0x1d, 0x1d, 0xd7, 0x2f, 0xaf, 0x9c, 0xce, 0x3a, 0xbc, 0xf6, 0x65, 0x4e,
	0x33, 0x7b, 0x40, 0xf8, 0xc5, 0xb5, 0xf7, 0xff, 0x66, 0x99, 0x05, 0xf2, 0x05, 0x0f, 0x15, 0xd3,
	0xac, 0x34, 0x13, 0x14, 0xa2, 0xf9, 0xe1, 0xd9, 0x34, 0x7a, 0x9a, 0x17, 0x2d, 0x9c, 0xe2, 0x62,
	0x70, 0xf7, 0xf0, 0xb1, 0x97, 0x34, 0x7a, 0x32, 0xc6, 0x43, 0x05, 0x9a, 0xcb, 0x6c, 0x72, 0x10,
	0xa2, 0xf9, 0x20, 0x69, 0x7e, 0xe4, 0x18, 0xbf, 0x06, 0x25, 0xd3, 0x35, 0xe5, 0x19, 0x08, 0xcb,
	0x57, 0x1c, 0xf4, 0xa4, 0x1f, 0xa2, 0xf9, 0x28, 0x39, 0x72, 0xf8, 0x4d, 0x0b, 0x93, 0x13, 0xfc,
	0xc6, 0x41, 0x86, 0x2a, 0xd0, 0xb4, 0x71, 0x1b, 0x84, 0x68, 0xde, 0x6f, 0xb4, 0x66, 0x01, 0x7a,
	0xe1, 0x6d, 0x3f, 0xe1, 0x57, 0xe6, 0x07, 0x57, 0x0a, 0x32, 0xea, 0xa9, 0xc9, 0x33, 0x17, 0xfb,
	0xb2, 0x41, 0xaf, 0x1c, 0x38, 0xfb, 0x7d, 0x80, 0x87, 0xbe, 0x5d, 0xf2, 0x01, 0xe3, 0x92, 0x0b,
	0x4b, 0x33, 0x10, 0xb2, 0x74, 0xe3, 0x8d, 0x92, 0x51, 0x8d, 0x5c, 0xd6, 0x00, 0xe1, 0xf8, 0x3d,
	0xdc, 0x2a, 0x29, 0xea, 0x6e, 0x58, 0x41, 0x53, 0x56, 0xa4, 0x1b, 0x3f, 0xb2, 0x1b, 0xe8, 0xf0,
	0xec, 0xa4, 0x6b, 0x15, 0x57, 0xfb, 0x92, 0xaf, 0xfb, 0x8a, 0x66, 0x35, 0x63, 0xe8, 0x64, 0xc9,
	0x0a, 0x8f, 0x5b, 0x13, 0x9a, 0x71, 0x63, 0x35, 0x5f, 0x6e, 0x5c, 0x52, 0xdf, 0x25, 0x1d, 0x77,
	0x25, 0xdd, 0xfc, 0xff, 0x5c, 0x3e, 0x2a, 0x68, 0x82, 0xde, 0xf1, 0x2e, 0xd2, 0xad, 0x5e, 0xb0,
	0x65, 0x01, 0xb4, 0xe5, 0xdd, 0x3a, 0x9f, 0x27, 0x47, 0x1e, 0x6f, 0x3d, 0x2f, 0xae, 0xef, 0xb6,
	0x01, 0xba, 0xdf, 0x06, 0xe8, 0xef, 0x36, 0x40, 0xbf, 0x76, 0x41, 0xef, 0x7e, 0x17, 0xf4, 0xfe,
	0xec, 0x82, 0xde, 0xf7, 0xcf, 0x39, 0xb7, 0xeb, 0xcd, 0x32, 0x4a, 0x65, 0x19, 0xdb, 0x35, 0xd3,
	0x86, 0x9b, 0xd8, 0x5f, 0x56, 0x75, 0x1e, 0xdf, 0x3e, 0x3a, 0x2f, 0xfb, 0x53, 0x81, 0x59, 0x0e,
	0xdd, 0x61, 0x9d, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x3c, 0x1f, 0xb7, 0xca, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SkippedEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SkippedEpochs))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochsPerPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochsPerPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EnableInflation {
		i--
		if m.EnableInflation {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.InflationDistribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ExponentialCalculation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.EpochsPerPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.EpochsPerPeriod))
	}
	if m.SkippedEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.SkippedEpochs))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ExponentialCalculation.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EnableInflation {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochsPerPeriod", wireType)
			}
			m.EpochsPerPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochsPerPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkippedEpochs", wireType)
			}
			m.SkippedEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkippedEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentialCalculation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExponentialCalculation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationDistribution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableInflation", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableInflation = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
