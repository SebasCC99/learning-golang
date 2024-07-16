package main

import (
	"context"
	"log"
	"net"

	"github.com/sebascc99/grpc-demo/proto"
	"google.golang.org/grpc"
)

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

func (s greeterServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello world",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("There was an error startings the server, %s", err)
	}

	serviceRegistrar := grpc.NewServer()

	proto.RegisterGreeterServer(serviceRegistrar, &greeterServer{})
	err = serviceRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("There was an error registering the service, %s", err)
	}

}
