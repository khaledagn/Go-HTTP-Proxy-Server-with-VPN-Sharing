package httpserver

import (
	"agn-proxy-server/reverseproxy"
	"context"
	"log"
	"net/http"
)



// this function starts the HTTP proxy server with a context for graceful shutdown
func StartHTTPServerWithContext(ctx context.Context, port string) {
	server := &http.Server{
		Addr: ":" + port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy := reverseproxy.NewReverseProxy(r.URL)
			proxy.ServeHTTP(w, r)
		}),
	}

	// start the server in a goroutine
	go func() {
		log.Printf("HTTP server listening on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// listen for context cancellation
	<-ctx.Done()


	// shutdown the server gracefully
	log.Println("Shutting down HTTP server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
}
