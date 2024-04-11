package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	order "my.go/microservicesgrpc/service/order/generated"
	"my.go/microservicesgrpc/service/order/handler"
)

var (
	_log = log.New(log.Writer(), "<RPC-SERVER> ", log.Lmsgprefix)
)

func init() {
	handler.HandlerInit()
}

func main() {
	serviceServer := order.NewOrderServiceServer()
	grpcServer := grpc.NewServer()
	order.RegisterOrderServiceServer(grpcServer, serviceServer)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:9000")
	if err != nil {
		_log.Fatal(err)
	}
	listener, lErr := net.ListenTCP("tcp", tcpAddr)
	if lErr != nil {
		_log.Fatal(lErr)
	}
	go grpcServer.Serve(listener)
	http.ListenAndServe(":8081", nil)
}