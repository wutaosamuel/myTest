package main

import (
	"context"
	"fmt"
	"net"

	hw_pb "../helloworld"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, HReq *hw_pb.HelloRequest) (*hw_pb.HelloResponse, error) {
	return &hw_pb.HelloResponse{Message: "Hello" + HReq.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	hw_pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		fmt.Println(err)
		return
	}
}
