package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sebascc99/grpc-jwt-psql-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test Register
	r, err := c.Register(ctx, &pb.RegisterRequest{Username: "testuser", Password: "testpass"})
	if err != nil {
		log.Fatalf("could not register: %v", err)
	}
	log.Printf("Register Response: %s", r.GetMessage())

	// Test Login
	l, err := c.Login(ctx, &pb.LoginRequest{Username: "testuser", Password: "testpass"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("Login Response: %s", l.GetToken())
}
