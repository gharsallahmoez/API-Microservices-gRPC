package main

import (
	"context"
	"fmt"
	"github.com/gharsallahmoez/API-Microservices-gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("greet function was invoked with : %v",req)
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen : %v: ", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
