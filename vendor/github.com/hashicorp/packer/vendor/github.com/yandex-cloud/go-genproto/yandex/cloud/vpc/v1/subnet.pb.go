// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/vpc/v1/subnet.proto

package vpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A Subnet resource. For more information, see [Subnets](/docs/vpc/concepts/subnets).
type Subnet struct {
	// ID of the subnet.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the subnet belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the subnet. The name is unique within the project. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description of the subnet. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the network the subnet belongs to.
	NetworkId string `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// ID of the availability zone where the subnet resides.
	ZoneId string `protobuf:"bytes,8,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// CIDR block.
	// The range of internal addresses that are defined for this subnet.
	// This field can be set only at Subnet resource creation time and cannot be changed.
	// For example, 10.0.0.0/22 or 192.168.0.0/24.
	// Minimum subnet size is /28, maximum subnet size is /16.
	V4CidrBlocks []string `protobuf:"bytes,10,rep,name=v4_cidr_blocks,json=v4CidrBlocks,proto3" json:"v4_cidr_blocks,omitempty"`
	// IPv6 not available yet.
	V6CidrBlocks []string `protobuf:"bytes,11,rep,name=v6_cidr_blocks,json=v6CidrBlocks,proto3" json:"v6_cidr_blocks,omitempty"`
	// ID of route table the subnet is linked to.
	RouteTableId         string   `protobuf:"bytes,12,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}
func (*Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c0de762dc72cc6, []int{0}
}

func (m *Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet.Unmarshal(m, b)
}
func (m *Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet.Marshal(b, m, deterministic)
}
func (m *Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet.Merge(m, src)
}
func (m *Subnet) XXX_Size() int {
	return xxx_messageInfo_Subnet.Size(m)
}
func (m *Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet proto.InternalMessageInfo

func (m *Subnet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subnet) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Subnet) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Subnet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subnet) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Subnet) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Subnet) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *Subnet) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *Subnet) GetV4CidrBlocks() []string {
	if m != nil {
		return m.V4CidrBlocks
	}
	return nil
}

func (m *Subnet) GetV6CidrBlocks() []string {
	if m != nil {
		return m.V6CidrBlocks
	}
	return nil
}

func (m *Subnet) GetRouteTableId() string {
	if m != nil {
		return m.RouteTableId
	}
	return ""
}

func init() {
	proto.RegisterType((*Subnet)(nil), "yandex.cloud.vpc.v1.Subnet")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.Subnet.LabelsEntry")
}

func init() { proto.RegisterFile("yandex/cloud/vpc/v1/subnet.proto", fileDescriptor_40c0de762dc72cc6) }

var fileDescriptor_40c0de762dc72cc6 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x6b, 0xd4, 0x40,
	0x14, 0x85, 0xd9, 0x4d, 0x9b, 0x36, 0x37, 0x4b, 0x91, 0x51, 0x30, 0xac, 0x88, 0x41, 0x04, 0xf7,
	0xa5, 0x13, 0x5a, 0x4b, 0xb1, 0x16, 0x11, 0x2b, 0x3e, 0x2c, 0xf8, 0xb4, 0xf6, 0xc9, 0x97, 0x90,
	0xcc, 0xdc, 0xc6, 0x61, 0x93, 0x4c, 0x98, 0x4c, 0x46, 0xd7, 0xdf, 0xe8, 0x8f, 0x2a, 0xb9, 0x93,
	0x42, 0x0b, 0xfb, 0x36, 0x73, 0xce, 0x37, 0x73, 0x38, 0x73, 0x07, 0xd2, 0x5d, 0xd1, 0x4a, 0xfc,
	0x9b, 0x89, 0x5a, 0x0f, 0x32, 0x73, 0x9d, 0xc8, 0xdc, 0x59, 0xd6, 0x0f, 0x65, 0x8b, 0x96, 0x77,
	0x46, 0x5b, 0xcd, 0x9e, 0x7b, 0x82, 0x13, 0xc1, 0x5d, 0x27, 0xb8, 0x3b, 0x5b, 0xbe, 0xa9, 0xb4,
	0xae, 0x6a, 0xcc, 0x08, 0x29, 0x87, 0xbb, 0xcc, 0xaa, 0x06, 0x7b, 0x5b, 0x34, 0x9d, 0x3f, 0xf5,
	0xf6, 0x7f, 0x00, 0xe1, 0x4f, 0xba, 0x86, 0x9d, 0xc0, 0x5c, 0xc9, 0x64, 0x96, 0xce, 0x56, 0xd1,
	0x66, 0xae, 0x24, 0x7b, 0x05, 0xd1, 0x9d, 0xae, 0x25, 0x9a, 0x5c, 0xc9, 0x64, 0x4e, 0xf2, 0xb1,
	0x17, 0xd6, 0x92, 0x5d, 0x01, 0x08, 0x83, 0x85, 0x45, 0x99, 0x17, 0x36, 0x09, 0xd2, 0xd9, 0x2a,
	0x3e, 0x5f, 0x72, 0x9f, 0xc6, 0x1f, 0xd2, 0xf8, 0xed, 0x43, 0xda, 0x26, 0x9a, 0xe8, 0xaf, 0x96,
	0x31, 0x38, 0x68, 0x8b, 0x06, 0x93, 0x03, 0xba, 0x92, 0xd6, 0x2c, 0x85, 0x58, 0x62, 0x2f, 0x8c,
	0xea, 0xac, 0xd2, 0x6d, 0x72, 0x48, 0xd6, 0x63, 0x89, 0x7d, 0x81, 0xb0, 0x2e, 0x4a, 0xac, 0xfb,
	0x24, 0x4c, 0x83, 0x55, 0x7c, 0xfe, 0x9e, 0xef, 0xe9, 0xcb, 0x7d, 0x15, 0xfe, 0x83, 0xc8, 0xef,
	0xad, 0x35, 0xbb, 0xcd, 0x74, 0x8c, 0xbd, 0x06, 0x68, 0xd1, 0xfe, 0xd1, 0x66, 0x3b, 0xf6, 0x39,
	0xa2, 0x84, 0x68, 0x52, 0xd6, 0x92, 0xbd, 0x84, 0xa3, 0x7f, 0xba, 0xc5, 0xd1, 0x3b, 0x26, 0x2f,
	0x1c, 0xb7, 0x6b, 0xc9, 0xde, 0xc1, 0x89, 0xbb, 0xc8, 0x85, 0x92, 0x26, 0x2f, 0x6b, 0x2d, 0xb6,
	0x7d, 0x02, 0x69, 0xb0, 0x8a, 0x36, 0x0b, 0x77, 0xf1, 0x4d, 0x49, 0x73, 0x43, 0x1a, 0x51, 0x97,
	0x4f, 0xa8, 0x78, 0xa2, 0x2e, 0x9f, 0x52, 0x46, 0x0f, 0x16, 0x73, 0x5b, 0x94, 0x35, 0x65, 0x2d,
	0x28, 0x6b, 0x41, 0xea, 0xed, 0x28, 0xae, 0xe5, 0xf2, 0x0a, 0xe2, 0x47, 0x05, 0xd8, 0x33, 0x08,
	0xb6, 0xb8, 0x9b, 0x06, 0x33, 0x2e, 0xd9, 0x0b, 0x38, 0x74, 0x45, 0x3d, 0xe0, 0x34, 0x15, 0xbf,
	0xf9, 0x34, 0xff, 0x38, 0xbb, 0xf9, 0xfc, 0xeb, 0xba, 0x52, 0xf6, 0xf7, 0x50, 0x72, 0xa1, 0x9b,
	0xcc, 0xbf, 0xd0, 0xa9, 0xff, 0x33, 0x95, 0x3e, 0xad, 0xb0, 0xa5, 0xd1, 0x64, 0x7b, 0x3e, 0xd3,
	0xb5, 0xeb, 0x44, 0x19, 0x92, 0xfd, 0xe1, 0x3e, 0x00, 0x00, 0xff, 0xff, 0x42, 0x30, 0x95, 0x96,
	0x6e, 0x02, 0x00, 0x00,
}