package handler

import (
	"net/http"

	"github.com/kalpesh172000/gsaras/services/common/genproto/orders"
	"github.com/kalpesh172000/gsaras/services/common/util"
	"github.com/kalpesh172000/gsaras/services/orders/types"
)

type OrderHTTPHandler struct {
	orderService types.OrderService
	/* orders.UnimplementedOrderServiceServer */
}

func NewHTTPOrderHandler(orderService types.OrderService) *OrderHTTPHandler {
	handler := &OrderHTTPHandler{
		orderService: orderService,
	}

	return handler
}

func (h *OrderHTTPHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest

	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
	}

	order := &orders.Order{
		OrderId:    42,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
	}

		res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
