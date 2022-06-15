package main

import (
	"fmt"

	"github.com/xiaolaji422/golink/config"
	"github.com/xiaolaji422/golink/grpc"
	"github.com/xiaolaji422/golink/server"
	_ "google.golang.org/grpc"
	// _ ""
)

func main() {
	// 启动grpc
	grpc_port := config.Conf.Get("grpc.port")
	go grpc.Run(fmt.Sprintf("%d", grpc_port))
	// // 启动ws
	ws_port := config.Conf.Get("websocket.port")
	server.Run(fmt.Sprintf("%s", ws_port))

}
