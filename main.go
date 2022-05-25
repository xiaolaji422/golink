package main

import (
	"fmt"

	"git.woa.com/chengdukf/go-link/config"
	"git.woa.com/chengdukf/go-link/grpc"
	"git.woa.com/chengdukf/go-link/server"
	_ "google.golang.org/grpc"
	// _ ""
)

func main() {
	// 启动grpc
	grpc_port := config.Conf.Get("grpc.port")
	go grpc.Run(fmt.Sprintf("%d", grpc_port))
	// 启动ws
	ws_port := config.Conf.Get("websocket.port")
	server.Run(fmt.Sprintf("%s", ws_port))
}
