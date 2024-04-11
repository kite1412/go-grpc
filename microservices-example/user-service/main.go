package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	order "my.go/microservicesgrpc/service/order/generated"
	"my.go/microservicesgrpc/service/user/handler"
)

var (
	orderClient order.OrderServiceClient
)

func main() {
	c := grpc.WithTransportCredentials(insecure.NewCredentials())
	con, err := grpc.Dial("localhost:9000", c)
	if err != nil {
		log.Fatal("can't connect to order's rpc server")
	}
	orderClient = order.NewOrderServiceClient(con)
	handler.Init(orderClient)
	http.ListenAndServe(":8080", nil)
}