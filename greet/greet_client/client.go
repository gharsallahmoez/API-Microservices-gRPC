package main

import (
	"context"
	"fmt"
	"github.com/gharsallahmoez/API-Microservices-gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting unary client rpc")
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
	fmt.Println("starting server streaming rpc")
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting client streaming rpc")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "moez",
				LastName:  "gharsallah",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bob",
				LastName:  "Bob",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Alice",
				LastName:  "Alice",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while reading steam : %v", err)
	}
	// iterate over the slice and send each message individually
	for _, req := range requests {
		fmt.Printf("sending request %v", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet : %v", err)
	}
	fmt.Printf("LongSteam response : %v\n",res.Result)
}
