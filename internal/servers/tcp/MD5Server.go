package tcp

import (
	"context"
	"log"
	pb "md5er/api/protomd5"
	"md5er/internal/md5"
)

type MD5Server struct {
	pb.UnimplementedMd5ErServer
	md5service *md5.MD5service
}

func (s MD5Server) Md5(ctx context.Context, req *pb.Md5Request) (*pb.Md5Response, error) {
	log.Println("got input", string(req.Data))
	result := s.md5service.AsByteArray(req.Data)
	log.Println("produced output", string(result[:]))
	return &pb.Md5Response{Hash: result[:]}, nil
}
