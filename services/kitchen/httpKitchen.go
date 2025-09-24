package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/kalpesh172000/gsaras/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (h *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGrpcClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelation := context.WithTimeout(r.Context(), time.Second*2)
		defer cancelation()

		c := orders.NewOrderServiceClient(conn)
		// this is what it means to do the internal grpc call
		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerId: 1,
			ProductId:  123,
			Quantity:   3,
		})
		if err != nil {
			log.Println("failed to call internal service : ", err)
		}

		res, err := c.GetOrders(ctx, &orders.GetOrderRequest{
			CustomerId: 1,
		})
		if err != nil {
			log.Println("failed to call internal service : ", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalln("Template Error: ",err)
		}

	})

	log.Print("Starting Server On", h.addr)
	return http.ListenAndServe(h.addr, router)
}

var ordersTemplate = `
	<!DOCTYPE html> 
	<html>
	<head>
<title>Kitchen Orders</title>
	</head>
	<body>

	<h1>Orders List</h1>
	<table order="1">
	<tr>
	<th>Order ID</th>
	<th>Custmer ID</th>
	<th>Quantity</th>
	</tr>

	{{range.}}
	<tr>
	<td>{{.OrderId}}</td>
	<td>{{.CustomerId}}</td>
	<td>{{.Quantity}}</td>
	</tr>
	{{end}}

	</table>
	</body> 
	</html>
	`
