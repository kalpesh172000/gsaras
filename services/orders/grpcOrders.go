package main

import (
	"fmt"
	"net"
	"log"

	"google.golang.org/grpc"
	"github.com/kalpesh172000/gsaras/services/orders/services"
	"github.com/kalpesh172000/gsaras/services/orders/handler"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}


func (g *gRPCServer) Run() error{
	lis, err := net.Listen("tcp", g.addr)
	if err != nil {
		fmt.Println("error occured",err)
	}

	grpcServer := grpc.NewServer()

	orderService := services.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Server started on",g.addr)
	return grpcServer.Serve(lis)
}

