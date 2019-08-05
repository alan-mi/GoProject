package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	pb "proto/hello"
)

const (
	Address = "192.168.0.83:50052"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	print(in.GetName())
	resp.Message = fmt.Sprintf("hello%s", in.Name)
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln("aaaaa", err)

	}
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.crt", "../../keys/server.key")
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServer(s, HelloService)
	s.Serve(lis)
}
