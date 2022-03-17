package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ydnju/go_tour/rpc/proto"
	"google.golang.org/grpc"
)

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "yucdong"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Resp: %v", resp)
	}

	return nil
}

func SayRecord(client pb.GreeterClient) error {
	stream, _ := client.SayRecord(context.Background())
	stream.Send(&pb.HelloRequest{Name: "What"})
	stream.Send(&pb.HelloRequest{Name: "a"})
	stream.Send(&pb.HelloRequest{Name: "Movie"})
	reply, err := stream.CloseAndRecv()
	if err == nil {
		fmt.Printf("Resp: %v", reply)
	}

	return err
}

func SayRoute(client pb.GreeterClient) error {
	stream, _ := client.SayRoute(context.Background())
	hint, err := stream.Recv()
	fmt.Printf("Received %v, start to transmit messages.\n", hint)

	stream.Send(&pb.HelloRequest{Name: "What"})
	stream.Send(&pb.HelloRequest{Name: "a"})
	stream.Send(&pb.HelloRequest{Name: "Movie"})

	// In our code, this must be called before
	// stream.Recv() because the server will read
	// EOF first then break out of the loop and return
	// a Joined message.
	stream.CloseSend()

	reply, err := stream.Recv()
	if err == nil {
		fmt.Printf("Resp is: %v\n", reply)
	}

	return err
}

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayRoute(client)
}
