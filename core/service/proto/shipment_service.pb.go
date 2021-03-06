// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipment_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type LogisticFlowTrackRequest struct {
	ShipperCode          string   `protobuf:"bytes,1,opt,name=shipperCode,proto3" json:"shipperCode,omitempty"`
	LogisticCode         string   `protobuf:"bytes,2,opt,name=logisticCode,proto3" json:"logisticCode,omitempty"`
	Invert               bool     `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogisticFlowTrackRequest) Reset()         { *m = LogisticFlowTrackRequest{} }
func (m *LogisticFlowTrackRequest) String() string { return proto.CompactTextString(m) }
func (*LogisticFlowTrackRequest) ProtoMessage()    {}
func (*LogisticFlowTrackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_service_97bec78194ff804c, []int{0}
}
func (m *LogisticFlowTrackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogisticFlowTrackRequest.Unmarshal(m, b)
}
func (m *LogisticFlowTrackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogisticFlowTrackRequest.Marshal(b, m, deterministic)
}
func (dst *LogisticFlowTrackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogisticFlowTrackRequest.Merge(dst, src)
}
func (m *LogisticFlowTrackRequest) XXX_Size() int {
	return xxx_messageInfo_LogisticFlowTrackRequest.Size(m)
}
func (m *LogisticFlowTrackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogisticFlowTrackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogisticFlowTrackRequest proto.InternalMessageInfo

func (m *LogisticFlowTrackRequest) GetShipperCode() string {
	if m != nil {
		return m.ShipperCode
	}
	return ""
}

func (m *LogisticFlowTrackRequest) GetLogisticCode() string {
	if m != nil {
		return m.LogisticCode
	}
	return ""
}

func (m *LogisticFlowTrackRequest) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

type OrderLogisticTrackRequest struct {
	ShipOrderId          int64    `protobuf:"zigzag64,1,opt,name=shipOrderId,proto3" json:"shipOrderId,omitempty"`
	Invert               bool     `protobuf:"varint,2,opt,name=invert,proto3" json:"invert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderLogisticTrackRequest) Reset()         { *m = OrderLogisticTrackRequest{} }
func (m *OrderLogisticTrackRequest) String() string { return proto.CompactTextString(m) }
func (*OrderLogisticTrackRequest) ProtoMessage()    {}
func (*OrderLogisticTrackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_service_97bec78194ff804c, []int{1}
}
func (m *OrderLogisticTrackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLogisticTrackRequest.Unmarshal(m, b)
}
func (m *OrderLogisticTrackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLogisticTrackRequest.Marshal(b, m, deterministic)
}
func (dst *OrderLogisticTrackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLogisticTrackRequest.Merge(dst, src)
}
func (m *OrderLogisticTrackRequest) XXX_Size() int {
	return xxx_messageInfo_OrderLogisticTrackRequest.Size(m)
}
func (m *OrderLogisticTrackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLogisticTrackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLogisticTrackRequest proto.InternalMessageInfo

func (m *OrderLogisticTrackRequest) GetShipOrderId() int64 {
	if m != nil {
		return m.ShipOrderId
	}
	return 0
}

func (m *OrderLogisticTrackRequest) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

// 发货单追踪
type SShipOrderTrack struct {
	// 返回状态码
	Code int32 `protobuf:"zigzag32,1,opt,name=Code,proto3" json:"Code,omitempty"`
	// 返回错误信息
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	// 物流单号
	LogisticCode string `protobuf:"bytes,3,opt,name=LogisticCode,proto3" json:"LogisticCode,omitempty"`
	// 承运商代码
	ShipperCode string `protobuf:"bytes,4,opt,name=ShipperCode,proto3" json:"ShipperCode,omitempty"`
	// * 承运商名称
	ShipperName string `protobuf:"bytes,5,opt,name=ShipperName,proto3" json:"ShipperName,omitempty"`
	// 发货状态
	ShipState string `protobuf:"bytes,6,opt,name=ShipState,proto3" json:"ShipState,omitempty"`
	// 更新时间
	UpdateTime int64 `protobuf:"zigzag64,7,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	// 包含发货单流
	Flows                []*SShipFlow `protobuf:"bytes,8,rep,name=Flows,proto3" json:"Flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SShipOrderTrack) Reset()         { *m = SShipOrderTrack{} }
func (m *SShipOrderTrack) String() string { return proto.CompactTextString(m) }
func (*SShipOrderTrack) ProtoMessage()    {}
func (*SShipOrderTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_service_97bec78194ff804c, []int{2}
}
func (m *SShipOrderTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShipOrderTrack.Unmarshal(m, b)
}
func (m *SShipOrderTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShipOrderTrack.Marshal(b, m, deterministic)
}
func (dst *SShipOrderTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShipOrderTrack.Merge(dst, src)
}
func (m *SShipOrderTrack) XXX_Size() int {
	return xxx_messageInfo_SShipOrderTrack.Size(m)
}
func (m *SShipOrderTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_SShipOrderTrack.DiscardUnknown(m)
}

var xxx_messageInfo_SShipOrderTrack proto.InternalMessageInfo

func (m *SShipOrderTrack) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SShipOrderTrack) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SShipOrderTrack) GetLogisticCode() string {
	if m != nil {
		return m.LogisticCode
	}
	return ""
}

func (m *SShipOrderTrack) GetShipperCode() string {
	if m != nil {
		return m.ShipperCode
	}
	return ""
}

func (m *SShipOrderTrack) GetShipperName() string {
	if m != nil {
		return m.ShipperName
	}
	return ""
}

func (m *SShipOrderTrack) GetShipState() string {
	if m != nil {
		return m.ShipState
	}
	return ""
}

func (m *SShipOrderTrack) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *SShipOrderTrack) GetFlows() []*SShipFlow {
	if m != nil {
		return m.Flows
	}
	return nil
}

// 发货流
type SShipFlow struct {
	// 记录标题
	Subject string `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	// 记录时间
	CreateTime string `protobuf:"bytes,2,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// 备注
	Remark               string   `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShipFlow) Reset()         { *m = SShipFlow{} }
func (m *SShipFlow) String() string { return proto.CompactTextString(m) }
func (*SShipFlow) ProtoMessage()    {}
func (*SShipFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_shipment_service_97bec78194ff804c, []int{3}
}
func (m *SShipFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShipFlow.Unmarshal(m, b)
}
func (m *SShipFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShipFlow.Marshal(b, m, deterministic)
}
func (dst *SShipFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShipFlow.Merge(dst, src)
}
func (m *SShipFlow) XXX_Size() int {
	return xxx_messageInfo_SShipFlow.Size(m)
}
func (m *SShipFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_SShipFlow.DiscardUnknown(m)
}

var xxx_messageInfo_SShipFlow proto.InternalMessageInfo

func (m *SShipFlow) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SShipFlow) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *SShipFlow) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*LogisticFlowTrackRequest)(nil), "LogisticFlowTrackRequest")
	proto.RegisterType((*OrderLogisticTrackRequest)(nil), "OrderLogisticTrackRequest")
	proto.RegisterType((*SShipOrderTrack)(nil), "SShipOrderTrack")
	proto.RegisterType((*SShipFlow)(nil), "SShipFlow")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShipmentServiceClient is the client API for ShipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipmentServiceClient interface {
	// * 物流追踪
	GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
}

type shipmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewShipmentServiceClient(cc *grpc.ClientConn) ShipmentServiceClient {
	return &shipmentServiceClient{cc}
}

func (c *shipmentServiceClient) GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, "/ShipmentService/GetLogisticFlowTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, "/ShipmentService/ShipOrderLogisticTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentServiceServer is the server API for ShipmentService service.
type ShipmentServiceServer interface {
	// * 物流追踪
	GetLogisticFlowTrack(context.Context, *LogisticFlowTrackRequest) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(context.Context, *OrderLogisticTrackRequest) (*SShipOrderTrack, error)
}

func RegisterShipmentServiceServer(s *grpc.Server, srv ShipmentServiceServer) {
	s.RegisterService(&_ShipmentService_serviceDesc, srv)
}

func _ShipmentService_GetLogisticFlowTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogisticFlowTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/GetLogisticFlowTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, req.(*LogisticFlowTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ShipOrderLogisticTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderLogisticTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/ShipOrderLogisticTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, req.(*OrderLogisticTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShipmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShipmentService",
	HandlerType: (*ShipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLogisticFlowTrack",
			Handler:    _ShipmentService_GetLogisticFlowTrack_Handler,
		},
		{
			MethodName: "ShipOrderLogisticTrack",
			Handler:    _ShipmentService_ShipOrderLogisticTrack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipment_service.proto",
}

func init() {
	proto.RegisterFile("shipment_service.proto", fileDescriptor_shipment_service_97bec78194ff804c)
}

var fileDescriptor_shipment_service_97bec78194ff804c = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0x2d, 0xdf, 0x1d, 0x4c, 0xd0, 0x8d, 0x21, 0x0b, 0x31, 0xa6, 0xe9, 0x89, 0x53, 0x0f,
	0x78, 0xf4, 0x26, 0xf1, 0x2b, 0x41, 0x4d, 0x5a, 0xb8, 0x98, 0x18, 0x53, 0xda, 0x09, 0x56, 0x28,
	0xad, 0xbb, 0x0b, 0xfa, 0x06, 0xbe, 0x87, 0x4f, 0x6a, 0x76, 0xdb, 0xe2, 0x56, 0xe1, 0xd4, 0xce,
	0x6f, 0x67, 0x67, 0x76, 0xfe, 0xf3, 0x87, 0x2e, 0x7f, 0x8d, 0xd2, 0x18, 0x57, 0xe2, 0x85, 0x23,
	0xdb, 0x44, 0x01, 0x3a, 0x29, 0x4b, 0x44, 0x62, 0x7f, 0x02, 0x1d, 0x27, 0xf3, 0x88, 0x8b, 0x28,
	0xb8, 0x5e, 0x26, 0x1f, 0x13, 0xe6, 0x07, 0x0b, 0x17, 0xdf, 0xd7, 0xc8, 0x05, 0xb1, 0xa0, 0x2d,
	0x6f, 0xa5, 0xc8, 0x46, 0x49, 0x88, 0xd4, 0xb0, 0x8c, 0x81, 0xe9, 0xea, 0x88, 0xd8, 0x70, 0xb8,
	0xcc, 0x6f, 0xab, 0x94, 0x8a, 0x4a, 0x29, 0x31, 0xd2, 0x85, 0x46, 0xb4, 0xda, 0x20, 0x13, 0xb4,
	0x6a, 0x19, 0x83, 0x96, 0x9b, 0x47, 0xf6, 0x14, 0x7a, 0x8f, 0x2c, 0x44, 0x56, 0xb4, 0xdf, 0xd5,
	0x5a, 0x25, 0xdc, 0x85, 0xaa, 0x35, 0x71, 0x75, 0xa4, 0x95, 0xad, 0x94, 0xca, 0x7e, 0x55, 0xa0,
	0xe3, 0x79, 0x45, 0xa2, 0x2a, 0x4a, 0x08, 0xd4, 0xb6, 0x13, 0x1c, 0xbb, 0xea, 0x9f, 0x50, 0x68,
	0xde, 0x23, 0xe7, 0xfe, 0xbc, 0x78, 0x75, 0x11, 0xca, 0xa1, 0xc6, 0xfa, 0x50, 0xd5, 0x6c, 0x28,
	0x9d, 0xc9, 0xf7, 0x79, 0x9a, 0x34, 0xb5, 0x4c, 0x1a, 0x0d, 0x69, 0x19, 0x0f, 0x7e, 0x8c, 0xb4,
	0x5e, 0xca, 0x90, 0x88, 0x9c, 0x82, 0x29, 0x43, 0x4f, 0xf8, 0x02, 0x69, 0x43, 0x9d, 0xff, 0x02,
	0x72, 0x06, 0x30, 0x4d, 0x43, 0x5f, 0xe0, 0x24, 0x8a, 0x91, 0x36, 0x95, 0x00, 0x1a, 0x21, 0x16,
	0xd4, 0xe5, 0xc2, 0x38, 0x6d, 0x59, 0xd5, 0x41, 0x7b, 0x08, 0x8e, 0x1a, 0x5a, 0x22, 0x37, 0x3b,
	0xb0, 0x9f, 0xc1, 0xdc, 0x32, 0x39, 0xae, 0xb7, 0x9e, 0xbd, 0x61, 0x20, 0xf2, 0x3d, 0x16, 0xa1,
	0x6c, 0x34, 0x62, 0x58, 0x34, 0xca, 0xb4, 0xd0, 0x88, 0x14, 0xda, 0xc5, 0xd8, 0x67, 0x8b, 0x5c,
	0x88, 0x3c, 0x1a, 0x7e, 0x1b, 0xd0, 0xf1, 0x72, 0x53, 0x79, 0x99, 0xa7, 0xc8, 0x15, 0x9c, 0xdc,
	0xa0, 0xf8, 0x67, 0x28, 0xd2, 0x73, 0xf6, 0x99, 0xac, 0x7f, 0xe4, 0xfc, 0xd9, 0x96, 0x7d, 0x40,
	0x6e, 0xa1, 0xbb, 0x65, 0x25, 0x7b, 0x90, 0xbe, 0xb3, 0xd7, 0x33, 0xbb, 0x2a, 0x5d, 0x9a, 0x4f,
	0x4d, 0xe7, 0x42, 0x39, 0x7d, 0xd6, 0x50, 0x9f, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08,
	0xdf, 0xfe, 0xe4, 0x0a, 0x03, 0x00, 0x00,
}
