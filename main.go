package main

import (
	"context"
	"syscall"
)

func main() {
	// Startup configuration
	config := new(appConfigData)
	config.loadAndConfigure()

	rootCtx := context.Background()
	sigCtx, cancel := context.WithCancel(rootCtx)

	// Setup signal handler goroutine and get a channel
	// to receive signals caught by the handler.
	recvSigCh := setupSignalHandling(sigCtx)

	for {
		server := startHttpServer(config)

		sig := <-recvSigCh
		switch sig {
		case syscall.SIGHUP:
			if err := server.Shutdown(rootCtx); err != nil {
				panic(err)
			}
			config = new(appConfigData)
			config.loadAndConfigure()

		case syscall.SIGTERM, syscall.SIGINT:
			cancel()
			if err := server.Shutdown(rootCtx); err != nil {
				panic(err)
			}
			return
		}
	}
}
