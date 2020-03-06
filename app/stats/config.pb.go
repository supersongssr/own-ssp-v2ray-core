package stats

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

type Config struct {
	TrackIp              bool     `protobuf:"varint,1,opt,name=track_ip,json=trackIp,proto3" json:"track_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_d494ded44ceaa50d, []int{0}
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

func (m *Config) GetTrackIp() bool {
	if m != nil {
		return m.TrackIp
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.app.stats.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/stats/config.proto", fileDescriptor_d494ded44ceaa50d)
}

var fileDescriptor_d494ded44ceaa50d = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2c, 0x28, 0xd0, 0x2f,
	0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x81, 0x29, 0x2b, 0x4a, 0xd5, 0x4b, 0x2c, 0x28, 0xd0, 0x03, 0x2b, 0x51, 0x52,
	0xe6, 0x62, 0x73, 0x06, 0xab, 0x12, 0x92, 0xe4, 0xe2, 0x28, 0x29, 0x4a, 0x4c, 0xce, 0x8e, 0xcf,
	0x2c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x62, 0x07, 0xf3, 0x3d, 0x0b, 0x9c, 0xac, 0xb8,
	0x24, 0x92, 0xf3, 0x73, 0xf5, 0xb0, 0x19, 0x10, 0xc0, 0x18, 0xc5, 0x0a, 0x66, 0xac, 0x62, 0x12,
	0x09, 0x33, 0x0a, 0x4a, 0xac, 0xd4, 0x73, 0x06, 0xc9, 0x3b, 0x16, 0x14, 0xe8, 0x05, 0x83, 0x84,
	0x93, 0xd8, 0xc0, 0xb6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xdc, 0x6f, 0x3e, 0xa6,
	0x00, 0x00, 0x00,
}