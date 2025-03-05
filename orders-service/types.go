package main

import (
	"context"

	pb "github.com/Desk888/common/api"
)

type OrdersService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
	GetOrder(context.Context, string) (*pb.Order, error)
	UpdateOrder(context.Context, *pb.Order) (*pb.Order, error)
}

type OrdersStore interface {
	Create(context.Context) error
	SaveOrder(order *pb.Order) error
	GetOrderById(id string) (*pb.Order, error)
}
