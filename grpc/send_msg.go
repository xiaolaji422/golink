package grpc

import (
	"context"

	"git.woa.com/chengdukf/go-link/pb"
	"git.woa.com/chengdukf/go-link/server"
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
