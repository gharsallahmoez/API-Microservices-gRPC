package main

import (
	"fmt"
	"github.com/gharsallahmoez/API-Microservices-gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main(){
	fmt.Println("hello I'm a client")

	// create connection to server
	cc ,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v: ", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client %v",c)
}
