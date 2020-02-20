package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc"
	"io"
	"iris-grpc-example/proto"
	"log"
	"strconv"
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
	}
}

func uploadImage(ctx iris.Context) {
	stream, err := streamClient.UploadFile(context.Background())
	if err != nil {
		ctx.JSON(map[string]string{
			"err": err.Error(),
		})
		return
	}
	for i := 1; i <= 10; i++ {
		img := &proto.Image{FileName: "image" + strconv.Itoa(i), File: "file data"}
		images := &proto.StreamImageList{Image: img}
		err := stream.Send(images)
		if err != nil {
			ctx.JSON(map[string]string{
				"err": err.Error(),
			})
			return
		}
	}
	//发送完毕 关闭并获取服务端返回的消息
	resp, err := stream.CloseAndRecv()
	if err != nil {
		ctx.JSON(map[string]string{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(map[string]interface{}{"result": resp, "message": "success"})
	log.Println(resp)
}

func sumData(ctx iris.Context) {
	stream, err := streamClient.SumData(context.Background())
	if err != nil {
		ctx.JSON(map[string]string{
			"err": err.Error(),
		})
		return
	}
	for i := 1; i <= 10; i++ {
		err = stream.Send(&proto.StreamSumData{Number: int32(i)})
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		log.Printf("res number:%d", res.Number)
	}
	stream.CloseSend()
	return
}
