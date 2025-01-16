package main

import (
	"agn-proxy-server/proxyserver"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	
	// start the proxy server
	fmt.Println(proxyserver.StartProxy())

	// wait for an interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// block until a signal is received
	<-quit

	// stop the proxy server
	fmt.Println(proxyserver.StopProxy())
	fmt.Println("Proxy server has been stopped gracefully.")
}
