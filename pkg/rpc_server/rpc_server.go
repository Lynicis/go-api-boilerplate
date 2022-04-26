package rpc_server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	config_model "go-rest-api-boilerplate/pkg/config/model"
)

type RPCServer interface {
	GetRPCServer() *grpc.Server
	Start() error
	Stop()
}

type rpcServer struct {
	server *grpc.Server
	config config_model.RPCServer
}

func NewRPCServer(serverConfig config_model.RPCServer) RPCServer {
	grpcInstance := grpc.NewServer()

	return &rpcServer{
		config: serverConfig,
		server: grpcInstance,
	}
}

func (rpcServer *rpcServer) GetRPCServer() *grpc.Server {
	return rpcServer.server
}

func (rpcServer *rpcServer) Start() error {
	var err error

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

func (rpcServer *rpcServer) Stop() {
	rpcServer.server.Stop()
}
