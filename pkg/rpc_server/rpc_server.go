package rpcserver

import (
	"fmt"
	"go-rest-api-boilerplate/pkg/logger"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

// RPCServer interface provide getter and starting methods
type RPCServer interface {
	GetRPCServer() *grpc.Server
	Start()
}

type rpcServer struct {
	server *grpc.Server
	config configmodel.RPCServer
	logger logger.Logger
}

// NewRPCServer create new rpc instance
func NewRPCServer(serverConfig configmodel.RPCServer, logger logger.Logger) RPCServer {
	grpcInstance := grpc.NewServer()

	return &rpcServer{
		config: serverConfig,
		server: grpcInstance,
		logger: logger,
	}
}

func (rpcServer *rpcServer) GetRPCServer() *grpc.Server {
	return rpcServer.server
}

func (rpcServer *rpcServer) Start() {
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChannel
		rpcServer.server.Stop()
	}()

	serverAddress := fmt.Sprintf(":%d", rpcServer.config.Port)
	tcpServer, err := net.Listen("tcp", serverAddress)
	if err != nil {
		rpcServer.logger.Fatal("got error while TCP server starting", zap.Error(err))
	}

	err = rpcServer.server.Serve(tcpServer)
	if err != nil {
		rpcServer.logger.Fatal("got error while gRPC server starting", zap.Error(err))
	}
}
