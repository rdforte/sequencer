package main

import (
	"context"
	"fmt"
	"github.com/rdforte/sequencer/internal/config"
	"github.com/rdforte/sequencer/internal/handlers"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var buildEnv = "dev"

func main() {
	if err := run(); err != nil {
		fmt.Printf("server error %v", err)
		os.Exit(1)
	}
}

func run() error {
	log.Printf("Setting up server...")

	cfg, err := config.CreateAPIConfig()
	if err != nil {
		return fmt.Errorf("unable to create application config error %w", err)
	}

	/** The Debug function returns a mux to listen and serve on for all the debug
	related endpoints. This includes the standard library endpoints.
	*/
	debugMux := handlers.CreateDebugMux(buildEnv)

	// start the service listening for debug requests.
	// not concerned about shutting this down with load shedding.
	go func() {
		log.Printf("debug service started on port %v", cfg.DebugPort())
		if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.DebugPort()), debugMux); err != nil {
			fmt.Printf("debug server shutdown on host %v, error %v", cfg.DebugPort(), err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	apiMux := handlers.CreateAPIMux(buildEnv)

	srv := &http.Server{
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Addr:         fmt.Sprintf(":%v", cfg.ApiPort()),
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      apiMux,
	}

	srvErr := make(chan error, 1)
	go func() {
		log.Printf("service started on port %v", cfg.ApiPort())
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err := <-srvErr:
		// Error when starting HTTP server.
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		// give the server 10 seconds to shut down gracefully. Allows for load shedding.

		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout())
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	return nil
}
