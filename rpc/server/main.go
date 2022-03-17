package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	pb "github.com/ydnju/go_tour/rpc/proto"
	"google.golang.org/grpc"
)

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello, " + r.Name}, nil
}

func (s *GreeterServer) SayList(r *pb.HelloRequest, server pb.Greeter_SayListServer) error {
	for n := 0; n < 6; n++ {
		_ = server.Send(&pb.HelloReply{Message: "hello.list" + r.Name})
	}

	return nil
}

func (s *GreeterServer) SayRecord(server pb.Greeter_SayRecordServer) error {
	words := []string{}
	for {
		resp, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		words = append(words, resp.Name)
	}

	message := strings.Join(words[:], " ")
	server.SendAndClose(&pb.HelloReply{Message: message})
	return nil
}

func (s *GreeterServer) SayRoute(server pb.Greeter_SayRouteServer) error {
	n := 0
	_ = server.Send(&pb.HelloReply{Message: "Say Route"})

	pieces := []string{}
	for {
		resp, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		pieces = append(pieces, resp.Name)
		n++
	}

	message := strings.Join(pieces, ",")
	fmt.Printf("The message to response is %v\n", message)
	server.Send(&pb.HelloReply{Message: "blabla"})

	return nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8080")
	server.Serve(lis)
}
