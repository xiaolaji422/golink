package main

import (
	"fmt"

	"github.com/xiaolaji422/golink/config"
	"github.com/xiaolaji422/golink/lib/log"
	_ "google.golang.org/grpc"
	// _ ""
)

func main() {
	// 启动grpc
	grpc_port := config.Conf.Get("grpc.port")
	// go grpc.Run(fmt.Sprintf("%d", grpc_port))
	// // 启动ws
	// ws_port := config.Conf.Get("websocket.port")
	// server.Run(fmt.Sprintf("%s", ws_port))
	err := log.Instance().Info(fmt.Sprintf("%s", grpc_port))
	if err != nil {
		fmt.Println(err.Error(), "main")
	}
}
