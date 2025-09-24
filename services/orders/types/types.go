package types

import (
	"context"
	"github.com/kalpesh172000/gsaras/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order 
}
