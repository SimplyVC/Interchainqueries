// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icq/pending_icq.proto

package types

import (
	fmt "fmt"
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

//*
//PendingICQInstance is one instance of a request query. This is done to remove
//duplicate fields, we keep a pointer of the periodic query in this instance
//so we can trace back to all the requested fields.
type PendingICQInstance struct {
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TimeoutHeight uint64 `protobuf:"varint,2,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	TargetHeight  uint64 `protobuf:"varint,3,opt,name=target_height,json=targetHeight,proto3" json:"target_height,omitempty"`
	PeriodicId    uint64 `protobuf:"varint,4,opt,name=periodic_id,json=periodicId,proto3" json:"periodic_id,omitempty"`
}

func (m *PendingICQInstance) Reset()         { *m = PendingICQInstance{} }
func (m *PendingICQInstance) String() string { return proto.CompactTextString(m) }
func (*PendingICQInstance) ProtoMessage()    {}
func (*PendingICQInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb23dfe41bec70fc, []int{0}
}
func (m *PendingICQInstance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingICQInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingICQInstance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingICQInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingICQInstance.Merge(m, src)
}
func (m *PendingICQInstance) XXX_Size() int {
	return m.Size()
}
func (m *PendingICQInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingICQInstance.DiscardUnknown(m)
}

var xxx_messageInfo_PendingICQInstance proto.InternalMessageInfo

func (m *PendingICQInstance) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PendingICQInstance) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *PendingICQInstance) GetTargetHeight() uint64 {
	if m != nil {
		return m.TargetHeight
	}
	return 0
}

func (m *PendingICQInstance) GetPeriodicId() uint64 {
	if m != nil {
		return m.PeriodicId
	}
	return 0
}

//*
//PendingICQRequest is the full request we combine Periodic Query Data as well
//as the PendingICQ data, so that relayers can process the requests.
type PendingICQRequest struct {
	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path            string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	TimeoutHeight   uint64 `protobuf:"varint,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	TargetHeight    uint64 `protobuf:"varint,4,opt,name=target_height,json=targetHeight,proto3" json:"target_height,omitempty"`
	ClientId        string `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Creator         string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainId         string `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	QueryParameters []byte `protobuf:"bytes,8,opt,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	PeriodicId      uint64 `protobuf:"varint,9,opt,name=periodic_id,json=periodicId,proto3" json:"periodic_id,omitempty"`
}

func (m *PendingICQRequest) Reset()         { *m = PendingICQRequest{} }
func (m *PendingICQRequest) String() string { return proto.CompactTextString(m) }
func (*PendingICQRequest) ProtoMessage()    {}
func (*PendingICQRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb23dfe41bec70fc, []int{1}
}
func (m *PendingICQRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingICQRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingICQRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingICQRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingICQRequest.Merge(m, src)
}
func (m *PendingICQRequest) XXX_Size() int {
	return m.Size()
}
func (m *PendingICQRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingICQRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PendingICQRequest proto.InternalMessageInfo

func (m *PendingICQRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PendingICQRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PendingICQRequest) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *PendingICQRequest) GetTargetHeight() uint64 {
	if m != nil {
		return m.TargetHeight
	}
	return 0
}

func (m *PendingICQRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PendingICQRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PendingICQRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *PendingICQRequest) GetQueryParameters() []byte {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *PendingICQRequest) GetPeriodicId() uint64 {
	if m != nil {
		return m.PeriodicId
	}
	return 0
}

func init() {
	proto.RegisterType((*PendingICQInstance)(nil), "simplyvc.interchainqueries.icq.PendingICQInstance")
	proto.RegisterType((*PendingICQRequest)(nil), "simplyvc.interchainqueries.icq.PendingICQRequest")
}

func init() { proto.RegisterFile("icq/pending_icq.proto", fileDescriptor_fb23dfe41bec70fc) }

var fileDescriptor_fb23dfe41bec70fc = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x9b, 0x34, 0xb7, 0x7f, 0xe6, 0xb6, 0xbd, 0xf7, 0x0e, 0x5c, 0x88, 0x08, 0xb1, 0x54,
	0x84, 0xba, 0x49, 0x10, 0xdf, 0x40, 0x37, 0x06, 0x37, 0x35, 0x4b, 0x37, 0x21, 0x9d, 0x39, 0x24,
	0x07, 0x9a, 0x64, 0x32, 0x99, 0x88, 0x7d, 0x0b, 0x7d, 0x04, 0xdf, 0xc6, 0x65, 0x97, 0x2e, 0xa5,
	0x7d, 0x11, 0xc9, 0xf4, 0x8f, 0xd0, 0x76, 0xe1, 0x6e, 0xce, 0x77, 0xbe, 0xc3, 0xfc, 0xf8, 0xce,
	0x21, 0xff, 0x91, 0x15, 0x9e, 0x80, 0x8c, 0x63, 0x16, 0x87, 0xc8, 0x0a, 0x57, 0xc8, 0x5c, 0xe5,
	0xd4, 0x29, 0x31, 0x15, 0xb3, 0xf9, 0x13, 0x73, 0x31, 0x53, 0x20, 0x59, 0x12, 0x61, 0x56, 0x54,
	0x20, 0x11, 0x4a, 0x17, 0x59, 0x31, 0x7a, 0x35, 0x08, 0x9d, 0xac, 0xa7, 0xfc, 0xdb, 0x07, 0x3f,
	0x2b, 0x55, 0x94, 0x31, 0xa0, 0x03, 0x62, 0x22, 0xb7, 0x8d, 0xa1, 0x31, 0xb6, 0x02, 0x13, 0x39,
	0xbd, 0x20, 0x03, 0x85, 0x29, 0xe4, 0x95, 0x0a, 0x13, 0xc0, 0x38, 0x51, 0xb6, 0xa9, 0x7b, 0xfd,
	0x8d, 0x7a, 0xa7, 0x45, 0x7a, 0x4e, 0xfa, 0x2a, 0x92, 0x31, 0xec, 0x5c, 0x4d, 0xed, 0xea, 0xad,
	0xc5, 0x8d, 0xe9, 0x8c, 0xfc, 0x16, 0x20, 0x31, 0xe7, 0xc8, 0x42, 0xe4, 0xb6, 0xa5, 0x2d, 0x64,
	0x2b, 0xf9, 0x7c, 0xf4, 0x66, 0x92, 0x7f, 0xdf, 0x4c, 0x01, 0x14, 0x15, 0x94, 0xea, 0x00, 0x89,
	0x12, 0x4b, 0x44, 0x2a, 0xd1, 0x20, 0xdd, 0x40, 0xbf, 0x8f, 0x60, 0x36, 0x7f, 0x84, 0x69, 0x1d,
	0xc1, 0x3c, 0x25, 0x5d, 0x36, 0x43, 0xc8, 0x54, 0x0d, 0xf9, 0x4b, 0x7f, 0xd2, 0x59, 0x0b, 0x3e,
	0xa7, 0x36, 0x69, 0x33, 0x09, 0x91, 0xca, 0xa5, 0xdd, 0xd2, 0xad, 0x6d, 0x49, 0x4f, 0x48, 0x47,
	0x87, 0x5c, 0x4f, 0xb5, 0x37, 0xad, 0xba, 0xf6, 0x39, 0xbd, 0x24, 0x7f, 0xeb, 0xe8, 0xe7, 0xa1,
	0x88, 0x64, 0x94, 0x82, 0x02, 0x59, 0xda, 0x9d, 0xa1, 0x31, 0xee, 0x05, 0x7f, 0xb4, 0x3e, 0xd9,
	0xc9, 0xfb, 0x19, 0x75, 0xf7, 0x33, 0xba, 0xb9, 0x7f, 0x5f, 0x3a, 0xc6, 0x62, 0xe9, 0x18, 0x9f,
	0x4b, 0xc7, 0x78, 0x59, 0x39, 0x8d, 0xc5, 0xca, 0x69, 0x7c, 0xac, 0x9c, 0xc6, 0xe3, 0x55, 0x8c,
	0x2a, 0xa9, 0xa6, 0x2e, 0xcb, 0x53, 0x6f, 0xbb, 0x7c, 0xef, 0x60, 0xf9, 0xde, 0xb3, 0x57, 0x1f,
	0x8c, 0x9a, 0x0b, 0x28, 0xa7, 0x2d, 0x7d, 0x2b, 0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0x1a, 0xc3, 0x2b, 0x44, 0x02, 0x00, 0x00,
}

func (m *PendingICQInstance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingICQInstance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingICQInstance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PeriodicId != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.PeriodicId))
		i--
		dAtA[i] = 0x20
	}
	if m.TargetHeight != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.TargetHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PendingICQRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingICQRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingICQRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PeriodicId != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.PeriodicId))
		i--
		dAtA[i] = 0x48
	}
	if len(m.QueryParameters) > 0 {
		i -= len(m.QueryParameters)
		copy(dAtA[i:], m.QueryParameters)
		i = encodeVarintPendingIcq(dAtA, i, uint64(len(m.QueryParameters)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintPendingIcq(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPendingIcq(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintPendingIcq(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TargetHeight != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.TargetHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintPendingIcq(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPendingIcq(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPendingIcq(dAtA []byte, offset int, v uint64) int {
	offset -= sovPendingIcq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PendingICQInstance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPendingIcq(uint64(m.Id))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovPendingIcq(uint64(m.TimeoutHeight))
	}
	if m.TargetHeight != 0 {
		n += 1 + sovPendingIcq(uint64(m.TargetHeight))
	}
	if m.PeriodicId != 0 {
		n += 1 + sovPendingIcq(uint64(m.PeriodicId))
	}
	return n
}

func (m *PendingICQRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPendingIcq(uint64(m.Id))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovPendingIcq(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovPendingIcq(uint64(m.TimeoutHeight))
	}
	if m.TargetHeight != 0 {
		n += 1 + sovPendingIcq(uint64(m.TargetHeight))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovPendingIcq(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPendingIcq(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovPendingIcq(uint64(l))
	}
	l = len(m.QueryParameters)
	if l > 0 {
		n += 1 + l + sovPendingIcq(uint64(l))
	}
	if m.PeriodicId != 0 {
		n += 1 + sovPendingIcq(uint64(m.PeriodicId))
	}
	return n
}

func sovPendingIcq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPendingIcq(x uint64) (n int) {
	return sovPendingIcq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PendingICQInstance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPendingIcq
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
			return fmt.Errorf("proto: PendingICQInstance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingICQInstance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetHeight", wireType)
			}
			m.TargetHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicId", wireType)
			}
			m.PeriodicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeriodicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPendingIcq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPendingIcq
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
func (m *PendingICQRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPendingIcq
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
			return fmt.Errorf("proto: PendingICQRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingICQRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
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
				return ErrInvalidLengthPendingIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetHeight", wireType)
			}
			m.TargetHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
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
				return ErrInvalidLengthPendingIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
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
				return ErrInvalidLengthPendingIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
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
				return ErrInvalidLengthPendingIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryParameters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPendingIcq
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryParameters = append(m.QueryParameters[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryParameters == nil {
				m.QueryParameters = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicId", wireType)
			}
			m.PeriodicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeriodicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPendingIcq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPendingIcq
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
func skipPendingIcq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPendingIcq
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
					return 0, ErrIntOverflowPendingIcq
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
					return 0, ErrIntOverflowPendingIcq
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
				return 0, ErrInvalidLengthPendingIcq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPendingIcq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPendingIcq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPendingIcq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPendingIcq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPendingIcq = fmt.Errorf("proto: unexpected end of group")
)
