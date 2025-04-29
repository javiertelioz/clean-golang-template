package server

import (
	"context"
	"time"

	pb "github.com/javiertelioz/clean_architecture/pkg/interfaces/grpc/gen/helloworld/v1"
)

type HelloServiceServer struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *HelloServiceServer) SayHello(ctx context.Context, req *pb.GreeterServiceSayHelloRequest) (*pb.GreeterServiceSayHelloResponse, error) {
	return &pb.GreeterServiceSayHelloResponse{
		Message:   "Hello, " + req.GetName() + "!",
		Code:      200,
		Timestamp: time.Now().UnixMilli(),
	}, nil
}
