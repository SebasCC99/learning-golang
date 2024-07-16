package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sebascc99/grpc-jwt-psql-demo/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServer
}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// TODO: Implement actual registration logic
	return &pb.RegisterResponse{Success: true, Message: "User registered successfully"}, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// TODO: Implement actual login logic
	return &pb.LoginResponse{Success: true, Token: "dummy_token"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
