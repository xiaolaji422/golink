package server

import (
	"errors"

	"git.woa.com/chengdukf/go-link/lib/helper"
)

// 注册的方法
type AuthFunc func(req *Request) (appid string, userid string, err error)

// 默认注册方法
var DefaultAuth AuthFunc = defaultAuth

// 默认注册方法实体
func defaultAuth(req *Request) (appid string, userid string, err error) {
	appid = helper.GetMapVal(req.Body, "appid", "").String()
	userid = helper.GetMapVal(req.Body, "userid", "").String()
	if len(appid) == 0 {
		err = errors.New("params is none  on key:appid")
		return
	}
	if len(userid) == 0 {
		err = errors.New("params is none  on key:userid")
		return
	}
	return appid, userid, nil
}
