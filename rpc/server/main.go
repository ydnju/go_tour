package main

import (
	"context"
	"net"

	pb "github.com/ydnju/go_tour/rpc/proto"
	"google.golang.org/grpc"
)

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello, " + r.Name}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8080")
	server.Serve(lis)
}
