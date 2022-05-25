package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 启动服务
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 鉴定端口
// 监听链接
func Run(port string) {
	// 开始启动
	// http.HandleFunc("/", Accept)
	http.HandleFunc("/", Accept)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("websocket start:" + port)
}

// 获取到了链接
func Accept(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 多协程并发操作
	go NewIConn(conn).Read()
}

//
