package tcp

import (
	"google.golang.org/grpc"
	"log"
	pb "md5er/api/protomd5"
	"md5er/internal/md5"
	"net"
)

type TcpGrpcServer struct {
	grpcHandler *MD5Server
}

func New(service *md5.MD5service) *TcpGrpcServer {
	return &TcpGrpcServer{
		grpcHandler: &MD5Server{
			md5service: service,
		},
	}
}
func (s *TcpGrpcServer) Start() {
	lis, err := net.Listen("tcp", "127.0.0.1:9616")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterMd5ErServer(grpcSrv, s.grpcHandler)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
