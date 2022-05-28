package server

import (
	"fmt"

	"github.com/xiaolaji422/golink/lib/helper"
	lib "github.com/xiaolaji422/golink/lib/helper"
)

type Handler func(req *Request)

// 默认的业务回调
var DefaultFunc Handler = defaultFunc

func defaultFunc(req *Request) {
	var (
		params = req.Body
		cmd    = lib.GetMapVal(params, "cmd", "sendTo").String()
	)
	switch cmd {
	case "sendTo":
		// 给某人发消息
		var (
			AppID  = helper.GetMapVal(params, "appid", "").String()
			UserID = helper.GetMapVal(params, "userid", "").String()
			msg    = helper.GetMapVal(params, "msg", "有人让我给你发消息").String()
		)
		err := SendToUser(AppID, UserID, []byte(msg))
		if err != nil {
			req.Response().SendMsg(ERROR, err.Error())
			return
		}
		err = req.Response().SendMsg(SUCCESS, "发送成功")
		if err != nil {
			fmt.Println("error on sendMsg:", req.AppId+"_"+req.UserId)
			return
		}
		return
	default:
		// 路由到某个方法
		// 不知道干啥
	}
	req.Response().SendMsg(0, "没有鉴权", req.Body)
}

func HandleFunc(f Handler) {
	DefaultFunc = f
}
