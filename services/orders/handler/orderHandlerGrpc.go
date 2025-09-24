package handler

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"github.com/kalpesh172000/gsaras/services/common/genproto/orders"
	"github.com/kalpesh172000/gsaras/services/orders/types"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpcServer *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpcServer, gRPCHandler)
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}


func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, request *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	// data
	order := &orders.Order{
		OrderId: 42,
		CustomerId: 2,
		ProductId: 1,
		Quantity: 10,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		log.Fatal("error occured",err)
		return nil,err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res,nil
}
