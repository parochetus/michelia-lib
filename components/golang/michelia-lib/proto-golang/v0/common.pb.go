// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package v0

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TableLocation struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Table                string   `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableLocation) Reset()         { *m = TableLocation{} }
func (m *TableLocation) String() string { return proto.CompactTextString(m) }
func (*TableLocation) ProtoMessage()    {}
func (*TableLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *TableLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableLocation.Unmarshal(m, b)
}
func (m *TableLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableLocation.Marshal(b, m, deterministic)
}
func (m *TableLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableLocation.Merge(m, src)
}
func (m *TableLocation) XXX_Size() int {
	return xxx_messageInfo_TableLocation.Size(m)
}
func (m *TableLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_TableLocation.DiscardUnknown(m)
}

var xxx_messageInfo_TableLocation proto.InternalMessageInfo

func (m *TableLocation) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TableLocation) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

type SliceConfiguration struct {
	Slot                 int32    `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Segment              int32    `protobuf:"varint,2,opt,name=segment,proto3" json:"segment,omitempty"`
	Slice                int32    `protobuf:"varint,3,opt,name=slice,proto3" json:"slice,omitempty"`
	Cache                bool     `protobuf:"varint,4,opt,name=cache,proto3" json:"cache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SliceConfiguration) Reset()         { *m = SliceConfiguration{} }
func (m *SliceConfiguration) String() string { return proto.CompactTextString(m) }
func (*SliceConfiguration) ProtoMessage()    {}
func (*SliceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *SliceConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SliceConfiguration.Unmarshal(m, b)
}
func (m *SliceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SliceConfiguration.Marshal(b, m, deterministic)
}
func (m *SliceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SliceConfiguration.Merge(m, src)
}
func (m *SliceConfiguration) XXX_Size() int {
	return xxx_messageInfo_SliceConfiguration.Size(m)
}
func (m *SliceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_SliceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_SliceConfiguration proto.InternalMessageInfo

func (m *SliceConfiguration) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *SliceConfiguration) GetSegment() int32 {
	if m != nil {
		return m.Segment
	}
	return 0
}

func (m *SliceConfiguration) GetSlice() int32 {
	if m != nil {
		return m.Slice
	}
	return 0
}

func (m *SliceConfiguration) GetCache() bool {
	if m != nil {
		return m.Cache
	}
	return false
}

type TendencyConfiguration struct {
	// tendency mode
	// mode 0: strict mode
	// mode 1: threshold mode
	Mode int32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	// toward 0: increase
	// toward 1: keep
	// toward 2: decrease
	Toward int32 `protobuf:"varint,2,opt,name=toward,proto3" json:"toward,omitempty"`
	// treat delta/1000 as the effective threshold
	Delta                int32    `protobuf:"varint,3,opt,name=delta,proto3" json:"delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TendencyConfiguration) Reset()         { *m = TendencyConfiguration{} }
func (m *TendencyConfiguration) String() string { return proto.CompactTextString(m) }
func (*TendencyConfiguration) ProtoMessage()    {}
func (*TendencyConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *TendencyConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TendencyConfiguration.Unmarshal(m, b)
}
func (m *TendencyConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TendencyConfiguration.Marshal(b, m, deterministic)
}
func (m *TendencyConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TendencyConfiguration.Merge(m, src)
}
func (m *TendencyConfiguration) XXX_Size() int {
	return xxx_messageInfo_TendencyConfiguration.Size(m)
}
func (m *TendencyConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_TendencyConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_TendencyConfiguration proto.InternalMessageInfo

func (m *TendencyConfiguration) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *TendencyConfiguration) GetToward() int32 {
	if m != nil {
		return m.Toward
	}
	return 0
}

func (m *TendencyConfiguration) GetDelta() int32 {
	if m != nil {
		return m.Delta
	}
	return 0
}

type SlotWindowConfiguration struct {
	Hole                 *HoleConfiguration `protobuf:"bytes,1,opt,name=hole,proto3" json:"hole,omitempty"`
	Size                 int32              `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SlotWindowConfiguration) Reset()         { *m = SlotWindowConfiguration{} }
func (m *SlotWindowConfiguration) String() string { return proto.CompactTextString(m) }
func (*SlotWindowConfiguration) ProtoMessage()    {}
func (*SlotWindowConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *SlotWindowConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlotWindowConfiguration.Unmarshal(m, b)
}
func (m *SlotWindowConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlotWindowConfiguration.Marshal(b, m, deterministic)
}
func (m *SlotWindowConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlotWindowConfiguration.Merge(m, src)
}
func (m *SlotWindowConfiguration) XXX_Size() int {
	return xxx_messageInfo_SlotWindowConfiguration.Size(m)
}
func (m *SlotWindowConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_SlotWindowConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_SlotWindowConfiguration proto.InternalMessageInfo

func (m *SlotWindowConfiguration) GetHole() *HoleConfiguration {
	if m != nil {
		return m.Hole
	}
	return nil
}

func (m *SlotWindowConfiguration) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type HoleConfiguration struct {
	// hole mode
	// mode 0: no hole
	// mode 1: single hole available
	// mode 2: multi-hole available
	Mode int32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	// total hole count
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// max continuous hole number
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HoleConfiguration) Reset()         { *m = HoleConfiguration{} }
func (m *HoleConfiguration) String() string { return proto.CompactTextString(m) }
func (*HoleConfiguration) ProtoMessage()    {}
func (*HoleConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *HoleConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HoleConfiguration.Unmarshal(m, b)
}
func (m *HoleConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HoleConfiguration.Marshal(b, m, deterministic)
}
func (m *HoleConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HoleConfiguration.Merge(m, src)
}
func (m *HoleConfiguration) XXX_Size() int {
	return xxx_messageInfo_HoleConfiguration.Size(m)
}
func (m *HoleConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_HoleConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_HoleConfiguration proto.InternalMessageInfo

func (m *HoleConfiguration) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *HoleConfiguration) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *HoleConfiguration) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type WindowOperationConfiguration struct {
	// window operation mode
	// mode 0: sample
	// mode 1: average
	// mode 2: moving average
	Mode int32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	// total hole count
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// max continuous hole number
	Offset               int32    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WindowOperationConfiguration) Reset()         { *m = WindowOperationConfiguration{} }
func (m *WindowOperationConfiguration) String() string { return proto.CompactTextString(m) }
func (*WindowOperationConfiguration) ProtoMessage()    {}
func (*WindowOperationConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *WindowOperationConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WindowOperationConfiguration.Unmarshal(m, b)
}
func (m *WindowOperationConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WindowOperationConfiguration.Marshal(b, m, deterministic)
}
func (m *WindowOperationConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WindowOperationConfiguration.Merge(m, src)
}
func (m *WindowOperationConfiguration) XXX_Size() int {
	return xxx_messageInfo_WindowOperationConfiguration.Size(m)
}
func (m *WindowOperationConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WindowOperationConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WindowOperationConfiguration proto.InternalMessageInfo

func (m *WindowOperationConfiguration) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *WindowOperationConfiguration) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *WindowOperationConfiguration) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type TimestampConfiguration struct {
	Moment               int64    `protobuf:"varint,1,opt,name=moment,proto3" json:"moment,omitempty"`
	Dir                  bool     `protobuf:"varint,2,opt,name=dir,proto3" json:"dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimestampConfiguration) Reset()         { *m = TimestampConfiguration{} }
func (m *TimestampConfiguration) String() string { return proto.CompactTextString(m) }
func (*TimestampConfiguration) ProtoMessage()    {}
func (*TimestampConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *TimestampConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimestampConfiguration.Unmarshal(m, b)
}
func (m *TimestampConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimestampConfiguration.Marshal(b, m, deterministic)
}
func (m *TimestampConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimestampConfiguration.Merge(m, src)
}
func (m *TimestampConfiguration) XXX_Size() int {
	return xxx_messageInfo_TimestampConfiguration.Size(m)
}
func (m *TimestampConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_TimestampConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_TimestampConfiguration proto.InternalMessageInfo

func (m *TimestampConfiguration) GetMoment() int64 {
	if m != nil {
		return m.Moment
	}
	return 0
}

func (m *TimestampConfiguration) GetDir() bool {
	if m != nil {
		return m.Dir
	}
	return false
}

type TagReduceConfiguration struct {
	// reduce mode
	// mode 0: do not reduce
	// mode 1: must contain all
	// mode 2: must not contain all
	// mode 3: contain any
	Mode                 int32    `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Head                 []string `protobuf:"bytes,2,rep,name=head,proto3" json:"head,omitempty"`
	Detail               []string `protobuf:"bytes,3,rep,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagReduceConfiguration) Reset()         { *m = TagReduceConfiguration{} }
func (m *TagReduceConfiguration) String() string { return proto.CompactTextString(m) }
func (*TagReduceConfiguration) ProtoMessage()    {}
func (*TagReduceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *TagReduceConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagReduceConfiguration.Unmarshal(m, b)
}
func (m *TagReduceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagReduceConfiguration.Marshal(b, m, deterministic)
}
func (m *TagReduceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagReduceConfiguration.Merge(m, src)
}
func (m *TagReduceConfiguration) XXX_Size() int {
	return xxx_messageInfo_TagReduceConfiguration.Size(m)
}
func (m *TagReduceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_TagReduceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_TagReduceConfiguration proto.InternalMessageInfo

func (m *TagReduceConfiguration) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *TagReduceConfiguration) GetHead() []string {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *TagReduceConfiguration) GetDetail() []string {
	if m != nil {
		return m.Detail
	}
	return nil
}

type SelectorConfiguration struct {
	// selector mode
	// mode 0: *
	// mode 1: for uid,un,domain match
	// mode 2: for tag reduce match
	Mode                 int32                   `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Uid                  string                  `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Un                   string                  `protobuf:"bytes,3,opt,name=un,proto3" json:"un,omitempty"`
	Domain               string                  `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Tag                  *TagReduceConfiguration `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SelectorConfiguration) Reset()         { *m = SelectorConfiguration{} }
func (m *SelectorConfiguration) String() string { return proto.CompactTextString(m) }
func (*SelectorConfiguration) ProtoMessage()    {}
func (*SelectorConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{8}
}

func (m *SelectorConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectorConfiguration.Unmarshal(m, b)
}
func (m *SelectorConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectorConfiguration.Marshal(b, m, deterministic)
}
func (m *SelectorConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectorConfiguration.Merge(m, src)
}
func (m *SelectorConfiguration) XXX_Size() int {
	return xxx_messageInfo_SelectorConfiguration.Size(m)
}
func (m *SelectorConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectorConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_SelectorConfiguration proto.InternalMessageInfo

func (m *SelectorConfiguration) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *SelectorConfiguration) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SelectorConfiguration) GetUn() string {
	if m != nil {
		return m.Un
	}
	return ""
}

func (m *SelectorConfiguration) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *SelectorConfiguration) GetTag() *TagReduceConfiguration {
	if m != nil {
		return m.Tag
	}
	return nil
}

type PlainSelectFilterReduceConfiguration struct {
	Selector             *SelectorConfiguration   `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	Slice                *SliceConfiguration      `protobuf:"bytes,2,opt,name=slice,proto3" json:"slice,omitempty"`
	Timestamp            *TimestampConfiguration  `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Slot                 *SlotWindowConfiguration `protobuf:"bytes,4,opt,name=slot,proto3" json:"slot,omitempty"`
	Tendency             *TendencyConfiguration   `protobuf:"bytes,5,opt,name=tendency,proto3" json:"tendency,omitempty"`
	Reduce               *TagReduceConfiguration  `protobuf:"bytes,6,opt,name=reduce,proto3" json:"reduce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PlainSelectFilterReduceConfiguration) Reset()         { *m = PlainSelectFilterReduceConfiguration{} }
func (m *PlainSelectFilterReduceConfiguration) String() string { return proto.CompactTextString(m) }
func (*PlainSelectFilterReduceConfiguration) ProtoMessage()    {}
func (*PlainSelectFilterReduceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{9}
}

func (m *PlainSelectFilterReduceConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainSelectFilterReduceConfiguration.Unmarshal(m, b)
}
func (m *PlainSelectFilterReduceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainSelectFilterReduceConfiguration.Marshal(b, m, deterministic)
}
func (m *PlainSelectFilterReduceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainSelectFilterReduceConfiguration.Merge(m, src)
}
func (m *PlainSelectFilterReduceConfiguration) XXX_Size() int {
	return xxx_messageInfo_PlainSelectFilterReduceConfiguration.Size(m)
}
func (m *PlainSelectFilterReduceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainSelectFilterReduceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_PlainSelectFilterReduceConfiguration proto.InternalMessageInfo

func (m *PlainSelectFilterReduceConfiguration) GetSelector() *SelectorConfiguration {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *PlainSelectFilterReduceConfiguration) GetSlice() *SliceConfiguration {
	if m != nil {
		return m.Slice
	}
	return nil
}

func (m *PlainSelectFilterReduceConfiguration) GetTimestamp() *TimestampConfiguration {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PlainSelectFilterReduceConfiguration) GetSlot() *SlotWindowConfiguration {
	if m != nil {
		return m.Slot
	}
	return nil
}

func (m *PlainSelectFilterReduceConfiguration) GetTendency() *TendencyConfiguration {
	if m != nil {
		return m.Tendency
	}
	return nil
}

func (m *PlainSelectFilterReduceConfiguration) GetReduce() *TagReduceConfiguration {
	if m != nil {
		return m.Reduce
	}
	return nil
}

func init() {
	proto.RegisterType((*TableLocation)(nil), "mechelia.v0.TableLocation")
	proto.RegisterType((*SliceConfiguration)(nil), "mechelia.v0.SliceConfiguration")
	proto.RegisterType((*TendencyConfiguration)(nil), "mechelia.v0.TendencyConfiguration")
	proto.RegisterType((*SlotWindowConfiguration)(nil), "mechelia.v0.SlotWindowConfiguration")
	proto.RegisterType((*HoleConfiguration)(nil), "mechelia.v0.HoleConfiguration")
	proto.RegisterType((*WindowOperationConfiguration)(nil), "mechelia.v0.WindowOperationConfiguration")
	proto.RegisterType((*TimestampConfiguration)(nil), "mechelia.v0.TimestampConfiguration")
	proto.RegisterType((*TagReduceConfiguration)(nil), "mechelia.v0.TagReduceConfiguration")
	proto.RegisterType((*SelectorConfiguration)(nil), "mechelia.v0.SelectorConfiguration")
	proto.RegisterType((*PlainSelectFilterReduceConfiguration)(nil), "mechelia.v0.PlainSelectFilterReduceConfiguration")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xe6, 0x2e, 0x77, 0xe7, 0xdd, 0x54, 0xa5, 0x5d, 0xed, 0x99, 0x87, 0xa2, 0x25, 0xf6, 0xa1,
	0x2f, 0xa6, 0xa5, 0x52, 0x10, 0x84, 0x82, 0x2d, 0x88, 0x0f, 0x82, 0xb2, 0x77, 0xe2, 0x8f, 0x07,
	0x61, 0x9b, 0xcc, 0xe5, 0x16, 0x36, 0xbb, 0x47, 0xb2, 0x69, 0xd1, 0x7f, 0xc4, 0x17, 0xff, 0x58,
	0xd9, 0xc9, 0xe6, 0x6a, 0xbc, 0x48, 0xfb, 0x36, 0xdf, 0x66, 0xe7, 0xfb, 0xe6, 0x9b, 0x9d, 0x09,
	0xdc, 0x4f, 0x4c, 0x9e, 0x1b, 0x1d, 0xaf, 0x0a, 0x63, 0x0d, 0xdb, 0xca, 0x31, 0x59, 0xa2, 0x92,
	0x22, 0xbe, 0x3a, 0x8e, 0x2e, 0xe0, 0xc1, 0x5c, 0x5c, 0x2a, 0x7c, 0x6f, 0x12, 0x61, 0xa5, 0xd1,
	0x6c, 0x0f, 0x26, 0x5a, 0xe4, 0x58, 0xae, 0x44, 0x82, 0x61, 0x6f, 0xbf, 0x77, 0x38, 0xe1, 0x37,
	0x07, 0xec, 0x31, 0x0c, 0xad, 0xbb, 0x1e, 0xf6, 0xe9, 0x4b, 0x0d, 0x22, 0x0d, 0x6c, 0xa6, 0x64,
	0x82, 0x17, 0x46, 0x2f, 0x64, 0x56, 0x15, 0x35, 0x13, 0x83, 0x41, 0xa9, 0x8c, 0x25, 0x92, 0x21,
	0xa7, 0x98, 0x85, 0x70, 0xaf, 0xc4, 0x2c, 0x47, 0x6d, 0x89, 0x61, 0xc8, 0x1b, 0xe8, 0x98, 0x4b,
	0xc7, 0x11, 0x06, 0x74, 0x5e, 0x03, 0x77, 0x9a, 0x88, 0x64, 0x89, 0xe1, 0x60, 0xbf, 0x77, 0x38,
	0xe6, 0x35, 0x88, 0xbe, 0xc2, 0xee, 0x1c, 0x75, 0x8a, 0x3a, 0xf9, 0xb1, 0x21, 0x99, 0x9b, 0x14,
	0x1b, 0x49, 0x17, 0xb3, 0x29, 0x8c, 0xac, 0xb9, 0x16, 0x45, 0xea, 0x15, 0x3d, 0x72, 0xd4, 0x29,
	0x2a, 0x2b, 0x1a, 0x41, 0x02, 0x91, 0x80, 0x27, 0x33, 0x65, 0xec, 0x67, 0xa9, 0x53, 0x73, 0xdd,
	0x26, 0x3f, 0x81, 0xc1, 0xd2, 0xa8, 0x9a, 0x7c, 0xeb, 0xe4, 0x69, 0xfc, 0x57, 0x1b, 0xe3, 0x77,
	0x46, 0xb5, 0xdd, 0x73, 0xba, 0x4b, 0x3d, 0x90, 0x3f, 0xd1, 0x4b, 0x53, 0x1c, 0x7d, 0x82, 0x9d,
	0x8d, 0xeb, 0x9d, 0x95, 0x3b, 0xf3, 0xa6, 0x5a, 0xb7, 0xaa, 0x06, 0xce, 0x8f, 0x59, 0x2c, 0x4a,
	0xb4, 0xbe, 0x70, 0x8f, 0xa2, 0xef, 0xb0, 0x57, 0x57, 0xfd, 0x61, 0x85, 0x35, 0xe9, 0xed, 0x0a,
	0x1d, 0xe5, 0xfd, 0x97, 0xff, 0x1c, 0xa6, 0x73, 0x99, 0x63, 0x69, 0x45, 0xbe, 0x6a, 0x33, 0x4f,
	0x61, 0x94, 0x1b, 0x7a, 0x53, 0xc7, 0x1d, 0x70, 0x8f, 0xd8, 0x36, 0x04, 0xa9, 0x2c, 0x88, 0x7c,
	0xcc, 0x5d, 0x18, 0x7d, 0x81, 0xe9, 0x5c, 0x64, 0x1c, 0xd3, 0x2a, 0xc1, 0x3b, 0x55, 0xb7, 0x44,
	0xe1, 0xde, 0x2d, 0x38, 0x9c, 0x70, 0x8a, 0x9d, 0x56, 0x8a, 0x56, 0x48, 0x15, 0x06, 0x74, 0xea,
	0x51, 0xf4, 0xbb, 0x07, 0xbb, 0x33, 0x54, 0x98, 0x58, 0x53, 0xdc, 0xce, 0xbc, 0x0d, 0x41, 0x25,
	0x53, 0x3f, 0xc4, 0x2e, 0x64, 0x0f, 0xa1, 0x5f, 0x69, 0x72, 0x3c, 0xe1, 0xfd, 0x8a, 0x3c, 0xa5,
	0x26, 0x17, 0x52, 0xd3, 0xe4, 0x39, 0x1d, 0x42, 0xec, 0x14, 0x02, 0x2b, 0xb2, 0x70, 0x48, 0x33,
	0xf0, 0xbc, 0x35, 0x03, 0xdd, 0xce, 0xb8, 0xbb, 0x1f, 0xfd, 0x0a, 0xe0, 0xe0, 0xa3, 0x12, 0x52,
	0xd7, 0x35, 0xbe, 0x95, 0xca, 0x62, 0xd1, 0xd5, 0x87, 0x33, 0x18, 0x97, 0xde, 0x86, 0x1f, 0xb4,
	0xa8, 0x25, 0xd2, 0xe9, 0x91, 0xaf, 0x73, 0xd8, 0x69, 0xb3, 0x46, 0x7d, 0x4a, 0x7e, 0xd6, 0x4e,
	0xde, 0x58, 0xd2, 0x66, 0xcf, 0xde, 0xc0, 0xc4, 0x36, 0x8f, 0x4b, 0x5d, 0xd8, 0x30, 0xd7, 0xf9,
	0xf4, 0xfc, 0x26, 0x8b, 0xbd, 0xf2, 0xeb, 0x3e, 0xa0, 0xec, 0x83, 0x7f, 0x84, 0x3b, 0x57, 0xca,
	0xff, 0x14, 0xce, 0x60, 0x6c, 0xfd, 0x3a, 0xfb, 0xc6, 0xb6, 0x3d, 0x77, 0xee, 0x3a, 0x5f, 0xe7,
	0xb0, 0xd7, 0x30, 0x2a, 0xa8, 0x95, 0xe1, 0xe8, 0xee, 0xcf, 0xe2, 0x53, 0xce, 0x1f, 0x7d, 0xdb,
	0x89, 0x8f, 0xe8, 0xc7, 0xf8, 0x22, 0x33, 0x4a, 0xe8, 0xec, 0xe8, 0xea, 0xf8, 0x72, 0x44, 0x07,
	0x2f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x86, 0x1a, 0xca, 0xa2, 0x39, 0x05, 0x00, 0x00,
}
