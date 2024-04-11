module my.go/microservicesgrpc/service/user

go 1.22.1

require my.go/microservicesgrpc/service/order v1.0.0

require (
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace my.go/microservicesgrpc/service/order => ../order-service
