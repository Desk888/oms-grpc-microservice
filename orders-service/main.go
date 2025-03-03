package main

import (
	"context"
	"log"
	"net"

	"github.com/Desk888/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()
	db := NewDB()
	svc := NewService(db)
	svc.CreateOrder(context.Background())
	NewGRPCHandler(grpcServer)

	log.Println("Started gRPC server on", grpcAddr)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

