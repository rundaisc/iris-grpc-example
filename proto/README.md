### proto 定义 rpc的数据类型，以及暴露 rpc接口


修改只需要修改 `servers.proto`文件，修改完成之后执行
`protoc --go_out=plugins=grpc:. *.proto`

更多语法请参考 `https://developers.google.com/protocol-buffers/docs/proto3`