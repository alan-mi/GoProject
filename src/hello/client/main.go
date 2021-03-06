package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
	pb "proto/hello" // 引入proto包
)

const (
	// Address gRPC服务地址
	Address = "192.168.0.83:50052"
)

func main() {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.crt", "mimi.com")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := pb.NewHelloClient(conn)
	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	print(res.Message)
}
