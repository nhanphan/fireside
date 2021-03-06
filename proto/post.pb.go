// Code generated by protoc-gen-go.
// source: proto/post.proto
// DO NOT EDIT!

/*
Package post is a generated protocol buffer package.

It is generated from these files:
	proto/post.proto

It has these top-level messages:
	Post
	Posts
	GetByUUIDRequest
	SearchRequest
	Filter
	OrderBy
*/
package post

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PostField int32

const (
	PostField_UUID         PostField = 0
	PostField_CREATOR_UUID PostField = 1
	PostField_TYPE         PostField = 2
)

var PostField_name = map[int32]string{
	0: "UUID",
	1: "CREATOR_UUID",
	2: "TYPE",
}
var PostField_value = map[string]int32{
	"UUID":         0,
	"CREATOR_UUID": 1,
	"TYPE":         2,
}

func (x PostField) String() string {
	return proto.EnumName(PostField_name, int32(x))
}
func (PostField) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Operator int32

const (
	Operator_EQ        Operator = 0
	Operator_IEQ       Operator = 1
	Operator_NE        Operator = 2
	Operator_NIE       Operator = 3
	Operator_GT        Operator = 4
	Operator_GTE       Operator = 5
	Operator_LT        Operator = 6
	Operator_LTE       Operator = 7
	Operator_LIKE      Operator = 8
	Operator_ISNULL    Operator = 9
	Operator_ISNOTNULL Operator = 10
)

var Operator_name = map[int32]string{
	0:  "EQ",
	1:  "IEQ",
	2:  "NE",
	3:  "NIE",
	4:  "GT",
	5:  "GTE",
	6:  "LT",
	7:  "LTE",
	8:  "LIKE",
	9:  "ISNULL",
	10: "ISNOTNULL",
}
var Operator_value = map[string]int32{
	"EQ":        0,
	"IEQ":       1,
	"NE":        2,
	"NIE":       3,
	"GT":        4,
	"GTE":       5,
	"LT":        6,
	"LTE":       7,
	"LIKE":      8,
	"ISNULL":    9,
	"ISNOTNULL": 10,
}

func (x Operator) String() string {
	return proto.EnumName(Operator_name, int32(x))
}
func (Operator) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Condition int32

const (
	Condition_AND Condition = 0
	Condition_OR  Condition = 1
)

var Condition_name = map[int32]string{
	0: "AND",
	1: "OR",
}
var Condition_value = map[string]int32{
	"AND": 0,
	"OR":  1,
}

func (x Condition) String() string {
	return proto.EnumName(Condition_name, int32(x))
}
func (Condition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Post struct {
	Uuid        string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	CreatorUuid string `protobuf:"bytes,2,opt,name=creatorUuid" json:"creatorUuid,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Posts struct {
	Posts []*Post `protobuf:"bytes,1,rep,name=posts" json:"posts,omitempty"`
}

func (m *Posts) Reset()                    { *m = Posts{} }
func (m *Posts) String() string            { return proto.CompactTextString(m) }
func (*Posts) ProtoMessage()               {}
func (*Posts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Posts) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type GetByUUIDRequest struct {
	Uuid   string      `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Fields []PostField `protobuf:"varint,2,rep,name=fields,enum=post.PostField" json:"fields,omitempty"`
}

func (m *GetByUUIDRequest) Reset()                    { *m = GetByUUIDRequest{} }
func (m *GetByUUIDRequest) String() string            { return proto.CompactTextString(m) }
func (*GetByUUIDRequest) ProtoMessage()               {}
func (*GetByUUIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SearchRequest struct {
	Filters []*Filter   `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
	Fields  []PostField `protobuf:"varint,2,rep,name=fields,enum=post.PostField" json:"fields,omitempty"`
	Limit   int32       `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset  int32       `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	OrderBy *OrderBy    `protobuf:"bytes,5,opt,name=orderBy" json:"orderBy,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchRequest) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *SearchRequest) GetOrderBy() *OrderBy {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

type Filter struct {
	Field PostField `protobuf:"varint,1,opt,name=field,enum=post.PostField" json:"field,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Filter_StringVal
	//	*Filter_IntVal
	Value     isFilter_Value `protobuf_oneof:"value"`
	Operator  Operator       `protobuf:"varint,4,opt,name=operator,enum=post.Operator" json:"operator,omitempty"`
	Condition Condition      `protobuf:"varint,5,opt,name=condition,enum=post.Condition" json:"condition,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isFilter_Value interface {
	isFilter_Value()
}

type Filter_StringVal struct {
	StringVal string `protobuf:"bytes,2,opt,name=stringVal,oneof"`
}
type Filter_IntVal struct {
	IntVal int32 `protobuf:"varint,3,opt,name=intVal,oneof"`
}

func (*Filter_StringVal) isFilter_Value() {}
func (*Filter_IntVal) isFilter_Value()    {}

func (m *Filter) GetValue() isFilter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Filter) GetStringVal() string {
	if x, ok := m.GetValue().(*Filter_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (m *Filter) GetIntVal() int32 {
	if x, ok := m.GetValue().(*Filter_IntVal); ok {
		return x.IntVal
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Filter_OneofMarshaler, _Filter_OneofUnmarshaler, _Filter_OneofSizer, []interface{}{
		(*Filter_StringVal)(nil),
		(*Filter_IntVal)(nil),
	}
}

func _Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Filter)
	// value
	switch x := m.Value.(type) {
	case *Filter_StringVal:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringVal)
	case *Filter_IntVal:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntVal))
	case nil:
	default:
		return fmt.Errorf("Filter.Value has unexpected type %T", x)
	}
	return nil
}

func _Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Filter)
	switch tag {
	case 2: // value.stringVal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Filter_StringVal{x}
		return true, err
	case 3: // value.intVal
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Filter_IntVal{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Filter)
	// value
	switch x := m.Value.(type) {
	case *Filter_StringVal:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringVal)))
		n += len(x.StringVal)
	case *Filter_IntVal:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntVal))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type OrderBy struct {
	Field        PostField `protobuf:"varint,1,opt,name=field,enum=post.PostField" json:"field,omitempty"`
	IsDescending bool      `protobuf:"varint,2,opt,name=isDescending" json:"isDescending,omitempty"`
}

func (m *OrderBy) Reset()                    { *m = OrderBy{} }
func (m *OrderBy) String() string            { return proto.CompactTextString(m) }
func (*OrderBy) ProtoMessage()               {}
func (*OrderBy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Post)(nil), "post.Post")
	proto.RegisterType((*Posts)(nil), "post.Posts")
	proto.RegisterType((*GetByUUIDRequest)(nil), "post.GetByUUIDRequest")
	proto.RegisterType((*SearchRequest)(nil), "post.SearchRequest")
	proto.RegisterType((*Filter)(nil), "post.Filter")
	proto.RegisterType((*OrderBy)(nil), "post.OrderBy")
	proto.RegisterEnum("post.PostField", PostField_name, PostField_value)
	proto.RegisterEnum("post.Operator", Operator_name, Operator_value)
	proto.RegisterEnum("post.Condition", Condition_name, Condition_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PostSvc service

type PostSvcClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Posts, error)
	GetByUUID(ctx context.Context, in *GetByUUIDRequest, opts ...grpc.CallOption) (*Post, error)
}

type postSvcClient struct {
	cc *grpc.ClientConn
}

func NewPostSvcClient(cc *grpc.ClientConn) PostSvcClient {
	return &postSvcClient{cc}
}

func (c *postSvcClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Posts, error) {
	out := new(Posts)
	err := grpc.Invoke(ctx, "/post.PostSvc/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postSvcClient) GetByUUID(ctx context.Context, in *GetByUUIDRequest, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := grpc.Invoke(ctx, "/post.PostSvc/GetByUUID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PostSvc service

type PostSvcServer interface {
	Search(context.Context, *SearchRequest) (*Posts, error)
	GetByUUID(context.Context, *GetByUUIDRequest) (*Post, error)
}

func RegisterPostSvcServer(s *grpc.Server, srv PostSvcServer) {
	s.RegisterService(&_PostSvc_serviceDesc, srv)
}

func _PostSvc_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostSvcServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostSvc/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostSvcServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostSvc_GetByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostSvcServer).GetByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostSvc/GetByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostSvcServer).GetByUUID(ctx, req.(*GetByUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostSvc",
	HandlerType: (*PostSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _PostSvc_Search_Handler,
		},
		{
			MethodName: "GetByUUID",
			Handler:    _PostSvc_GetByUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/post.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x93, 0xd8, 0x8e, 0x27, 0xfd, 0x59, 0x96, 0xaa, 0x8a, 0xa2, 0x0a, 0x55, 0x96, 0x28,
	0x25, 0x12, 0x8d, 0x1a, 0xde, 0x78, 0xeb, 0x4f, 0x5a, 0x22, 0xa2, 0xa4, 0x6c, 0x1d, 0x10, 0xbc,
	0x20, 0x93, 0x6c, 0x82, 0xa5, 0xe0, 0x0d, 0xde, 0x4d, 0xa5, 0x08, 0xf1, 0xc2, 0x15, 0xb8, 0x05,
	0x17, 0xe0, 0x14, 0x3c, 0x71, 0x05, 0x0e, 0xc2, 0xce, 0x6e, 0x9c, 0xb4, 0x08, 0xa4, 0xbe, 0xcd,
	0x7e, 0xdf, 0xcc, 0xe7, 0x6f, 0x76, 0x67, 0x0c, 0x64, 0x9a, 0x09, 0x25, 0x1a, 0x53, 0x21, 0xd5,
	0xa1, 0x09, 0x69, 0x09, 0xe3, 0xda, 0xee, 0x58, 0x88, 0xf1, 0x84, 0x37, 0xe2, 0x69, 0xd2, 0x88,
	0xd3, 0x54, 0xa8, 0x58, 0x25, 0x22, 0x95, 0x36, 0x27, 0xbc, 0x84, 0xd2, 0xa5, 0xce, 0xa2, 0x14,
	0x4a, 0xb3, 0x59, 0x32, 0xac, 0x3a, 0x7b, 0xce, 0x41, 0xc0, 0x4c, 0x4c, 0xf7, 0xa0, 0x32, 0xc8,
	0x78, 0xac, 0x44, 0xd6, 0x47, 0xaa, 0x60, 0xa8, 0x9b, 0x10, 0x56, 0xa9, 0xf9, 0x94, 0x57, 0x8b,
	0xb6, 0x0a, 0xe3, 0xf0, 0x31, 0xb8, 0xa8, 0x28, 0x75, 0xb9, 0x8b, 0x06, 0xa4, 0xd6, 0x2c, 0x1e,
	0x54, 0x9a, 0x70, 0x68, 0xac, 0x21, 0xc7, 0x2c, 0x11, 0xf6, 0x80, 0x5c, 0x70, 0x75, 0x32, 0xef,
	0xf7, 0xdb, 0x67, 0x8c, 0x7f, 0x9a, 0xf1, 0xff, 0x18, 0x79, 0x04, 0xde, 0x28, 0xe1, 0x93, 0xa1,
	0xd4, 0x1e, 0x8a, 0x07, 0x9b, 0xcd, 0xad, 0x95, 0xd4, 0x39, 0xe2, 0x6c, 0x41, 0x87, 0x3f, 0x1c,
	0xd8, 0xb8, 0xe2, 0x71, 0x36, 0xf8, 0x90, 0xcb, 0xed, 0x83, 0x3f, 0x4a, 0x26, 0x8a, 0x67, 0xb9,
	0x8d, 0x75, 0x5b, 0x7b, 0x6e, 0x40, 0x96, 0x93, 0x77, 0xfe, 0x04, 0xdd, 0x06, 0x77, 0x92, 0x7c,
	0x4c, 0x94, 0xe9, 0xd9, 0x65, 0xf6, 0x40, 0x77, 0xc0, 0x13, 0xa3, 0x91, 0xe4, 0xaa, 0x5a, 0x32,
	0xf0, 0xe2, 0xa4, 0x65, 0x7d, 0x91, 0x0d, 0x79, 0x76, 0x32, 0xaf, 0xba, 0x9a, 0xa8, 0x34, 0x37,
	0xac, 0x6e, 0xcf, 0x82, 0x2c, 0x67, 0xc3, 0x9f, 0x0e, 0x78, 0xd6, 0x13, 0x7d, 0x08, 0xae, 0xf9,
	0x96, 0xb9, 0x82, 0x7f, 0x38, 0xb1, 0x2c, 0x7d, 0x00, 0x81, 0x54, 0x59, 0x92, 0x8e, 0x5f, 0xc5,
	0x13, 0xfb, 0x36, 0xcf, 0xd7, 0xd8, 0x0a, 0xa2, 0x55, 0xf0, 0x92, 0x54, 0x21, 0x69, 0x9c, 0x6a,
	0x72, 0x71, 0xa6, 0x75, 0x28, 0x8b, 0x29, 0xcf, 0xf0, 0x15, 0x8d, 0xdd, 0xcd, 0xe6, 0xe6, 0xc2,
	0xd5, 0x02, 0x65, 0x4b, 0x9e, 0x3e, 0x81, 0x60, 0x20, 0xd2, 0x61, 0x82, 0x33, 0x63, 0x5a, 0x58,
	0x1a, 0x3a, 0xcd, 0x61, 0xb6, 0xca, 0x38, 0xf1, 0xc1, 0xbd, 0x8e, 0x27, 0x33, 0x1e, 0x46, 0xe0,
	0x2f, 0x7a, 0xbc, 0x6b, 0x3f, 0x21, 0xac, 0x27, 0xf2, 0x8c, 0xcb, 0x01, 0xd7, 0x62, 0xe9, 0xd8,
	0xb4, 0x54, 0x66, 0xb7, 0xb0, 0xfa, 0x11, 0x04, 0xcb, 0x3a, 0x5a, 0x86, 0x12, 0x0e, 0x0e, 0x59,
	0xa3, 0x04, 0xd6, 0x4f, 0x59, 0xeb, 0x38, 0xea, 0xb1, 0x77, 0x06, 0x71, 0x90, 0x8b, 0xde, 0x5c,
	0xb6, 0x48, 0xa1, 0x2e, 0xa1, 0x9c, 0xb7, 0x45, 0x3d, 0x28, 0xb4, 0x5e, 0xea, 0x7c, 0x1f, 0x8a,
	0x6d, 0x1d, 0x38, 0x08, 0x74, 0x75, 0x12, 0x02, 0xdd, 0x76, 0x8b, 0x14, 0x11, 0xb8, 0x88, 0x48,
	0x09, 0x81, 0x8b, 0xa8, 0x45, 0x5c, 0x04, 0x3a, 0x11, 0xf1, 0x10, 0xe8, 0x68, 0xc0, 0x47, 0xe5,
	0x4e, 0xfb, 0x45, 0x8b, 0x94, 0x29, 0x80, 0xd7, 0xbe, 0xea, 0xf6, 0x3b, 0x1d, 0x12, 0xd0, 0x0d,
	0x08, 0x74, 0xdc, 0x8b, 0xcc, 0x11, 0xea, 0xbb, 0x10, 0x2c, 0xaf, 0x07, 0x4b, 0x8f, 0xbb, 0x68,
	0x53, 0x6b, 0xf5, 0x18, 0x71, 0x9a, 0xdf, 0x1d, 0xf0, 0xb1, 0x8d, 0xab, 0xeb, 0x01, 0x3d, 0x07,
	0xcf, 0x0e, 0x2c, 0xbd, 0x6f, 0xef, 0xe5, 0xd6, 0xf8, 0xd6, 0x2a, 0xab, 0xcb, 0x92, 0x61, 0xed,
	0xeb, 0xaf, 0xdf, 0xdf, 0x0a, 0xdb, 0xe1, 0x56, 0xe3, 0xfa, 0xc8, 0xec, 0x79, 0x43, 0x9a, 0xe4,
	0x67, 0x4e, 0x9d, 0xbe, 0x86, 0x60, 0xb9, 0x4a, 0x74, 0xc7, 0x56, 0xfd, 0xbd, 0x5b, 0xb5, 0x1b,
	0x2b, 0x18, 0xee, 0x1b, 0xb1, 0xbd, 0xb0, 0x9c, 0x8b, 0x69, 0x95, 0xb7, 0xf7, 0xe8, 0x4a, 0xfb,
	0x33, 0xae, 0xde, 0x97, 0xf7, 0x9e, 0xf9, 0x4f, 0x3c, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb0,
	0x26, 0xc7, 0xaf, 0x5f, 0x04, 0x00, 0x00,
}
