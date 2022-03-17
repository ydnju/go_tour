package main

import (
	"context"
	"log"

	pb "github.com/ydnju/go_tour/rpc/proto"
	"google.golang.org/grpc"
)

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "yucdong"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}
