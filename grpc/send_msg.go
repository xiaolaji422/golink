package grpc

import (
	"context"

	"github.com/xiaolaji422/golink/pb"
	"github.com/xiaolaji422/golink/server"
)

type grpcSer struct{}

// 给指定某人发消息
func (*grpcSer) SendMsg(ctx context.Context, in *pb.SendMsgReq) (*pb.SendMsgRep, error) {
	//
	var res *pb.SendMsgRep = &pb.SendMsgRep{}
	wxRsp, err := server.NewResponse(in.Appid, in.Userid)
	if err != nil {
		return res, err
	}
	err = wxRsp.Conn.Write([]byte(in.Message))
	return res, err
}

// 发送给指定的很多人
func (*grpcSer) SendAll(ctx context.Context, in *pb.SendMsgAllReq) (*pb.SendMsgAllRep, error) {
	//
	var res *pb.SendMsgAllRep = &pb.SendMsgAllRep{}
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
