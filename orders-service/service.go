package main

import (
	"context"

	pb "github.com/Desk888/common/api"
	"github.com/google/uuid"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	validItems, err := s.ValidateOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	order := &pb.Order{
		Id:         uuid.New().String(),
		CustomerId: req.CustomerId,
		Status:     "PENDING",
		Items:      validItems,
	}

	err = s.store.SaveOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *service) GetOrder(ctx context.Context, orderId string) (*pb.Order, error) {
	order, err := s.store.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.ItemsWithQuantity, error) {
	validItems := make([]*pb.ItemsWithQuantity, 0)
	for _, item := range p.Items {
		if item.Quantity > 0 {
			validItems = append(validItems, &pb.ItemsWithQuantity{Id: item.Id, Quantity: item.Quantity})
		}
	}
	return validItems, nil
}

func (s *service) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	existingOrder, err := s.store.GetOrderById(o.Id)
	if err != nil {
		return nil, err
	}

	existingOrder.Items = o.Items
	err = s.store.SaveOrder(existingOrder)
	if err != nil {
		return nil, err
	}
	return existingOrder, nil
}

func (s *service) DeleteOrder(ctx context.Context, orderId string) error {
	_, err := s.store.GetOrderById(orderId)
	if err != nil {
		return err
	}

	err = s.store.DeleteOrder(orderId)
	if err != nil {
		return err
	}
	return nil
}