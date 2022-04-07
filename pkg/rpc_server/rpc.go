package rpcserver

import (
	"fmt"
	"go-rest-api-boilerplate/pkg/config"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type RPCServer interface {
	GetRPCServer() *grpc.Server
	StartServer() error
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
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		rpcServer.server.Stop()
	}()

	rpcConfig := rpcServer.config.GetRPCConfig()
	tcpServer, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcConfig.Port))
	if err != nil {
		return err
	}

	return rpcServer.server.Serve(tcpServer)
}
