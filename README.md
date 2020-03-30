# API-Microservices-gRPC

## prerequisite
##### install protocol buffer compiler from [here](https://developers.google.com/protocol-buffers/) or use [snap](https://snapcraft.io/protobuf) for linux
##### check if protoc is successfully installed 
    protoc --version
#### gRPC package for golang 
    go get -u google.golang.org/grpc
#### protobuf packages for golang 
    go get -u github.com/golang/protobuf/protoc-gen-go

## Generate proto files : 
:warning: without --proto_path because input and outpul files are in the same folder

    protoc greet/greetpb/greet.proto --go_out=plugins=grpc:. 