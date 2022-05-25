// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icq/icq_result.proto

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

type DataPointResult struct {
	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocalHeight     uint64 `protobuf:"varint,2,opt,name=local_height,json=localHeight,proto3" json:"local_height,omitempty"`
	TargetHeight    uint64 `protobuf:"varint,3,opt,name=target_height,json=targetHeight,proto3" json:"target_height,omitempty"`
	Data            []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	LastDataPointId string `protobuf:"bytes,5,opt,name=last_data_point_id,json=lastDataPointId,proto3" json:"last_data_point_id,omitempty"`
}

func (m *DataPointResult) Reset()         { *m = DataPointResult{} }
func (m *DataPointResult) String() string { return proto.CompactTextString(m) }
func (*DataPointResult) ProtoMessage()    {}
func (*DataPointResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b751af249ed86cac, []int{0}
}
func (m *DataPointResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPointResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPointResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPointResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPointResult.Merge(m, src)
}
func (m *DataPointResult) XXX_Size() int {
	return m.Size()
}
func (m *DataPointResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPointResult.DiscardUnknown(m)
}

var xxx_messageInfo_DataPointResult proto.InternalMessageInfo

func (m *DataPointResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataPointResult) GetLocalHeight() uint64 {
	if m != nil {
		return m.LocalHeight
	}
	return 0
}

func (m *DataPointResult) GetTargetHeight() uint64 {
	if m != nil {
		return m.TargetHeight
	}
	return 0
}

func (m *DataPointResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DataPointResult) GetLastDataPointId() string {
	if m != nil {
		return m.LastDataPointId
	}
	return ""
}

type ICQResult struct {
	PeriodicId   uint64 `protobuf:"varint,1,opt,name=periodic_id,json=periodicId,proto3" json:"periodic_id,omitempty"`
	LastResultId string `protobuf:"bytes,2,opt,name=last_result_id,json=lastResultId,proto3" json:"last_result_id,omitempty"`
}

func (m *ICQResult) Reset()         { *m = ICQResult{} }
func (m *ICQResult) String() string { return proto.CompactTextString(m) }
func (*ICQResult) ProtoMessage()    {}
func (*ICQResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b751af249ed86cac, []int{1}
}
func (m *ICQResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ICQResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ICQResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ICQResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICQResult.Merge(m, src)
}
func (m *ICQResult) XXX_Size() int {
	return m.Size()
}
func (m *ICQResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ICQResult.DiscardUnknown(m)
}

var xxx_messageInfo_ICQResult proto.InternalMessageInfo

func (m *ICQResult) GetPeriodicId() uint64 {
	if m != nil {
		return m.PeriodicId
	}
	return 0
}

func (m *ICQResult) GetLastResultId() string {
	if m != nil {
		return m.LastResultId
	}
	return ""
}

func init() {
	proto.RegisterType((*DataPointResult)(nil), "simplyvc.interchainqueries.icq.DataPointResult")
	proto.RegisterType((*ICQResult)(nil), "simplyvc.interchainqueries.icq.ICQResult")
}

func init() { proto.RegisterFile("icq/icq_result.proto", fileDescriptor_b751af249ed86cac) }

var fileDescriptor_b751af249ed86cac = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xeb, 0x10, 0x90, 0xea, 0x86, 0x56, 0xb2, 0x18, 0x32, 0x99, 0x52, 0x18, 0x2a, 0x21,
	0x25, 0x42, 0xdc, 0x00, 0x18, 0x88, 0x58, 0x20, 0x23, 0x4b, 0xe4, 0xda, 0x56, 0xf3, 0x4b, 0x69,
	0x9d, 0x38, 0x7f, 0x11, 0xbd, 0x05, 0xe7, 0xe0, 0x24, 0x8c, 0x1d, 0x19, 0x51, 0x7b, 0x11, 0x64,
	0xb7, 0x61, 0x61, 0x4b, 0x3e, 0x3f, 0x3f, 0x3f, 0x7d, 0xf4, 0x0c, 0x64, 0x93, 0x82, 0x6c, 0x0a,
	0xab, 0xdb, 0x55, 0x85, 0x49, 0x6d, 0x0d, 0x1a, 0xc6, 0x5b, 0x58, 0xd4, 0xd5, 0xfa, 0x4d, 0x26,
	0xb0, 0x44, 0x6d, 0x65, 0x29, 0x60, 0xd9, 0xac, 0xb4, 0x05, 0xdd, 0x26, 0x20, 0x9b, 0xc9, 0x27,
	0xa1, 0xa3, 0x07, 0x81, 0xe2, 0xd9, 0xc0, 0x12, 0x73, 0x7f, 0x93, 0x0d, 0x69, 0x00, 0x2a, 0x26,
	0x63, 0x32, 0xed, 0xe7, 0x01, 0x28, 0x76, 0x41, 0xa3, 0xca, 0x48, 0x51, 0x15, 0xa5, 0x86, 0x79,
	0x89, 0x71, 0x30, 0x26, 0xd3, 0x30, 0x1f, 0x78, 0xf6, 0xe8, 0x11, 0xbb, 0xa4, 0xa7, 0x28, 0xec,
	0x5c, 0x63, 0x97, 0x39, 0xf2, 0x99, 0x68, 0x0f, 0x0f, 0x21, 0x46, 0x43, 0x25, 0x50, 0xc4, 0xe1,
	0x98, 0x4c, 0xa3, 0xdc, 0x7f, 0xb3, 0x6b, 0xca, 0x2a, 0xd1, 0x62, 0xe1, 0x7e, 0x8a, 0xda, 0x8d,
	0x28, 0x40, 0xc5, 0xc7, 0xfe, 0xed, 0x91, 0x3b, 0xf9, 0x1b, 0x97, 0xa9, 0x49, 0x4e, 0xfb, 0xd9,
	0xfd, 0xcb, 0x61, 0xe5, 0x39, 0x1d, 0xd4, 0xda, 0x82, 0x51, 0x20, 0x8b, 0xc3, 0xdc, 0x30, 0xa7,
	0x1d, 0xca, 0x14, 0xbb, 0xa2, 0x43, 0x5f, 0xbd, 0xf7, 0xe1, 0x32, 0x81, 0xaf, 0x8d, 0x1c, 0xdd,
	0x97, 0x64, 0xea, 0xee, 0xe9, 0x6b, 0xcb, 0xc9, 0x66, 0xcb, 0xc9, 0xcf, 0x96, 0x93, 0x8f, 0x1d,
	0xef, 0x6d, 0x76, 0xbc, 0xf7, 0xbd, 0xe3, 0xbd, 0xd7, 0x9b, 0x39, 0x60, 0xb9, 0x9a, 0x25, 0xd2,
	0x2c, 0xd2, 0xce, 0x62, 0xfa, 0xcf, 0x62, 0xfa, 0xee, 0xa4, 0xa7, 0xb8, 0xae, 0x75, 0x3b, 0x3b,
	0xf1, 0xd2, 0x6f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x30, 0xba, 0x9d, 0x84, 0x8c, 0x01, 0x00,
	0x00,
}

func (m *DataPointResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPointResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPointResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastDataPointId) > 0 {
		i -= len(m.LastDataPointId)
		copy(dAtA[i:], m.LastDataPointId)
		i = encodeVarintIcqResult(dAtA, i, uint64(len(m.LastDataPointId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintIcqResult(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.TargetHeight != 0 {
		i = encodeVarintIcqResult(dAtA, i, uint64(m.TargetHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.LocalHeight != 0 {
		i = encodeVarintIcqResult(dAtA, i, uint64(m.LocalHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIcqResult(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ICQResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ICQResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ICQResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LastResultId) > 0 {
		i -= len(m.LastResultId)
		copy(dAtA[i:], m.LastResultId)
		i = encodeVarintIcqResult(dAtA, i, uint64(len(m.LastResultId)))
		i--
		dAtA[i] = 0x12
	}
	if m.PeriodicId != 0 {
		i = encodeVarintIcqResult(dAtA, i, uint64(m.PeriodicId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcqResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcqResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataPointResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIcqResult(uint64(l))
	}
	if m.LocalHeight != 0 {
		n += 1 + sovIcqResult(uint64(m.LocalHeight))
	}
	if m.TargetHeight != 0 {
		n += 1 + sovIcqResult(uint64(m.TargetHeight))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovIcqResult(uint64(l))
	}
	l = len(m.LastDataPointId)
	if l > 0 {
		n += 1 + l + sovIcqResult(uint64(l))
	}
	return n
}

func (m *ICQResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeriodicId != 0 {
		n += 1 + sovIcqResult(uint64(m.PeriodicId))
	}
	l = len(m.LastResultId)
	if l > 0 {
		n += 1 + l + sovIcqResult(uint64(l))
	}
	return n
}

func sovIcqResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcqResult(x uint64) (n int) {
	return sovIcqResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataPointResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcqResult
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
			return fmt.Errorf("proto: DataPointResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPointResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcqResult
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
				return ErrInvalidLengthIcqResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcqResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalHeight", wireType)
			}
			m.LocalHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcqResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LocalHeight |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowIcqResult
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcqResult
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
				return ErrInvalidLengthIcqResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIcqResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastDataPointId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcqResult
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
				return ErrInvalidLengthIcqResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcqResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastDataPointId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcqResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcqResult
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
func (m *ICQResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcqResult
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
			return fmt.Errorf("proto: ICQResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ICQResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicId", wireType)
			}
			m.PeriodicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcqResult
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastResultId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcqResult
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
				return ErrInvalidLengthIcqResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcqResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastResultId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcqResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcqResult
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
func skipIcqResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcqResult
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
					return 0, ErrIntOverflowIcqResult
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
					return 0, ErrIntOverflowIcqResult
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
				return 0, ErrInvalidLengthIcqResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcqResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcqResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcqResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcqResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcqResult = fmt.Errorf("proto: unexpected end of group")
)