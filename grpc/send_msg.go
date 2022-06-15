package grpc

import (
	"context"
	"errors"

	"github.com/xiaolaji422/golink/lib/log"
	"github.com/xiaolaji422/golink/pb"
	"github.com/xiaolaji422/golink/server"
)

type grpcSer struct{}

// 给指定某人发消息
func (*grpcSer) SendMsg(ctx context.Context, in *pb.SendMsgReq) (*pb.SendMsgRep, error) {
	//
	var res *pb.SendMsgRep = &pb.SendMsgRep{}
	log.Instance().Info(in.Appid, in.Userid, in.Message)
	wxRsp, err := server.NewResponse(in.Appid, in.Userid)
	if err != nil {
		return res, err
	}
	err = wxRsp.Conn.Write([]byte(in.Message))
	return res, err
}

// 发送给所有用户
func (*grpcSer) SendAll(ctx context.Context, in *pb.SendMsgAllReq) (*pb.SendMsgAllRep, error) {
	var res *pb.SendMsgAllRep = &pb.SendMsgAllRep{}
	row, err := server.SendAll(in.Appid, []byte(in.Message))
	res.Count = int32(row)
	return res, err
}

//发送给指定的很多人
func (*grpcSer) SendToUsers(ctx context.Context, in *pb.SendToUsersReq) (*pb.SendToUsersRsp, error) {
	//
	var res *pb.SendToUsersRsp = &pb.SendToUsersRsp{}
	if len(in.Userid) == 0 {
		return nil, errors.New("请指定发送对象")
	}
	for _, v := range in.Userid {
		wxRsp, err := server.NewResponse(in.Appid, v)
		if err != nil {
			continue
		}
		err = wxRsp.Conn.Write([]byte(in.Message))
		if err != nil {
			res.Count++
		}
	}
	return res, nil
}
