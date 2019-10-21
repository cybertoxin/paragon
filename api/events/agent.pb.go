// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	codec "github.com/kcarretto/paragon/api/codec"
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

type AgentCheckin struct {
	PublicIP             string               `protobuf:"bytes,1,opt,name=publicIP,proto3" json:"publicIP,omitempty"`
	SeenTime             int64                `protobuf:"varint,2,opt,name=SeenTime,proto3" json:"SeenTime,omitempty"`
	Agent                *codec.AgentMetadata `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AgentCheckin) Reset()         { *m = AgentCheckin{} }
func (m *AgentCheckin) String() string { return proto.CompactTextString(m) }
func (*AgentCheckin) ProtoMessage()    {}
func (*AgentCheckin) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *AgentCheckin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentCheckin.Unmarshal(m, b)
}
func (m *AgentCheckin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentCheckin.Marshal(b, m, deterministic)
}
func (m *AgentCheckin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentCheckin.Merge(m, src)
}
func (m *AgentCheckin) XXX_Size() int {
	return xxx_messageInfo_AgentCheckin.Size(m)
}
func (m *AgentCheckin) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentCheckin.DiscardUnknown(m)
}

var xxx_messageInfo_AgentCheckin proto.InternalMessageInfo

func (m *AgentCheckin) GetPublicIP() string {
	if m != nil {
		return m.PublicIP
	}
	return ""
}

func (m *AgentCheckin) GetSeenTime() int64 {
	if m != nil {
		return m.SeenTime
	}
	return 0
}

func (m *AgentCheckin) GetAgent() *codec.AgentMetadata {
	if m != nil {
		return m.Agent
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentCheckin)(nil), "events.AgentCheckin")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x96,
	0x12, 0x4d, 0x2c, 0xc8, 0xd4, 0x4f, 0xce, 0x4f, 0x49, 0x4d, 0xd6, 0x47, 0x92, 0x56, 0x2a, 0xe2,
	0xe2, 0x71, 0x04, 0x71, 0x9d, 0x33, 0x52, 0x93, 0xb3, 0x33, 0xf3, 0x84, 0xa4, 0xb8, 0x38, 0x0a,
	0x4a, 0x93, 0x72, 0x32, 0x93, 0x3d, 0x03, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c,
	0x90, 0x5c, 0x70, 0x6a, 0x6a, 0x5e, 0x48, 0x66, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x73,
	0x10, 0x9c, 0x2f, 0xa4, 0xc5, 0xc5, 0x0a, 0x36, 0x56, 0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x44, 0x0f, 0x6c, 0x95, 0x1e, 0xd8, 0x6c, 0xdf, 0xd4, 0x92, 0xc4, 0x94, 0xc4, 0x92, 0xc4, 0x20,
	0x88, 0x12, 0x27, 0xcd, 0x28, 0xf5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0xfd, 0xec, 0xe4, 0xc4, 0xa2, 0xa2, 0xd4, 0x92, 0x92, 0x7c, 0xfd, 0x82, 0xc4, 0xa2, 0xc4, 0xf4,
	0xfc, 0x3c, 0x7d, 0x90, 0x4b, 0x21, 0xae, 0x4e, 0x62, 0x03, 0xbb, 0xd2, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0xfd, 0x27, 0x3a, 0xd3, 0x00, 0x00, 0x00,
}