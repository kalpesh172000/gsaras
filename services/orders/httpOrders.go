package main

import (
	"log"
	"net/http"

	"github.com/kalpesh172000/gsaras/services/orders/handler"
	"github.com/kalpesh172000/gsaras/services/orders/services"
)

type hTTPServer struct {
	addr string
}

func NewHttpServer(addr string) *hTTPServer {
	return &hTTPServer{addr: addr}
}

func (h *hTTPServer) Run() error {
	router := http.NewServeMux()

	orderService := services.NewOrderService()
	orderHandler := handler.NewHTTPOrderHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Server started on", h.addr)

	return http.ListenAndServe(h.addr, router)
}
