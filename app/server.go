package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/logger"
)

const ServerShutdownTimeout = time.Second

func NewServer(router *mux.Router, port string) *http.Server {
	server := &http.Server{
		Handler:      router,
		Addr:         net.JoinHostPort("0.0.0.0", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server
}

func (a *App) RunServer() error {
	return a.server.ListenAndServe()
}

func (a *App) StopServer() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), ServerShutdownTimeout)
	defer cancel()

	logger.Info("Closing DB connections....")
	a.Store.Close(context.TODO())

	a.server.Shutdown(ctx)
	logger.Info("Server is stopped")

	appCount--
	os.Exit(0)
}

func (a *App) InitServer() {
	go func() {
		err := a.RunServer()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	logger.Info("Server is listening on " + a.server.Addr)

	a.StopServer()
}
