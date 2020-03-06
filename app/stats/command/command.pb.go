package command

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetStatsRequest struct {
	// Name of the stat counter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether or not to reset the counter to fetching its value.
	Reset_               bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{0}
}

func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type Stat struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{1}
}

func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stat) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetStatsResponse struct {
	Stat                 *Stat    `protobuf:"bytes,1,opt,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{2}
}

func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(m, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetStat() *Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type QueryStatsRequest struct {
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Reset_               bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatsRequest) Reset()         { *m = QueryStatsRequest{} }
func (m *QueryStatsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStatsRequest) ProtoMessage()    {}
func (*QueryStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{3}
}

func (m *QueryStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatsRequest.Unmarshal(m, b)
}
func (m *QueryStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatsRequest.Marshal(b, m, deterministic)
}
func (m *QueryStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatsRequest.Merge(m, src)
}
func (m *QueryStatsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryStatsRequest.Size(m)
}
func (m *QueryStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatsRequest proto.InternalMessageInfo

func (m *QueryStatsRequest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *QueryStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type QueryStatsResponse struct {
	Stat                 []*Stat  `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatsResponse) Reset()         { *m = QueryStatsResponse{} }
func (m *QueryStatsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStatsResponse) ProtoMessage()    {}
func (*QueryStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{4}
}

func (m *QueryStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatsResponse.Unmarshal(m, b)
}
func (m *QueryStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatsResponse.Marshal(b, m, deterministic)
}
func (m *QueryStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatsResponse.Merge(m, src)
}
func (m *QueryStatsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryStatsResponse.Size(m)
}
func (m *QueryStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatsResponse proto.InternalMessageInfo

func (m *QueryStatsResponse) GetStat() []*Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type SysStatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SysStatsRequest) Reset()         { *m = SysStatsRequest{} }
func (m *SysStatsRequest) String() string { return proto.CompactTextString(m) }
func (*SysStatsRequest) ProtoMessage()    {}
func (*SysStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{5}
}

func (m *SysStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SysStatsRequest.Unmarshal(m, b)
}
func (m *SysStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SysStatsRequest.Marshal(b, m, deterministic)
}
func (m *SysStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SysStatsRequest.Merge(m, src)
}
func (m *SysStatsRequest) XXX_Size() int {
	return xxx_messageInfo_SysStatsRequest.Size(m)
}
func (m *SysStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SysStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SysStatsRequest proto.InternalMessageInfo

type SysStatsResponse struct {
	NumGoroutine         uint32   `protobuf:"varint,1,opt,name=NumGoroutine,proto3" json:"NumGoroutine,omitempty"`
	NumGC                uint32   `protobuf:"varint,2,opt,name=NumGC,proto3" json:"NumGC,omitempty"`
	Alloc                uint64   `protobuf:"varint,3,opt,name=Alloc,proto3" json:"Alloc,omitempty"`
	TotalAlloc           uint64   `protobuf:"varint,4,opt,name=TotalAlloc,proto3" json:"TotalAlloc,omitempty"`
	Sys                  uint64   `protobuf:"varint,5,opt,name=Sys,proto3" json:"Sys,omitempty"`
	Mallocs              uint64   `protobuf:"varint,6,opt,name=Mallocs,proto3" json:"Mallocs,omitempty"`
	Frees                uint64   `protobuf:"varint,7,opt,name=Frees,proto3" json:"Frees,omitempty"`
	LiveObjects          uint64   `protobuf:"varint,8,opt,name=LiveObjects,proto3" json:"LiveObjects,omitempty"`
	PauseTotalNs         uint64   `protobuf:"varint,9,opt,name=PauseTotalNs,proto3" json:"PauseTotalNs,omitempty"`
	Uptime               uint32   `protobuf:"varint,10,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SysStatsResponse) Reset()         { *m = SysStatsResponse{} }
func (m *SysStatsResponse) String() string { return proto.CompactTextString(m) }
func (*SysStatsResponse) ProtoMessage()    {}
func (*SysStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{6}
}

func (m *SysStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SysStatsResponse.Unmarshal(m, b)
}
func (m *SysStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SysStatsResponse.Marshal(b, m, deterministic)
}
func (m *SysStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SysStatsResponse.Merge(m, src)
}
func (m *SysStatsResponse) XXX_Size() int {
	return xxx_messageInfo_SysStatsResponse.Size(m)
}
func (m *SysStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SysStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SysStatsResponse proto.InternalMessageInfo

func (m *SysStatsResponse) GetNumGoroutine() uint32 {
	if m != nil {
		return m.NumGoroutine
	}
	return 0
}

func (m *SysStatsResponse) GetNumGC() uint32 {
	if m != nil {
		return m.NumGC
	}
	return 0
}

func (m *SysStatsResponse) GetAlloc() uint64 {
	if m != nil {
		return m.Alloc
	}
	return 0
}

func (m *SysStatsResponse) GetTotalAlloc() uint64 {
	if m != nil {
		return m.TotalAlloc
	}
	return 0
}

func (m *SysStatsResponse) GetSys() uint64 {
	if m != nil {
		return m.Sys
	}
	return 0
}

func (m *SysStatsResponse) GetMallocs() uint64 {
	if m != nil {
		return m.Mallocs
	}
	return 0
}

func (m *SysStatsResponse) GetFrees() uint64 {
	if m != nil {
		return m.Frees
	}
	return 0
}

func (m *SysStatsResponse) GetLiveObjects() uint64 {
	if m != nil {
		return m.LiveObjects
	}
	return 0
}

func (m *SysStatsResponse) GetPauseTotalNs() uint64 {
	if m != nil {
		return m.PauseTotalNs
	}
	return 0
}

func (m *SysStatsResponse) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

type GetUserIPStatsResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserIPStatsResponse) Reset()         { *m = GetUserIPStatsResponse{} }
func (m *GetUserIPStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserIPStatsResponse) ProtoMessage()    {}
func (*GetUserIPStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{7}
}

func (m *GetUserIPStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserIPStatsResponse.Unmarshal(m, b)
}
func (m *GetUserIPStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserIPStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserIPStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserIPStatsResponse.Merge(m, src)
}
func (m *GetUserIPStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserIPStatsResponse.Size(m)
}
func (m *GetUserIPStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserIPStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserIPStatsResponse proto.InternalMessageInfo

func (m *GetUserIPStatsResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserIPStatsResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{8}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetStatsRequest)(nil), "v2ray.core.app.stats.command.GetStatsRequest")
	proto.RegisterType((*Stat)(nil), "v2ray.core.app.stats.command.Stat")
	proto.RegisterType((*GetStatsResponse)(nil), "v2ray.core.app.stats.command.GetStatsResponse")
	proto.RegisterType((*QueryStatsRequest)(nil), "v2ray.core.app.stats.command.QueryStatsRequest")
	proto.RegisterType((*QueryStatsResponse)(nil), "v2ray.core.app.stats.command.QueryStatsResponse")
	proto.RegisterType((*SysStatsRequest)(nil), "v2ray.core.app.stats.command.SysStatsRequest")
	proto.RegisterType((*SysStatsResponse)(nil), "v2ray.core.app.stats.command.SysStatsResponse")
	proto.RegisterType((*GetUserIPStatsResponse)(nil), "v2ray.core.app.stats.command.GetUserIPStatsResponse")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.stats.command.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/stats/command/command.proto", fileDescriptor_c902411c4948f26b)
}

var fileDescriptor_c902411c4948f26b = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xcd, 0x57, 0xf3, 0x31, 0x69, 0x69, 0x6a, 0xa1, 0xca, 0xaa, 0x2a, 0x14, 0xf9, 0x94, 0x0b,
	0x4e, 0x15, 0x10, 0x17, 0x4e, 0x6d, 0x24, 0x22, 0x50, 0x28, 0x61, 0x43, 0x39, 0x70, 0x73, 0x97,
	0x01, 0x05, 0xb2, 0x6b, 0xd7, 0xf6, 0x06, 0xe5, 0xef, 0x70, 0xe4, 0x0f, 0xf1, 0x77, 0x90, 0xbd,
	0xbb, 0x24, 0x9b, 0xb6, 0x49, 0x7b, 0x8a, 0xdf, 0xcc, 0x3c, 0xcf, 0x9b, 0xd9, 0xe7, 0x00, 0x5f,
	0x0c, 0xb4, 0x58, 0xf2, 0x50, 0x46, 0xfd, 0x50, 0x6a, 0xec, 0x0b, 0xa5, 0xfa, 0xc6, 0x0a, 0x6b,
	0xfa, 0xa1, 0x8c, 0x22, 0x11, 0x7f, 0xcd, 0x7f, 0xb9, 0xd2, 0xd2, 0x4a, 0x72, 0x9a, 0xd7, 0x6b,
	0xe4, 0x42, 0x29, 0xee, 0x6b, 0x79, 0x56, 0xc3, 0x5e, 0xc3, 0xe1, 0x08, 0xed, 0xd4, 0xc5, 0x02,
	0xbc, 0x49, 0xd0, 0x58, 0x42, 0xa0, 0x16, 0x8b, 0x08, 0x69, 0xb9, 0x5b, 0xee, 0xb5, 0x02, 0x7f,
	0x26, 0x4f, 0x61, 0x4f, 0xa3, 0x41, 0x4b, 0x2b, 0xdd, 0x72, 0xaf, 0x19, 0xa4, 0x80, 0x9d, 0x41,
	0xcd, 0x31, 0xef, 0x63, 0x2c, 0xc4, 0x3c, 0x41, 0xcf, 0xa8, 0x06, 0x29, 0x60, 0xef, 0xa0, 0xb3,
	0x6a, 0x67, 0x94, 0x8c, 0x0d, 0x92, 0x57, 0x50, 0x73, 0x9a, 0x3c, 0xbb, 0x3d, 0x60, 0x7c, 0x9b,
	0x5e, 0xee, 0xa8, 0x81, 0xaf, 0x67, 0x43, 0x38, 0xfa, 0x98, 0xa0, 0x5e, 0x16, 0xc4, 0x53, 0x68,
	0x28, 0x61, 0x2d, 0xea, 0x38, 0x53, 0x93, 0xc3, 0x7b, 0x46, 0x18, 0x03, 0x59, 0xbf, 0xe4, 0x96,
	0xa4, 0xea, 0xa3, 0x24, 0x1d, 0xc1, 0xe1, 0x74, 0x69, 0xd6, 0x05, 0xb1, 0xdf, 0x15, 0xe8, 0xac,
	0x62, 0xd9, 0xfd, 0x0c, 0xf6, 0x2f, 0x93, 0x68, 0x24, 0xb5, 0x4c, 0xec, 0x2c, 0x4e, 0x17, 0x77,
	0x10, 0x14, 0x62, 0x4e, 0xaf, 0xc3, 0x43, 0xaf, 0xf7, 0x20, 0x48, 0x81, 0x8b, 0x9e, 0xcf, 0xe7,
	0x32, 0xa4, 0xd5, 0x6e, 0xb9, 0x57, 0x0b, 0x52, 0x40, 0x9e, 0x01, 0x7c, 0x92, 0x56, 0xcc, 0xd3,
	0x54, 0xcd, 0xa7, 0xd6, 0x22, 0xa4, 0x03, 0xd5, 0xe9, 0xd2, 0xd0, 0x3d, 0x9f, 0x70, 0x47, 0xb7,
	0xa7, 0xf7, 0xc2, 0xe5, 0x0c, 0xad, 0xfb, 0x68, 0x0e, 0x5d, 0x87, 0x37, 0x1a, 0xd1, 0xd0, 0x46,
	0xda, 0xc1, 0x03, 0xd2, 0x85, 0xf6, 0x78, 0xb6, 0xc0, 0x0f, 0xd7, 0x3f, 0x30, 0xb4, 0x86, 0x36,
	0x7d, 0x6e, 0x3d, 0xe4, 0x66, 0x9a, 0x88, 0xc4, 0xa0, 0x6f, 0x7b, 0x69, 0x68, 0xcb, 0x97, 0x14,
	0x62, 0xe4, 0x18, 0xea, 0x57, 0xca, 0xce, 0x22, 0xa4, 0xe0, 0x87, 0xca, 0x10, 0xbb, 0x80, 0xe3,
	0x11, 0xda, 0x2b, 0x83, 0xfa, 0xed, 0xa4, 0xb8, 0xa9, 0x9d, 0xd6, 0x6a, 0xe5, 0xd6, 0x6a, 0x42,
	0x7d, 0x28, 0xe3, 0x6f, 0xb3, 0xef, 0x83, 0xbf, 0x55, 0xd8, 0xf7, 0xb7, 0x4c, 0x51, 0x2f, 0x66,
	0x21, 0x92, 0x9f, 0xd0, 0xcc, 0x5d, 0x47, 0x9e, 0x6f, 0xff, 0x98, 0x1b, 0x8f, 0xe1, 0x84, 0x3f,
	0xb4, 0x3c, 0xd5, 0xcb, 0x4a, 0xe4, 0x06, 0x60, 0xe5, 0x28, 0xd2, 0xdf, 0xce, 0xbf, 0x65, 0xe0,
	0x93, 0xb3, 0x87, 0x13, 0xfe, 0xb7, 0x8c, 0xa1, 0xed, 0x84, 0x64, 0x2e, 0xdb, 0x35, 0xe2, 0x86,
	0x43, 0x77, 0x8d, 0xb8, 0x69, 0x5e, 0x56, 0x22, 0xbf, 0xe0, 0x49, 0xf1, 0x73, 0x3d, 0x76, 0xab,
	0x2f, 0x77, 0x96, 0xdf, 0xe1, 0x05, 0x56, 0xba, 0x18, 0x43, 0x37, 0x94, 0xd1, 0x56, 0xf2, 0xa4,
	0xfc, 0xa5, 0x91, 0x1d, 0xff, 0x54, 0x4e, 0x3f, 0x0f, 0x02, 0xb1, 0xe4, 0x43, 0x57, 0x79, 0xae,
	0x94, 0x7f, 0xaa, 0x86, 0x0f, 0xd3, 0xf4, 0x75, 0xdd, 0xff, 0x41, 0xbe, 0xf8, 0x17, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0xe7, 0x88, 0x07, 0x52, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*QueryStatsResponse, error)
	GetSysStats(ctx context.Context, in *SysStatsRequest, opts ...grpc.CallOption) (*SysStatsResponse, error)
	GetUserIPStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetUserIPStatsResponse, error)
}

type statsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatsServiceClient(cc *grpc.ClientConn) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*QueryStatsResponse, error) {
	out := new(QueryStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetSysStats(ctx context.Context, in *SysStatsRequest, opts ...grpc.CallOption) (*SysStatsResponse, error) {
	out := new(SysStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetSysStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetUserIPStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetUserIPStatsResponse, error) {
	out := new(GetUserIPStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetUserIPStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
type StatsServiceServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	QueryStats(context.Context, *QueryStatsRequest) (*QueryStatsResponse, error)
	GetSysStats(context.Context, *SysStatsRequest) (*SysStatsResponse, error)
	GetUserIPStats(context.Context, *GetStatsRequest) (*GetUserIPStatsResponse, error)
}

// UnimplementedStatsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (*UnimplementedStatsServiceServer) GetStats(ctx context.Context, req *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (*UnimplementedStatsServiceServer) QueryStats(ctx context.Context, req *QueryStatsRequest) (*QueryStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStats not implemented")
}
func (*UnimplementedStatsServiceServer) GetSysStats(ctx context.Context, req *SysStatsRequest) (*SysStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysStats not implemented")
}
func (*UnimplementedStatsServiceServer) GetUserIPStats(ctx context.Context, req *GetStatsRequest) (*GetUserIPStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIPStats not implemented")
}

func RegisterStatsServiceServer(s *grpc.Server, srv StatsServiceServer) {
	s.RegisterService(&_StatsService_serviceDesc, srv)
}

func _StatsService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).QueryStats(ctx, req.(*QueryStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetSysStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetSysStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetSysStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetSysStats(ctx, req.(*SysStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetUserIPStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetUserIPStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetUserIPStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetUserIPStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2ray.core.app.stats.command.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _StatsService_GetStats_Handler,
		},
		{
			MethodName: "QueryStats",
			Handler:    _StatsService_QueryStats_Handler,
		},
		{
			MethodName: "GetSysStats",
			Handler:    _StatsService_GetSysStats_Handler,
		},
		{
			MethodName: "GetUserIPStats",
			Handler:    _StatsService_GetUserIPStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2ray.com/core/app/stats/command/command.proto",
}