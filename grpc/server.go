package grpc

import (
	"fmt"
	"log"
	"net"

	"git.woa.com/chengdukf/go-link/pb"
	"google.golang.org/grpc"
)

func Run(port string) {

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServer(s, &grpcSer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("grpc start:" + port)
}
