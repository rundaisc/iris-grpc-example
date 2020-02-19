package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"iris-grpc-example/proto"
	"log"
	"net"
	"time"
)

type StreamServices struct{}

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamServices{})

	lis, err := net.Listen("tcp", "127.0.0.1:9528")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
func (services *StreamServices) OrderList(params *proto.OrderSearchParams, stream proto.StreamService_OrderListServer) error {
	for i := 0; i <= 10; i++ {
		order := proto.Order{
			Id:      int32(i),
			OrderSn: time.Now().Format("20060102150405") + "order_sn",
			Date:    time.Now().Format("2006-01-02 15:04:05"),
		}
		err := stream.Send(&proto.StreamOrderList{
			Order: &order,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (services *StreamServices) UploadFile(stream proto.StreamService_UploadFileServer) error {
	for  {
		 res,err := stream.Recv()
		 //接收消息结束，发送结果，并关闭
		if err == io.EOF {
			return stream.SendAndClose(&proto.UploadResponse{})
		}
		if err !=nil {
			return err
		}
		fmt.Println(res)
	}
	return nil
}

func (services *StreamServices) SumData(stream proto.StreamService_SumDataServer) error {
	return nil
}
