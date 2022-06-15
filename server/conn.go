package server

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/xiaolaji422/golink/lib/log"
)

// 长连接
type Conn struct {
	Mutex  *sync.Mutex     // Conn操作锁   回写，注册
	conn   *websocket.Conn // ws链接
	appid  string          // 链表分离
	userid string          // 链接的唯一id
}

const error_log = "send_msg_error.log"

type IConn interface {
	// 读取内容
	Read()
	// 回写数据
	Write(msg []byte) error
	// 关闭数据
	Close() error
	// 注册链接
	Register(appid, userid string) *Conn
}

func NewIConn(conn *websocket.Conn) IConn {
	return &Conn{
		conn:  conn,
		Mutex: &sync.Mutex{},
	}
}

func (c *Conn) Read() {
	if r := recover(); r != nil {
		fmt.Println("panic on Read:", r)
		goto Read
	}
	// var reply []byte
	// 建立连接后 接收来自客户端的信息reply
Read:
	typ, reply, err := c.conn.ReadMessage()

	if err != nil {
		fmt.Println("Read on error:", err.Error(), reply)
		return
	}
	if typ == 2 {
		// 文件传输
		goto Read
	}
	req, err := NewReq(c, reply)

	// ping
	var isping = strings.EqualFold(strings.Trim(string(reply), " "), "PING")
	if isping {
		go req.Response().SendMsg(PING, "success")
		goto Read
	}

	log.Instance().Info("read a requestion:", string(reply), req.AppId, req.UserId)
	if err != nil {
		req.Response().SendMsg(ERROR_PARAMS, err.Error(), "参数错误", string(reply))
		goto Read
	}

	if typ == 1 {
		if len(c.appid) == 0 {
			appid, userid, err := DefaultAuth(req)
			if err != nil {
				req.Response().SendMsg(NOT_LOGIN, err.Error())
				goto Read
			}
			req.Conn.Register(appid, userid)
			req.Response().SendMsg(SUCCESS, "注册成功", "success")
			goto Read
		} else {
			go DefaultFunc(req) // 并行
			goto Read
		}
	} else {
		req.Response().SendMsg(0, "不支持此类型")
		goto Read
	}
}

// 发送消息
func (c *Conn) Write(msg []byte) error {
	if c.conn == nil {
		return errors.New("none of conn")
	}
	// 内存锁住
	err := c.conn.WriteMessage(1, msg)
	if err != nil {
		log.Instance().Error("Error to write  user:" + err.Error())
	}
	return err
}

func (c *Conn) Close() error {
	err := c.conn.Close()
	if err != nil {
		log.Instance().Error("Error at close conn:" + err.Error())
	}
	return err
}

// 获取AppID
func (c *Conn) Register(appid, userid string) *Conn {
	log.Instance().Success("Register a User:", appid, userid)
	Conns.AddConn(appid, userid, c)
	c.appid = appid
	c.userid = userid
	return c
}
