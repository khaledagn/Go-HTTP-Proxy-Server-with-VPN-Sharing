package proxyserver

import (
	"agn-proxy-server/internal/config"
	"agn-proxy-server/internal/httpserver"
	"agn-proxy-server/internal/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	globalCancel context.CancelFunc
	wg           sync.WaitGroup
	isRunning    bool
)
 
// this function starts the HTTP proxy server in detached mode
func StartProxy(port string) string {
	logger.InitLogger()

	// update the configuration with the provided port
	cfg := config.GetConfig()
	cfg.HTTPPort = port

	// initialize contexts for the server
	ctx, cancel := context.WithCancel(context.Background())
	globalCancel = cancel 

	isRunning = true

	// start the HTTP server in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpserver.StartHTTPServerWithContext(ctx, cfg.HTTPPort)
	}()

	return fmt.Sprintf("Proxy Server running on HTTP: %s", cfg.HTTPPort)
}

// this function stops the HTTP proxy server
func StopProxy() string {
	if globalCancel != nil {
		globalCancel()
	}

	// wait for goroutines to exit
	wg.Wait()

	isRunning = false
	return "Proxy Server has been stopped."
}

// this function starts the proxy server in a detached mode and listens for termination signals
func RunDetachedProxy() {
	logger.InitLogger()

	if isRunning {
		logger.Info("Proxy is already running")
		return
	}

	// handle OS signals for graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// start the proxy
	startMsg := StartProxy(config.GetConfig().HTTPPort)
	logger.Info(startMsg)

	// wait for termination signals
	select {
	case <-signalChan: 
		logger.Info("Received termination signal, stopping proxy...")
		stopMsg := StopProxy()
		logger.Info(stopMsg)
	case <-context.Background().Done(): 
		logger.Info("Stopping proxy via external signal...")
		stopMsg := StopProxy()
		logger.Info(stopMsg)
	}
}

// this func checks if the proxy is running
func IsProxyRunning() bool {
	return isRunning
}
