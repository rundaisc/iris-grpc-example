// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streaming.proto

package proto

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

type Order struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderSn              string   `protobuf:"bytes,2,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetOrderSn() string {
	if m != nil {
		return m.OrderSn
	}
	return ""
}

func (m *Order) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type StreamOrderList struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamOrderList) Reset()         { *m = StreamOrderList{} }
func (m *StreamOrderList) String() string { return proto.CompactTextString(m) }
func (*StreamOrderList) ProtoMessage()    {}
func (*StreamOrderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{1}
}

func (m *StreamOrderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamOrderList.Unmarshal(m, b)
}
func (m *StreamOrderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamOrderList.Marshal(b, m, deterministic)
}
func (m *StreamOrderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamOrderList.Merge(m, src)
}
func (m *StreamOrderList) XXX_Size() int {
	return xxx_messageInfo_StreamOrderList.Size(m)
}
func (m *StreamOrderList) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamOrderList.DiscardUnknown(m)
}

var xxx_messageInfo_StreamOrderList proto.InternalMessageInfo

func (m *StreamOrderList) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type OrderSearchParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderSearchParams) Reset()         { *m = OrderSearchParams{} }
func (m *OrderSearchParams) String() string { return proto.CompactTextString(m) }
func (*OrderSearchParams) ProtoMessage()    {}
func (*OrderSearchParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{2}
}

func (m *OrderSearchParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderSearchParams.Unmarshal(m, b)
}
func (m *OrderSearchParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderSearchParams.Marshal(b, m, deterministic)
}
func (m *OrderSearchParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderSearchParams.Merge(m, src)
}
func (m *OrderSearchParams) XXX_Size() int {
	return xxx_messageInfo_OrderSearchParams.Size(m)
}
func (m *OrderSearchParams) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderSearchParams.DiscardUnknown(m)
}

var xxx_messageInfo_OrderSearchParams proto.InternalMessageInfo

type Image struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	File                 string   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Image) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

type StreamImageList struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamImageList) Reset()         { *m = StreamImageList{} }
func (m *StreamImageList) String() string { return proto.CompactTextString(m) }
func (*StreamImageList) ProtoMessage()    {}
func (*StreamImageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{4}
}

func (m *StreamImageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamImageList.Unmarshal(m, b)
}
func (m *StreamImageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamImageList.Marshal(b, m, deterministic)
}
func (m *StreamImageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamImageList.Merge(m, src)
}
func (m *StreamImageList) XXX_Size() int {
	return xxx_messageInfo_StreamImageList.Size(m)
}
func (m *StreamImageList) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamImageList.DiscardUnknown(m)
}

var xxx_messageInfo_StreamImageList proto.InternalMessageInfo

func (m *StreamImageList) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type UploadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{5}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

type StreamSumData struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamSumData) Reset()         { *m = StreamSumData{} }
func (m *StreamSumData) String() string { return proto.CompactTextString(m) }
func (*StreamSumData) ProtoMessage()    {}
func (*StreamSumData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2e8ceba11142904, []int{6}
}

func (m *StreamSumData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamSumData.Unmarshal(m, b)
}
func (m *StreamSumData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamSumData.Marshal(b, m, deterministic)
}
func (m *StreamSumData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamSumData.Merge(m, src)
}
func (m *StreamSumData) XXX_Size() int {
	return xxx_messageInfo_StreamSumData.Size(m)
}
func (m *StreamSumData) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamSumData.DiscardUnknown(m)
}

var xxx_messageInfo_StreamSumData proto.InternalMessageInfo

func (m *StreamSumData) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "proto.Order")
	proto.RegisterType((*StreamOrderList)(nil), "proto.StreamOrderList")
	proto.RegisterType((*OrderSearchParams)(nil), "proto.OrderSearchParams")
	proto.RegisterType((*Image)(nil), "proto.Image")
	proto.RegisterType((*StreamImageList)(nil), "proto.StreamImageList")
	proto.RegisterType((*UploadResponse)(nil), "proto.uploadResponse")
	proto.RegisterType((*StreamSumData)(nil), "proto.StreamSumData")
}

func init() { proto.RegisterFile("streaming.proto", fileDescriptor_f2e8ceba11142904) }

var fileDescriptor_f2e8ceba11142904 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xbb, 0x7d, 0xdf, 0xb4, 0x76, 0xd4, 0x56, 0x47, 0x2d, 0x21, 0xa7, 0xb2, 0x17, 0x7b,
	0x2a, 0xa5, 0x22, 0x1e, 0x3c, 0x88, 0xa0, 0x82, 0x20, 0x2a, 0x1b, 0xfc, 0x00, 0xdb, 0x66, 0xac,
	0x0b, 0xdd, 0xa4, 0x6c, 0x52, 0x3f, 0xa7, 0x1f, 0x49, 0x32, 0xd9, 0xfe, 0x0b, 0x9e, 0x76, 0x9e,
	0x79, 0x66, 0x66, 0x7f, 0x33, 0xd0, 0xcb, 0x0b, 0x47, 0xda, 0x9a, 0x74, 0x3e, 0x5a, 0xba, 0xac,
	0xc8, 0x30, 0xe0, 0x47, 0x3e, 0x42, 0xf0, 0xe6, 0x12, 0x72, 0xd8, 0x85, 0xa6, 0x49, 0x42, 0x31,
	0x10, 0xc3, 0x40, 0x35, 0x4d, 0x82, 0x21, 0xb4, 0xb3, 0xd2, 0x88, 0xd3, 0xb0, 0x39, 0x10, 0xc3,
	0x8e, 0x5a, 0x4b, 0x44, 0xf8, 0x9f, 0xe8, 0x82, 0xc2, 0x7f, 0x9c, 0xe6, 0x58, 0x5e, 0x43, 0x2f,
	0xe6, 0x0f, 0x78, 0xd8, 0x8b, 0xc9, 0x0b, 0x94, 0x10, 0x70, 0x07, 0xcf, 0x3c, 0x9c, 0x1c, 0x55,
	0xff, 0x8e, 0xb8, 0x40, 0x55, 0x96, 0x3c, 0x83, 0x53, 0xd6, 0x31, 0x69, 0x37, 0xfb, 0x7a, 0xd7,
	0x4e, 0xdb, 0x5c, 0xde, 0x40, 0xf0, 0x6c, 0xf5, 0x9c, 0x30, 0x82, 0x83, 0x4f, 0xb3, 0xa0, 0x57,
	0x6d, 0x89, 0x87, 0x74, 0xd4, 0x46, 0x97, 0x10, 0x65, 0xec, 0xd9, 0x38, 0xde, 0x42, 0x70, 0xfb,
	0x1a, 0xc2, 0x94, 0xa2, 0x06, 0xc1, 0x05, 0xaa, 0xb2, 0xe4, 0x09, 0x74, 0x57, 0xcb, 0x45, 0xa6,
	0x13, 0x45, 0xf9, 0x32, 0x4b, 0x73, 0x92, 0x97, 0x70, 0x5c, 0x0d, 0x8a, 0x57, 0xf6, 0x41, 0x17,
	0x1a, 0xfb, 0xd0, 0x4a, 0x57, 0x76, 0xea, 0x97, 0x09, 0x94, 0x57, 0x93, 0x1f, 0xb1, 0xa9, 0x24,
	0xf7, 0x6d, 0x66, 0x84, 0xf7, 0xd0, 0xd9, 0x9e, 0x20, 0xdc, 0xdd, 0x79, 0x77, 0xc7, 0xa8, 0xef,
	0x9d, 0xda, 0xd1, 0x64, 0x63, 0x2c, 0xf0, 0x0e, 0xe0, 0x83, 0x79, 0x9e, 0xcc, 0x82, 0x70, 0xbf,
	0x72, 0xb3, 0x59, 0x74, 0xe1, 0xf3, 0x35, 0xf4, 0xc6, 0x50, 0xe0, 0x2d, 0xb4, 0xd7, 0xe0, 0xe7,
	0x7b, 0xdd, 0x3e, 0x1b, 0xfd, 0x99, 0x2d, 0x5b, 0xc7, 0x62, 0xda, 0x62, 0xeb, 0xea, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x40, 0xe6, 0x7f, 0xf0, 0x31, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	OrderList(ctx context.Context, in *OrderSearchParams, opts ...grpc.CallOption) (StreamService_OrderListClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (StreamService_UploadFileClient, error)
	SumData(ctx context.Context, opts ...grpc.CallOption) (StreamService_SumDataClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) OrderList(ctx context.Context, in *OrderSearchParams, opts ...grpc.CallOption) (StreamService_OrderListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/OrderList", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceOrderListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_OrderListClient interface {
	Recv() (*StreamOrderList, error)
	grpc.ClientStream
}

type streamServiceOrderListClient struct {
	grpc.ClientStream
}

func (x *streamServiceOrderListClient) Recv() (*StreamOrderList, error) {
	m := new(StreamOrderList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (StreamService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/proto.StreamService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceUploadFileClient{stream}
	return x, nil
}

type StreamService_UploadFileClient interface {
	Send(*StreamImageList) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type streamServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *streamServiceUploadFileClient) Send(m *StreamImageList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceUploadFileClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) SumData(ctx context.Context, opts ...grpc.CallOption) (StreamService_SumDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/proto.StreamService/SumData", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSumDataClient{stream}
	return x, nil
}

type StreamService_SumDataClient interface {
	Send(*StreamSumData) error
	Recv() (*StreamSumData, error)
	grpc.ClientStream
}

type streamServiceSumDataClient struct {
	grpc.ClientStream
}

func (x *streamServiceSumDataClient) Send(m *StreamSumData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSumDataClient) Recv() (*StreamSumData, error) {
	m := new(StreamSumData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	OrderList(*OrderSearchParams, StreamService_OrderListServer) error
	UploadFile(StreamService_UploadFileServer) error
	SumData(StreamService_SumDataServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) OrderList(req *OrderSearchParams, srv StreamService_OrderListServer) error {
	return status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (*UnimplementedStreamServiceServer) UploadFile(srv StreamService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedStreamServiceServer) SumData(srv StreamService_SumDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SumData not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_OrderList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderSearchParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).OrderList(m, &streamServiceOrderListServer{stream})
}

type StreamService_OrderListServer interface {
	Send(*StreamOrderList) error
	grpc.ServerStream
}

type streamServiceOrderListServer struct {
	grpc.ServerStream
}

func (x *streamServiceOrderListServer) Send(m *StreamOrderList) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).UploadFile(&streamServiceUploadFileServer{stream})
}

type StreamService_UploadFileServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*StreamImageList, error)
	grpc.ServerStream
}

type streamServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *streamServiceUploadFileServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceUploadFileServer) Recv() (*StreamImageList, error) {
	m := new(StreamImageList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_SumData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).SumData(&streamServiceSumDataServer{stream})
}

type StreamService_SumDataServer interface {
	Send(*StreamSumData) error
	Recv() (*StreamSumData, error)
	grpc.ServerStream
}

type streamServiceSumDataServer struct {
	grpc.ServerStream
}

func (x *streamServiceSumDataServer) Send(m *StreamSumData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSumDataServer) Recv() (*StreamSumData, error) {
	m := new(StreamSumData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OrderList",
			Handler:       _StreamService_OrderList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _StreamService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SumData",
			Handler:       _StreamService_SumData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "streaming.proto",
}