// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opinion/V1/opinion.proto

package opinion

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

type Opinion struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Asunto               string   `protobuf:"bytes,2,opt,name=asunto,proto3" json:"asunto,omitempty"`
	Autor                string   `protobuf:"bytes,3,opt,name=autor,proto3" json:"autor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Opinion) Reset()         { *m = Opinion{} }
func (m *Opinion) String() string { return proto.CompactTextString(m) }
func (*Opinion) ProtoMessage()    {}
func (*Opinion) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{0}
}

func (m *Opinion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Opinion.Unmarshal(m, b)
}
func (m *Opinion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Opinion.Marshal(b, m, deterministic)
}
func (m *Opinion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Opinion.Merge(m, src)
}
func (m *Opinion) XXX_Size() int {
	return xxx_messageInfo_Opinion.Size(m)
}
func (m *Opinion) XXX_DiscardUnknown() {
	xxx_messageInfo_Opinion.DiscardUnknown(m)
}

var xxx_messageInfo_Opinion proto.InternalMessageInfo

func (m *Opinion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Opinion) GetAsunto() string {
	if m != nil {
		return m.Asunto
	}
	return ""
}

func (m *Opinion) GetAutor() string {
	if m != nil {
		return m.Autor
	}
	return ""
}

// -- Get Opinion Messages types
type GetOpinionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOpinionRequest) Reset()         { *m = GetOpinionRequest{} }
func (m *GetOpinionRequest) String() string { return proto.CompactTextString(m) }
func (*GetOpinionRequest) ProtoMessage()    {}
func (*GetOpinionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{1}
}

func (m *GetOpinionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOpinionRequest.Unmarshal(m, b)
}
func (m *GetOpinionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOpinionRequest.Marshal(b, m, deterministic)
}
func (m *GetOpinionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOpinionRequest.Merge(m, src)
}
func (m *GetOpinionRequest) XXX_Size() int {
	return xxx_messageInfo_GetOpinionRequest.Size(m)
}
func (m *GetOpinionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOpinionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOpinionRequest proto.InternalMessageInfo

func (m *GetOpinionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetOpinionResponse struct {
	Opinion              *Opinion `protobuf:"bytes,1,opt,name=opinion,proto3" json:"opinion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOpinionResponse) Reset()         { *m = GetOpinionResponse{} }
func (m *GetOpinionResponse) String() string { return proto.CompactTextString(m) }
func (*GetOpinionResponse) ProtoMessage()    {}
func (*GetOpinionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{2}
}

func (m *GetOpinionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOpinionResponse.Unmarshal(m, b)
}
func (m *GetOpinionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOpinionResponse.Marshal(b, m, deterministic)
}
func (m *GetOpinionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOpinionResponse.Merge(m, src)
}
func (m *GetOpinionResponse) XXX_Size() int {
	return xxx_messageInfo_GetOpinionResponse.Size(m)
}
func (m *GetOpinionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOpinionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOpinionResponse proto.InternalMessageInfo

func (m *GetOpinionResponse) GetOpinion() *Opinion {
	if m != nil {
		return m.Opinion
	}
	return nil
}

// -- Add Opinion Message types
type AddOpinionRequest struct {
	Opinion              *Opinion `protobuf:"bytes,1,opt,name=opinion,proto3" json:"opinion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddOpinionRequest) Reset()         { *m = AddOpinionRequest{} }
func (m *AddOpinionRequest) String() string { return proto.CompactTextString(m) }
func (*AddOpinionRequest) ProtoMessage()    {}
func (*AddOpinionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{3}
}

func (m *AddOpinionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddOpinionRequest.Unmarshal(m, b)
}
func (m *AddOpinionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddOpinionRequest.Marshal(b, m, deterministic)
}
func (m *AddOpinionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddOpinionRequest.Merge(m, src)
}
func (m *AddOpinionRequest) XXX_Size() int {
	return xxx_messageInfo_AddOpinionRequest.Size(m)
}
func (m *AddOpinionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddOpinionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddOpinionRequest proto.InternalMessageInfo

func (m *AddOpinionRequest) GetOpinion() *Opinion {
	if m != nil {
		return m.Opinion
	}
	return nil
}

type AddOpinionResponse struct {
	Opinion              *Opinion `protobuf:"bytes,1,opt,name=opinion,proto3" json:"opinion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddOpinionResponse) Reset()         { *m = AddOpinionResponse{} }
func (m *AddOpinionResponse) String() string { return proto.CompactTextString(m) }
func (*AddOpinionResponse) ProtoMessage()    {}
func (*AddOpinionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{4}
}

func (m *AddOpinionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddOpinionResponse.Unmarshal(m, b)
}
func (m *AddOpinionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddOpinionResponse.Marshal(b, m, deterministic)
}
func (m *AddOpinionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddOpinionResponse.Merge(m, src)
}
func (m *AddOpinionResponse) XXX_Size() int {
	return xxx_messageInfo_AddOpinionResponse.Size(m)
}
func (m *AddOpinionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddOpinionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddOpinionResponse proto.InternalMessageInfo

func (m *AddOpinionResponse) GetOpinion() *Opinion {
	if m != nil {
		return m.Opinion
	}
	return nil
}

// -- Delete Opinion Message types
type DeleteOpinionRequest struct {
	Opinion              *Opinion `protobuf:"bytes,1,opt,name=opinion,proto3" json:"opinion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOpinionRequest) Reset()         { *m = DeleteOpinionRequest{} }
func (m *DeleteOpinionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOpinionRequest) ProtoMessage()    {}
func (*DeleteOpinionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{5}
}

func (m *DeleteOpinionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOpinionRequest.Unmarshal(m, b)
}
func (m *DeleteOpinionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOpinionRequest.Marshal(b, m, deterministic)
}
func (m *DeleteOpinionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOpinionRequest.Merge(m, src)
}
func (m *DeleteOpinionRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOpinionRequest.Size(m)
}
func (m *DeleteOpinionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOpinionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOpinionRequest proto.InternalMessageInfo

func (m *DeleteOpinionRequest) GetOpinion() *Opinion {
	if m != nil {
		return m.Opinion
	}
	return nil
}

type DeleteOpinionResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOpinionResponse) Reset()         { *m = DeleteOpinionResponse{} }
func (m *DeleteOpinionResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteOpinionResponse) ProtoMessage()    {}
func (*DeleteOpinionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e78773aa03024be2, []int{6}
}

func (m *DeleteOpinionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOpinionResponse.Unmarshal(m, b)
}
func (m *DeleteOpinionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOpinionResponse.Marshal(b, m, deterministic)
}
func (m *DeleteOpinionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOpinionResponse.Merge(m, src)
}
func (m *DeleteOpinionResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteOpinionResponse.Size(m)
}
func (m *DeleteOpinionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOpinionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOpinionResponse proto.InternalMessageInfo

func (m *DeleteOpinionResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Opinion)(nil), "opinion.Opinion")
	proto.RegisterType((*GetOpinionRequest)(nil), "opinion.GetOpinionRequest")
	proto.RegisterType((*GetOpinionResponse)(nil), "opinion.GetOpinionResponse")
	proto.RegisterType((*AddOpinionRequest)(nil), "opinion.AddOpinionRequest")
	proto.RegisterType((*AddOpinionResponse)(nil), "opinion.AddOpinionResponse")
	proto.RegisterType((*DeleteOpinionRequest)(nil), "opinion.DeleteOpinionRequest")
	proto.RegisterType((*DeleteOpinionResponse)(nil), "opinion.DeleteOpinionResponse")
}

func init() { proto.RegisterFile("opinion/V1/opinion.proto", fileDescriptor_e78773aa03024be2) }

var fileDescriptor_e78773aa03024be2 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0x84, 0x50,
	0x14, 0xc5, 0xd1, 0xc8, 0xa1, 0x13, 0xfd, 0x99, 0xcb, 0x34, 0x88, 0x51, 0x84, 0x6d, 0xa2, 0xc5,
	0x48, 0xd3, 0x07, 0x68, 0x8a, 0x86, 0xd9, 0x15, 0x4c, 0xd0, 0xde, 0xf2, 0x2d, 0x1e, 0x84, 0xcf,
	0x7c, 0xd7, 0xbe, 0x76, 0x5f, 0x21, 0xd0, 0xab, 0x63, 0x6a, 0x8b, 0x6a, 0xf7, 0xce, 0xbb, 0x87,
	0xdf, 0x3d, 0x1e, 0x1f, 0x7c, 0x93, 0xe9, 0x54, 0x9b, 0x34, 0x7a, 0xbe, 0x8a, 0xe4, 0x38, 0xcb,
	0x72, 0xc3, 0x86, 0x46, 0x22, 0xc3, 0x15, 0x46, 0x8f, 0xd5, 0x91, 0xf6, 0xe1, 0xea, 0xc4, 0x77,
	0xce, 0x9c, 0x8b, 0x9d, 0xb5, 0xab, 0x13, 0x9a, 0xc2, 0x8b, 0x6d, 0x91, 0xb2, 0xf1, 0xdd, 0xf2,
	0x4e, 0x14, 0x4d, 0xb0, 0x1d, 0x17, 0x6c, 0x72, 0x7f, 0xab, 0xbc, 0xae, 0x44, 0x78, 0x8e, 0xf1,
	0x4a, 0xb1, 0xb0, 0xd6, 0xea, 0xbd, 0x50, 0x96, 0xbb, 0xc8, 0x70, 0x01, 0x6a, 0x9b, 0x6c, 0x66,
	0x52, 0xab, 0xe8, 0x12, 0x75, 0x9c, 0xd2, 0xba, 0x3b, 0x3f, 0x9c, 0xd5, 0x69, 0x6b, 0x6b, 0x93,
	0xf7, 0x06, 0xe3, 0xdb, 0x24, 0xe9, 0xac, 0xf9, 0x0d, 0x60, 0x01, 0x6a, 0x03, 0xfe, 0x10, 0xe1,
	0x0e, 0x93, 0x7b, 0xf5, 0xa6, 0x58, 0xfd, 0x23, 0x45, 0x84, 0xa3, 0x0e, 0x43, 0x82, 0x4c, 0xe1,
	0x59, 0x8e, 0xb9, 0xb0, 0xd2, 0x9a, 0xa8, 0xf9, 0xa7, 0x83, 0x03, 0xf1, 0x3e, 0xa9, 0xfc, 0x43,
	0xbf, 0x6a, 0x43, 0x4b, 0x60, 0xd3, 0x26, 0x05, 0xcd, 0xb6, 0xde, 0x7f, 0x08, 0x8e, 0x07, 0x67,
	0xb2, 0x72, 0x09, 0x6c, 0x1a, 0x69, 0x61, 0x7a, 0x3d, 0xb7, 0x30, 0x03, 0x15, 0x3e, 0x60, 0xef,
	0xdb, 0x27, 0xd1, 0x49, 0xe3, 0x1e, 0xaa, 0x2b, 0x38, 0xfd, 0x69, 0x5c, 0xf1, 0x5e, 0xbc, 0xf2,
	0xa5, 0x5e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xda, 0xc5, 0xfa, 0xc7, 0xc5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OpinionServicioClient is the client API for OpinionServicio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OpinionServicioClient interface {
	GetOpinion(ctx context.Context, in *GetOpinionRequest, opts ...grpc.CallOption) (*GetOpinionResponse, error)
	AddOpinion(ctx context.Context, in *AddOpinionRequest, opts ...grpc.CallOption) (*AddOpinionResponse, error)
	DeleteOpinion(ctx context.Context, in *DeleteOpinionRequest, opts ...grpc.CallOption) (*DeleteOpinionResponse, error)
}

type opinionServicioClient struct {
	cc *grpc.ClientConn
}

func NewOpinionServicioClient(cc *grpc.ClientConn) OpinionServicioClient {
	return &opinionServicioClient{cc}
}

func (c *opinionServicioClient) GetOpinion(ctx context.Context, in *GetOpinionRequest, opts ...grpc.CallOption) (*GetOpinionResponse, error) {
	out := new(GetOpinionResponse)
	err := c.cc.Invoke(ctx, "/opinion.OpinionServicio/GetOpinion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opinionServicioClient) AddOpinion(ctx context.Context, in *AddOpinionRequest, opts ...grpc.CallOption) (*AddOpinionResponse, error) {
	out := new(AddOpinionResponse)
	err := c.cc.Invoke(ctx, "/opinion.OpinionServicio/AddOpinion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opinionServicioClient) DeleteOpinion(ctx context.Context, in *DeleteOpinionRequest, opts ...grpc.CallOption) (*DeleteOpinionResponse, error) {
	out := new(DeleteOpinionResponse)
	err := c.cc.Invoke(ctx, "/opinion.OpinionServicio/DeleteOpinion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpinionServicioServer is the server API for OpinionServicio service.
type OpinionServicioServer interface {
	GetOpinion(context.Context, *GetOpinionRequest) (*GetOpinionResponse, error)
	AddOpinion(context.Context, *AddOpinionRequest) (*AddOpinionResponse, error)
	DeleteOpinion(context.Context, *DeleteOpinionRequest) (*DeleteOpinionResponse, error)
}

// UnimplementedOpinionServicioServer can be embedded to have forward compatible implementations.
type UnimplementedOpinionServicioServer struct {
}

func (*UnimplementedOpinionServicioServer) GetOpinion(ctx context.Context, req *GetOpinionRequest) (*GetOpinionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpinion not implemented")
}
func (*UnimplementedOpinionServicioServer) AddOpinion(ctx context.Context, req *AddOpinionRequest) (*AddOpinionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOpinion not implemented")
}
func (*UnimplementedOpinionServicioServer) DeleteOpinion(ctx context.Context, req *DeleteOpinionRequest) (*DeleteOpinionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOpinion not implemented")
}

func RegisterOpinionServicioServer(s *grpc.Server, srv OpinionServicioServer) {
	s.RegisterService(&_OpinionServicio_serviceDesc, srv)
}

func _OpinionServicio_GetOpinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOpinionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpinionServicioServer).GetOpinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opinion.OpinionServicio/GetOpinion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpinionServicioServer).GetOpinion(ctx, req.(*GetOpinionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpinionServicio_AddOpinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOpinionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpinionServicioServer).AddOpinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opinion.OpinionServicio/AddOpinion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpinionServicioServer).AddOpinion(ctx, req.(*AddOpinionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpinionServicio_DeleteOpinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOpinionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpinionServicioServer).DeleteOpinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opinion.OpinionServicio/DeleteOpinion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpinionServicioServer).DeleteOpinion(ctx, req.(*DeleteOpinionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OpinionServicio_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opinion.OpinionServicio",
	HandlerType: (*OpinionServicioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOpinion",
			Handler:    _OpinionServicio_GetOpinion_Handler,
		},
		{
			MethodName: "AddOpinion",
			Handler:    _OpinionServicio_AddOpinion_Handler,
		},
		{
			MethodName: "DeleteOpinion",
			Handler:    _OpinionServicio_DeleteOpinion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opinion/V1/opinion.proto",
}
