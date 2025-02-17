// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/node_version_info.proto

package entities

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

type NodeVersionInfo struct {
	Semver               string   `protobuf:"bytes,1,opt,name=semver,proto3" json:"semver,omitempty"`
	Commit               string   `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	SporkId              []byte   `protobuf:"bytes,3,opt,name=spork_id,json=sporkId,proto3" json:"spork_id,omitempty"`
	ProtocolVersion      uint64   `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeVersionInfo) Reset()         { *m = NodeVersionInfo{} }
func (m *NodeVersionInfo) String() string { return proto.CompactTextString(m) }
func (*NodeVersionInfo) ProtoMessage()    {}
func (*NodeVersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9e447351cb4a8a, []int{0}
}

func (m *NodeVersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeVersionInfo.Unmarshal(m, b)
}
func (m *NodeVersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeVersionInfo.Marshal(b, m, deterministic)
}
func (m *NodeVersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeVersionInfo.Merge(m, src)
}
func (m *NodeVersionInfo) XXX_Size() int {
	return xxx_messageInfo_NodeVersionInfo.Size(m)
}
func (m *NodeVersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeVersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeVersionInfo proto.InternalMessageInfo

func (m *NodeVersionInfo) GetSemver() string {
	if m != nil {
		return m.Semver
	}
	return ""
}

func (m *NodeVersionInfo) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *NodeVersionInfo) GetSporkId() []byte {
	if m != nil {
		return m.SporkId
	}
	return nil
}

func (m *NodeVersionInfo) GetProtocolVersion() uint64 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeVersionInfo)(nil), "flow.entities.NodeVersionInfo")
}

func init() {
	proto.RegisterFile("flow/entities/node_version_info.proto", fileDescriptor_ff9e447351cb4a8a)
}

var fileDescriptor_ff9e447351cb4a8a = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4a, 0xc0, 0x30,
	0x18, 0x84, 0x89, 0x96, 0xaa, 0x41, 0xa9, 0x64, 0x90, 0x08, 0x0e, 0x45, 0x10, 0xea, 0x92, 0x08,
	0xbe, 0x81, 0x5b, 0x17, 0x91, 0x0e, 0x0e, 0x2e, 0x85, 0x36, 0x49, 0x0d, 0x36, 0xf9, 0x4b, 0x92,
	0xd6, 0x47, 0xf0, 0xb5, 0xa5, 0x49, 0x23, 0xb8, 0xfc, 0x70, 0x77, 0x1f, 0xdc, 0xf1, 0xe3, 0x07,
	0x35, 0xc3, 0x37, 0x97, 0x36, 0xe8, 0xa0, 0xa5, 0xe7, 0x16, 0x84, 0xec, 0x37, 0xe9, 0xbc, 0x06,
	0xdb, 0x6b, 0xab, 0x80, 0x2d, 0x0e, 0x02, 0x90, 0xab, 0x1d, 0x63, 0x19, 0xbb, 0xff, 0x41, 0xb8,
	0x7a, 0x05, 0x21, 0xdf, 0x13, 0xd9, 0x5a, 0x05, 0xe4, 0x06, 0x97, 0x5e, 0x9a, 0x4d, 0x3a, 0x8a,
	0x6a, 0xd4, 0x5c, 0x74, 0x87, 0xda, 0xfd, 0x11, 0x8c, 0xd1, 0x81, 0x9e, 0x24, 0x3f, 0x29, 0x72,
	0x8b, 0xcf, 0xfd, 0x02, 0xee, 0xab, 0xd7, 0x82, 0x9e, 0xd6, 0xa8, 0xb9, 0xec, 0xce, 0xa2, 0x6e,
	0x05, 0x79, 0xc4, 0xd7, 0xb1, 0x76, 0x84, 0x39, 0x8f, 0xa1, 0x45, 0x8d, 0x9a, 0xa2, 0xab, 0xb2,
	0x7f, 0x34, 0xbf, 0xbc, 0xe1, 0x3b, 0x70, 0x13, 0x03, 0x1b, 0x07, 0xc6, 0x74, 0x58, 0xd5, 0xdf,
	0xd2, 0x8f, 0xa7, 0x49, 0x87, 0xcf, 0x75, 0x60, 0x23, 0x18, 0x9e, 0x20, 0x1e, 0x4f, 0x26, 0xf9,
	0x04, 0xfc, 0xdf, 0x0b, 0x86, 0x32, 0x46, 0xcf, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01, 0x2b,
	0x0a, 0xa1, 0x1a, 0x01, 0x00, 0x00,
}
