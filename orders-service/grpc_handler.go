package main 

import (
	"context"
	"log"
	pb "github.com/Desk888/common/api"
	"google.golang.org/grpc"
)

func NewGRPCHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return handler
}

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func (h *grpcHandler) CreateOrder(ctx context.Context, r *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received")
	return &pb.Order{}, nil
}