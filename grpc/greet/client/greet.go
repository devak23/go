package main

import (
	"context"
	pb "github.com/devak23/go/grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Abhay",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greetings: %s\n", res.Result)
}
