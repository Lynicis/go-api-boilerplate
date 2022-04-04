package rpcserver

import (
	"fmt"
	"go-rest-api-boilerplate/pkg/config"
	"google.golang.org/grpc"
	"net"
)

type RPCServer interface {
	GetRPCServer() *grpc.Server
	StartServer() error
	Stop()
}

type rpcServer struct {
	server *grpc.Server
	config config.Config
}

func NewRPCServer(config config.Config) RPCServer {
	grpcInstance := grpc.NewServer()
	return &rpcServer{
		config: config,
		server: grpcInstance,
	}
}

func (rpcServer *rpcServer) GetRPCServer() *grpc.Server {
	return rpcServer.server
}

func (rpcServer *rpcServer) StartServer() error {
	rpcServerPort := fmt.Sprintf(":%d", rpcServer.config.GetRPCConfig().Port)
	tcpServer, err := net.Listen("tcp", rpcServerPort)
	if err != nil {
		return err
	}

	return rpcServer.server.Serve(tcpServer)
}

func (rpcServer *rpcServer) Stop() {
	rpcServer.server.GracefulStop() //todo: need to graceful shutdown
}
