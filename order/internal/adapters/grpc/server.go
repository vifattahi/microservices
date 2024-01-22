package grpc

import (
	"fmt"
	"net"

	"GitHub.com/vifattahi/microservices/order/internal/adapters/grpc"
	"GitHub.com/vifattahi/microservices/order/internal/ports"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{{api: api, port: port}}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
}
