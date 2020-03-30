package main

import (
	"context"
	"fmt"
	calculator "github.com/gharsallahmoez/API-Microservices-gRPC/calculator/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {

	s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	c := calculator.NewSumServiceClient(s)
	req := &calculator.SumRequest{
		A: 5,
		B: 2,
	}
	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("here the result", resp.Result)
}
