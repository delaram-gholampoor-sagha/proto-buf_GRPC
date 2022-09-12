package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Delaram-Gholampoor-Sagha/proto-buf_GRPC/greet/greetpb"
	"google.golang.org/grpc"
)

// protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     greet/greetpb/greet.proto

// https://github.com/techschool/pcbook-go/issues/3

func main() {
	fmt.Println(" Hello I'm a client")

	// cc = client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect %v", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("Created CLient %f", c)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Printf("starting to doa unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "delaram",
			LastName:  "majestic",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Print("Response from Greet: %v ", res.Result)
}
