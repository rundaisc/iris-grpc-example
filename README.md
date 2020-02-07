# go-grpc的基本使用。使用iris作为消费者

### 安装
- go version >= 1.13
- set GO111MODULE=on

### 设置go中国代理
`set GOPROXY=https://goproxy.cn,direct`

### 下载依赖包
`go mod download`

### 运行
- cd servers
- go run services.go
- go run main.go

### 浏览器访问
- http://127.0.0.1:8080/testSaveInfo
- http://127.0.0.1:8080/testUserInfo


如果想要自己修改 services内容，比如新增一个函数，那么需要安装`protoc`
- `https://github.com/protocolbuffers/protobuf/releases` 安装`protoc`
- `go get -u github.com/golang/protobuf/protoc-gen-go` 安装`go-protoc`

安装之后 即可修改`proto/servers.proto`


