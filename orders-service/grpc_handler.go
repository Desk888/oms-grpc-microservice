package main

import (
	"context"
	"log"

	pb "github.com/Desk888/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	svc OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, svc OrdersService) *grpcHandler {
	handler := &grpcHandler{svc: svc}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return handler
}

func (h *grpcHandler) CreateOrder(ctx context.Context, r *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received with customer ID:", r.CustomerId)
	order, err := h.svc.CreateOrder(ctx, r)
	if err != nil {
		return nil, err
	}
	return order, nil
}
