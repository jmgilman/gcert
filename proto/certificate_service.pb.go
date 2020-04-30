// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: certificate_service.proto

package gcert

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CertificateRequest_Endpoint int32

const (
	CertificateRequest_LE         CertificateRequest_Endpoint = 0
	CertificateRequest_LE_STAGING CertificateRequest_Endpoint = 1
)

// Enum value maps for CertificateRequest_Endpoint.
var (
	CertificateRequest_Endpoint_name = map[int32]string{
		0: "LE",
		1: "LE_STAGING",
	}
	CertificateRequest_Endpoint_value = map[string]int32{
		"LE":         0,
		"LE_STAGING": 1,
	}
)

func (x CertificateRequest_Endpoint) Enum() *CertificateRequest_Endpoint {
	p := new(CertificateRequest_Endpoint)
	*p = x
	return p
}

func (x CertificateRequest_Endpoint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificateRequest_Endpoint) Descriptor() protoreflect.EnumDescriptor {
	return file_certificate_service_proto_enumTypes[0].Descriptor()
}

func (CertificateRequest_Endpoint) Type() protoreflect.EnumType {
	return &file_certificate_service_proto_enumTypes[0]
}

func (x CertificateRequest_Endpoint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificateRequest_Endpoint.Descriptor instead.
func (CertificateRequest_Endpoint) EnumDescriptor() ([]byte, []int) {
	return file_certificate_service_proto_rawDescGZIP(), []int{0, 0}
}

type CertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domains  []string                    `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	Endpoint CertificateRequest_Endpoint `protobuf:"varint,2,opt,name=endpoint,proto3,enum=CertificateRequest_Endpoint" json:"endpoint,omitempty"`
}

func (x *CertificateRequest) Reset() {
	*x = CertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRequest) ProtoMessage() {}

func (x *CertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRequest.ProtoReflect.Descriptor instead.
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return file_certificate_service_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *CertificateRequest) GetEndpoint() CertificateRequest_Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return CertificateRequest_LE
}

type CertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultPaths []string `protobuf:"bytes,1,rep,name=vaultPaths,proto3" json:"vaultPaths,omitempty"`
	Success    bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CertificateResponse) Reset() {
	*x = CertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateResponse) ProtoMessage() {}

func (x *CertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateResponse.ProtoReflect.Descriptor instead.
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return file_certificate_service_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateResponse) GetVaultPaths() []string {
	if x != nil {
		return x.VaultPaths
	}
	return nil
}

func (x *CertificateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_certificate_service_proto protoreflect.FileDescriptor

var file_certificate_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x12,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x22, 0x4f, 0x0a, 0x13, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x68, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x53, 0x0a, 0x12, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x6d, 0x67, 0x69, 0x6c, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x63, 0x65, 0x72, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certificate_service_proto_rawDescOnce sync.Once
	file_certificate_service_proto_rawDescData = file_certificate_service_proto_rawDesc
)

func file_certificate_service_proto_rawDescGZIP() []byte {
	file_certificate_service_proto_rawDescOnce.Do(func() {
		file_certificate_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_service_proto_rawDescData)
	})
	return file_certificate_service_proto_rawDescData
}

var file_certificate_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_certificate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_certificate_service_proto_goTypes = []interface{}{
	(CertificateRequest_Endpoint)(0), // 0: CertificateRequest.Endpoint
	(*CertificateRequest)(nil),       // 1: CertificateRequest
	(*CertificateResponse)(nil),      // 2: CertificateResponse
}
var file_certificate_service_proto_depIdxs = []int32{
	0, // 0: CertificateRequest.endpoint:type_name -> CertificateRequest.Endpoint
	1, // 1: CertificateService.GetCertificate:input_type -> CertificateRequest
	2, // 2: CertificateService.GetCertificate:output_type -> CertificateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_certificate_service_proto_init() }
func file_certificate_service_proto_init() {
	if File_certificate_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certificate_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certificate_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_certificate_service_proto_goTypes,
		DependencyIndexes: file_certificate_service_proto_depIdxs,
		EnumInfos:         file_certificate_service_proto_enumTypes,
		MessageInfos:      file_certificate_service_proto_msgTypes,
	}.Build()
	File_certificate_service_proto = out.File
	file_certificate_service_proto_rawDesc = nil
	file_certificate_service_proto_goTypes = nil
	file_certificate_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateServiceClient interface {
	GetCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) GetCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/CertificateService/GetCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
type CertificateServiceServer interface {
	GetCertificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (*UnimplementedCertificateServiceServer) GetCertificate(context.Context, *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificate not implemented")
}

func RegisterCertificateServiceServer(s *grpc.Server, srv CertificateServiceServer) {
	s.RegisterService(&_CertificateService_serviceDesc, srv)
}

func _CertificateService_GetCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CertificateService/GetCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCertificate",
			Handler:    _CertificateService_GetCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate_service.proto",
}
