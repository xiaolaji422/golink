项目目标
 - 轻便的websocket服务中心，可以支持大量的websocket链接和转发。
 - 耗费资源小，可以动态收缩和扩容所需资源，epoll模型。
 - 多系统接入，每个系统有属于自己的用户管理，权限校验。
 - 简易的路由管理，可以像http接口一样管理websocket接口。
 - 可以实时的向不同系统的不同用户推送消息。


 ##  proto 安装
 - https://github.com/protocolbuffers/protobuf/releases 下载解压
 - go install github.com/golang/protobuf/protoc-gen-go    
 - 检查gopath目录（%GOPATH%\）的bin文件夹，是否有protoc-gen-go.exe，protoc-gen-go-grpc.exe,拷贝至goroot（%GOROOT%\）的bin目录下
 - cd pb && protoc --go_out=plugins=grpc:. send_msg.proto  

