// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/vpc/v1/route_table.proto

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

// A RouteTable resource. For more information, see [RouteTables](/docs/vpc/concepts/route_tables).
type RouteTable struct {
	// ID of the route table.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the route table belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the route table. The name is unique within the project. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description of the route table. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the network the route table belongs to.
	NetworkId string `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// List of static routes.
	StaticRoutes         []*StaticRoute `protobuf:"bytes,8,rep,name=static_routes,json=staticRoutes,proto3" json:"static_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RouteTable) Reset()         { *m = RouteTable{} }
func (m *RouteTable) String() string { return proto.CompactTextString(m) }
func (*RouteTable) ProtoMessage()    {}
func (*RouteTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_fefcab5534a5cf97, []int{0}
}

func (m *RouteTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTable.Unmarshal(m, b)
}
func (m *RouteTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTable.Marshal(b, m, deterministic)
}
func (m *RouteTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTable.Merge(m, src)
}
func (m *RouteTable) XXX_Size() int {
	return xxx_messageInfo_RouteTable.Size(m)
}
func (m *RouteTable) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTable.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTable proto.InternalMessageInfo

func (m *RouteTable) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RouteTable) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *RouteTable) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *RouteTable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteTable) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RouteTable) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *RouteTable) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *RouteTable) GetStaticRoutes() []*StaticRoute {
	if m != nil {
		return m.StaticRoutes
	}
	return nil
}

// A StaticRoute resource. For more information, see [StaticRoutes](/docs/vpc/concepts/static_routes).
type StaticRoute struct {
	// Types that are valid to be assigned to Destination:
	//	*StaticRoute_DestinationPrefix
	Destination isStaticRoute_Destination `protobuf_oneof:"destination"`
	// Types that are valid to be assigned to NextHop:
	//	*StaticRoute_NextHopAddress
	NextHop isStaticRoute_NextHop `protobuf_oneof:"next_hop"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StaticRoute) Reset()         { *m = StaticRoute{} }
func (m *StaticRoute) String() string { return proto.CompactTextString(m) }
func (*StaticRoute) ProtoMessage()    {}
func (*StaticRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_fefcab5534a5cf97, []int{1}
}

func (m *StaticRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticRoute.Unmarshal(m, b)
}
func (m *StaticRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticRoute.Marshal(b, m, deterministic)
}
func (m *StaticRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticRoute.Merge(m, src)
}
func (m *StaticRoute) XXX_Size() int {
	return xxx_messageInfo_StaticRoute.Size(m)
}
func (m *StaticRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticRoute.DiscardUnknown(m)
}

var xxx_messageInfo_StaticRoute proto.InternalMessageInfo

type isStaticRoute_Destination interface {
	isStaticRoute_Destination()
}

type StaticRoute_DestinationPrefix struct {
	DestinationPrefix string `protobuf:"bytes,1,opt,name=destination_prefix,json=destinationPrefix,proto3,oneof"`
}

func (*StaticRoute_DestinationPrefix) isStaticRoute_Destination() {}

func (m *StaticRoute) GetDestination() isStaticRoute_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *StaticRoute) GetDestinationPrefix() string {
	if x, ok := m.GetDestination().(*StaticRoute_DestinationPrefix); ok {
		return x.DestinationPrefix
	}
	return ""
}

type isStaticRoute_NextHop interface {
	isStaticRoute_NextHop()
}

type StaticRoute_NextHopAddress struct {
	NextHopAddress string `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3,oneof"`
}

func (*StaticRoute_NextHopAddress) isStaticRoute_NextHop() {}

func (m *StaticRoute) GetNextHop() isStaticRoute_NextHop {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *StaticRoute) GetNextHopAddress() string {
	if x, ok := m.GetNextHop().(*StaticRoute_NextHopAddress); ok {
		return x.NextHopAddress
	}
	return ""
}

func (m *StaticRoute) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StaticRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StaticRoute_DestinationPrefix)(nil),
		(*StaticRoute_NextHopAddress)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteTable)(nil), "yandex.cloud.vpc.v1.RouteTable")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.RouteTable.LabelsEntry")
	proto.RegisterType((*StaticRoute)(nil), "yandex.cloud.vpc.v1.StaticRoute")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.StaticRoute.LabelsEntry")
}

func init() {
	proto.RegisterFile("yandex/cloud/vpc/v1/route_table.proto", fileDescriptor_fefcab5534a5cf97)
}

var fileDescriptor_fefcab5534a5cf97 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x25, 0xc9, 0x36, 0xda, 0x2f, 0x6c, 0x1a, 0x86, 0x43, 0x54, 0x84, 0x88, 0x26, 0x21, 0x55,
	0xc0, 0x1c, 0x6d, 0x5c, 0x18, 0x13, 0x87, 0x15, 0x26, 0x75, 0x12, 0x07, 0x14, 0x76, 0xe2, 0x12,
	0x39, 0xf1, 0xd7, 0xcc, 0x5a, 0x1a, 0x47, 0xb6, 0x13, 0xda, 0x3f, 0xc0, 0x99, 0x9f, 0x8c, 0x62,
	0xa7, 0xac, 0x87, 0x4a, 0x1c, 0xb8, 0xd9, 0xdf, 0x7b, 0x7e, 0xcf, 0xef, 0x25, 0x86, 0xd7, 0x6b,
	0x56, 0x73, 0x5c, 0x25, 0x45, 0x25, 0x5b, 0x9e, 0x74, 0x4d, 0x91, 0x74, 0x67, 0x89, 0x92, 0xad,
	0xc1, 0xcc, 0xb0, 0xbc, 0x42, 0xda, 0x28, 0x69, 0x24, 0x79, 0xe6, 0x68, 0xd4, 0xd2, 0x68, 0xd7,
	0x14, 0xb4, 0x3b, 0x9b, 0xbc, 0x2a, 0xa5, 0x2c, 0x2b, 0x4c, 0x2c, 0x25, 0x6f, 0x17, 0x89, 0x11,
	0x4b, 0xd4, 0x86, 0x2d, 0x1b, 0x77, 0xea, 0xe4, 0x77, 0x00, 0x90, 0xf6, 0x5a, 0xb7, 0xbd, 0x14,
	0x39, 0x02, 0x5f, 0xf0, 0xc8, 0x8b, 0xbd, 0xe9, 0x38, 0xf5, 0x05, 0x27, 0x2f, 0x60, 0xbc, 0x90,
	0x15, 0x47, 0x95, 0x09, 0x1e, 0xf9, 0x76, 0x3c, 0x72, 0x83, 0x1b, 0x4e, 0x2e, 0x00, 0x0a, 0x85,
	0xcc, 0x20, 0xcf, 0x98, 0x89, 0x82, 0xd8, 0x9b, 0x86, 0xe7, 0x13, 0xea, 0x1c, 0xe9, 0xc6, 0x91,
	0xde, 0x6e, 0x1c, 0xd3, 0xf1, 0xc0, 0xbe, 0x32, 0x84, 0xc0, 0x5e, 0xcd, 0x96, 0x18, 0xed, 0x59,
	0x49, 0xbb, 0x26, 0x31, 0x84, 0x1c, 0x75, 0xa1, 0x44, 0x63, 0x84, 0xac, 0xa3, 0x7d, 0x0b, 0x6d,
	0x8f, 0xc8, 0x67, 0x38, 0xa8, 0x58, 0x8e, 0x95, 0x8e, 0x0e, 0xe2, 0x60, 0x1a, 0x9e, 0xbf, 0xa5,
	0x3b, 0x32, 0xd3, 0x87, 0x38, 0xf4, 0xab, 0x65, 0x5f, 0xd7, 0x46, 0xad, 0xd3, 0xe1, 0x28, 0x79,
	0x09, 0x50, 0xa3, 0xf9, 0x29, 0xd5, 0x7d, 0x9f, 0xe9, 0xb1, 0x75, 0x19, 0x0f, 0x93, 0x1b, 0x4e,
	0xae, 0xe1, 0x50, 0x1b, 0x66, 0x44, 0x91, 0xd9, 0x8a, 0x75, 0x34, 0xb2, 0x56, 0xf1, 0x4e, 0xab,
	0xef, 0x96, 0x69, 0x0d, 0xd3, 0x27, 0xfa, 0x61, 0xa3, 0x27, 0x17, 0x10, 0x6e, 0x99, 0x93, 0x63,
	0x08, 0xee, 0x71, 0x3d, 0x14, 0xdb, 0x2f, 0xc9, 0x73, 0xd8, 0xef, 0x58, 0xd5, 0xe2, 0xd0, 0xaa,
	0xdb, 0x7c, 0xf4, 0x3f, 0x78, 0x27, 0xbf, 0x7c, 0x08, 0xb7, 0x84, 0x49, 0x02, 0x84, 0xa3, 0x36,
	0xa2, 0x66, 0x7d, 0x09, 0x59, 0xa3, 0x70, 0x21, 0x56, 0x4e, 0x6a, 0xfe, 0x28, 0x7d, 0xba, 0x85,
	0x7d, 0xb3, 0x10, 0x79, 0x03, 0xc7, 0x35, 0xae, 0x4c, 0x76, 0x27, 0x9b, 0x8c, 0x71, 0xae, 0x50,
	0x6b, 0xe7, 0x32, 0xf7, 0xd2, 0xa3, 0x1e, 0x99, 0xcb, 0xe6, 0xca, 0xcd, 0xc9, 0x97, 0xbf, 0x95,
	0x06, 0x36, 0xe7, 0xbb, 0x7f, 0xe5, 0xdc, 0xd5, 0xe9, 0x7f, 0xa4, 0x9d, 0x1d, 0xda, 0xaf, 0xbe,
	0x49, 0x30, 0x03, 0x18, 0x6d, 0xee, 0x3e, 0xfb, 0xf4, 0xe3, 0xb2, 0x14, 0xe6, 0xae, 0xcd, 0x69,
	0x21, 0x97, 0x89, 0xbb, 0xd7, 0xa9, 0x7b, 0x05, 0xa5, 0x3c, 0x2d, 0xb1, 0xb6, 0xff, 0x58, 0xb2,
	0xe3, 0x79, 0x5c, 0x76, 0x4d, 0x91, 0x1f, 0x58, 0xf8, 0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x89, 0xe7, 0xdf, 0x40, 0x03, 0x00, 0x00,
}