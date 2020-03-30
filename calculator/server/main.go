package main

import (
	"context"
	calculator "github.com/gharsallahmoez/API-Microservices-gRPC/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculator.SumRequest) ( *calculator.SumResponse,  error) {
	a := req.GetA()
	b := req.GetB()
	result := a + b
	res := &calculator.SumResponse{
		Result: result,
	}
	return res,nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	calculator.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
