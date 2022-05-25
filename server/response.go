package server

import (
	"encoding/json"
	"errors"
	"fmt"

	"git.woa.com/chengdukf/go-link/lib/helper"
)

// 返回结构体
type Response struct {
	Conn   *Conn
	AppId  string //当前请求的系统
	UserId string // 当前请求的用户
}

func NewResponse(AppID, UserID string) (*Response, error) {
	var rep = &Response{
		AppId:  AppID,
		UserId: UserID,
	}
	var conn = Conns.GetConn(AppID, UserID)
	if helper.IsNil(conn) {
		fmt.Println(Conns, AppID, UserID, "当前用户暂不在线")
		return nil, errors.New("当前用户暂不在线")
	}
	rep.Conn = conn
	return rep, nil
}

//	发送消息
func (r *Response) SendMsg(code int, msg string, list ...interface{}) error {
	var res = map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	if len(list) == 1 {
		res["data"] = list[0]
		list = list[1:]
	}
	if len(list) > 0 {
		res["list"] = list
	}
	res_byt, err := json.Marshal(res)
	if err != nil {
		return err
	}
	// 获取Conn
	r.Conn.Write(res_byt)
	return err
}

// 发送消息到指定某人
func SendToUser(AppID, UserID string, msg []byte) error {
	rep, err := NewResponse(AppID, UserID)
	if err != nil {
		return err
	}
	return rep.Conn.Write(msg)
}
