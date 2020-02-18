package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc"
	"io"
	"iris-grpc-example/proto"
	"log"
	"time"
)

var streamClient proto.StreamServiceClient

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug") //debug
	app.Handle("GET", "/testOrderList", orderList)
	app.Handle("GET", "/testUploadImage", uploadImage)
	app.Handle("GET", "/testSumData", sumData)
	app.Run(iris.Addr("127.0.0.1:8080"))
}
func init() {
	connect, err := grpc.Dial("127.0.0.1:9528", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	streamClient = proto.NewStreamServiceClient(connect)
}

func orderList(ctx iris.Context) {
	stream, err := streamClient.OrderList(context.Background(), &proto.OrderSearchParams{})
	if err != nil {
		ctx.JSON(map[string]string{
			"err": err.Error(),
		})
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(map[string]string{
				"err": err.Error(),
			})

			return
		}
		ctx.JSON(res)
		log.Println(res)
		time.Sleep(time.Second)
	}
}

func uploadImage(ctx iris.Context) {

}

func sumData(ctx iris.Context) {

}
