package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	mgrpc "my.go/grpc/generated"
)

func main() {
	serviceServer := mgrpc.NewServiceServer()
	grpcServer := grpc.NewServer()
	tcpAddr, aErr := net.ResolveTCPAddr("tcp", "localhost:8080")
	if aErr != nil {
		log.Fatal(aErr)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	mgrpc.RegisterServiceServer(grpcServer, serviceServer)

	go grpcServer.Serve(listener)

	c := grpc.WithTransportCredentials(insecure.NewCredentials())

	connection, dErr := grpc.Dial("localhost:8080", c)
	if dErr != nil {
		log.Fatal(dErr)
	}
	serviceClient := mgrpc.NewServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	req := &mgrpc.Request{
		ReqData: "hello from client",
	}
	ress, rErr := serviceClient.Exchange(ctx, req)
	if rErr != nil {
		log.Fatal(rErr)
	}
	log.Println(ress.ResData)
}