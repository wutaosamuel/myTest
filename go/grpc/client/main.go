package main

import (
	"context"
	"fmt"

	hw_pb "../helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := hw_pb.NewGreeterClient(conn)
	res, err := c.SayHello(context.Background(), &hw_pb.HelloRequest{Name: "World!"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Greeting: ", res.Message)
}
