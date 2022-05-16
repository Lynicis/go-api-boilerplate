package main

import (
	"sync"

	"go-rest-api-boilerplate/pkg/http_server"
	"go-rest-api-boilerplate/pkg/logger"
	"go-rest-api-boilerplate/pkg/rpc_server"
)

func startServer(
	httpServer http_server.Server,
	rpcServer rpc_server.RPCServer,
	log logger.Logger,
) {
	var err error

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)

	go func() {
		err = httpServer.Start()
		if err != nil {
			log.Fatal(err)
		}

		waitGroup.Done()
	}()

	go func() {
		err = rpcServer.Start()
		if err != nil {
			log.Fatal(err)
		}

		waitGroup.Done()
	}()

	log.Info("Application started")
	waitGroup.Wait()
}
