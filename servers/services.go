package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"iris-grpc-example/proto"
	"log"
	"math/rand"
	"net"
)

type ServiceSearch struct{}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9527")
	if err != nil {
		log.Fatalf("tcp listen failed:%v", err)
	}
	server := grpc.NewServer()
	fmt.Println("services start success")
	proto.RegisterServiceSearchServer(server, &ServiceSearch{})
	server.Serve(listen)

}

//保存用户
func (Service *ServiceSearch) SaveUser(ctx context.Context, params *proto.UserParams) (*proto.Id, error) {
	id := rand.Int31n(10) //随机生成id 模式保存成功
	res := &proto.Id{Id: id}
	fmt.Printf("username:%s,age:%d\r\n", params.Name, params.Age)
	return res, nil
}

func (Service *ServiceSearch) UserInfo(ctx context.Context, id *proto.Id) (*proto.User, error) {
	res := &proto.User{Id: id, Name: &proto.Name{Name: "test"}, Age: &proto.Age{Age: 31}}
	return res, nil
}
