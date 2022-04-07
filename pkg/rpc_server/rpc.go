package rpcserver

import (
	"fmt"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
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
	config configmodel.RPCServer
}

func NewRPCServer(serverConfig configmodel.RPCServer) RPCServer {
	grpcInstance := grpc.NewServer()

	return &rpcServer{
		config: serverConfig,
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

	serverAddress := fmt.Sprintf(":%d", rpcServer.config.Port)
	tcpServer, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return err
	}

	return rpcServer.server.Serve(tcpServer)
}
