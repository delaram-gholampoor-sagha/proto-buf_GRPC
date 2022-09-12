package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Delaram-Gholampoor-Sagha/proto-buf_GRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked witj %v", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Print("Hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
