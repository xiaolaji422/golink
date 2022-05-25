package server

import "encoding/json"

// 请求结构体
type Request struct {
	Body   map[string]interface{} // 请求内容
	Conn   *Conn
	AppId  string //当前请求的系统
	UserId string // 当前请求的用户
}

func NewReq(conn *Conn, body []byte) (*Request, error) {
	var data = make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	return &Request{
		Conn:   conn,
		Body:   data,
		AppId:  conn.appid,
		UserId: conn.userid,
	}, err
}

func (r *Request) GetBody() map[string]interface{} {
	return r.Body
}

// 获取返回体
func (r *Request) Response() *Response {
	return &Response{
		Conn:   r.Conn,
		AppId:  r.AppId,
		UserId: r.UserId,
	}
}
