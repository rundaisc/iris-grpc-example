package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc"
	"iris-grpc-example/proto"
	"log"
)

var client proto.ServiceSearchClient

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug") //debug
	app.Handle("GET", "/testSaveUser", saveUser)
	app.Handle("GET", "/testUserInfo", userInfo)
	app.Run(iris.Addr("127.0.0.1:8080"))
}

func saveUser(ctx iris.Context) {
	params := proto.UserParams{}
	params.Age = &proto.Age{Age: 31}
	params.Name = &proto.Name{Name: "test"}
	res, err := client.SaveUser(context.Background(), &params)
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	ctx.JSON(res)
}
func userInfo(ctx iris.Context) {
	res, err := client.UserInfo(context.Background(), &proto.Id{Id: 1})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	ctx.JSON(res)
}

func init() {
	connect, err := grpc.Dial("127.0.0.1:9527", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client = proto.NewServiceSearchClient(connect)

}
