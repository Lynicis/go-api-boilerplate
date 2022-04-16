package rpcserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

// RPCServer gRPC server domain
type RPCServer interface {
	GetRPCServer() *grpc.Server
	Start() error
	Stop()
}

type rpcServer struct {
	server *grpc.Server
	config configmodel.RPCServer
}

// NewRPCServer create new rpc server instance
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

func (rpcServer *rpcServer) Start() error {
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
