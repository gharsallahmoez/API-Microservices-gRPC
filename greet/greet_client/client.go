package main

import (
	"context"
	"fmt"
	"github.com/gharsallahmoez/API-Microservices-gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("hello I'm a client")

	// create connection to server
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v: ", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting unary rpc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "moez",
			LastName:  "gharsallah",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC : %v", err)
	}
	log.Printf("response from greet : %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "moez",
			LastName:  "gharsallah",
		},
	}
	resSteam, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greetManyTime RPC : %v", err)
	}
	for {
		msg, err := resSteam.Recv()
		if err == io.EOF {
			// we've reached the end of steam
			break
		}
		if err != nil {
			log.Fatalf("error while reading steam : %v", err)
		}
		log.Printf("Response from GreetManyTimes : %v", msg.GetResult())

	}

}
