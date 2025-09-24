// Package services...
package services

import (
	"context"
	"github.com/kalpesh172000/gsaras/services/common/genproto/orders"

)

var ordersStore= make([]*orders.Order, 0)

type OrderService struct {
	//storage
}


func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersStore = append(ordersStore, order)
	return nil
}


func (s *OrderService) GetOrders (ctx context.Context) []*orders.Order {
	return ordersStore
}
