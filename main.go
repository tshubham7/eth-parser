package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/tshubham7/eth-parser/internal/parser/handler"
	"github.com/tshubham7/eth-parser/internal/parser/usecase"
	"github.com/tshubham7/eth-parser/internal/pkg/client"
	"github.com/tshubham7/eth-parser/internal/pkg/db"
	"github.com/tshubham7/eth-parser/internal/pkg/middleware"
	"github.com/tshubham7/eth-parser/internal/pkg/utils"
)

// HTTP Handlers
func main() {
	router := mux.NewRouter()
	v1 := router.PathPrefix("/api/v1").Subrouter()

	gets := v1.Methods(http.MethodGet).Subrouter()
	posts := v1.Methods(http.MethodPost).Subrouter()

	store := db.NewDBStore("default")
	apiClient := client.NewHttpClient()
	parserManager := usecase.NewParserUsecase(store, apiClient)
	parserHandler := handler.NewParserHttpHandler(parserManager)

	gets.HandleFunc("/transactions", parserHandler.RequestGetTransactions)
	gets.HandleFunc("/current-block", parserHandler.RequestGetCurrentBlockNumber)
	posts.HandleFunc("/subscribe", parserHandler.RequestPostSubscribe)

	adminV1 := v1.PathPrefix("/admin").Subrouter()
	adminV1.Use(middleware.AuthMiddleware)
	adminGets := adminV1.Methods(http.MethodGet).Subrouter()
	adminGets.HandleFunc("/transactions", parserHandler.RequestAllTransactions)

	closeCh := make(chan struct{})
	ctx := context.Background()
	ctx, _ = utils.NewSilentLogger(ctx)
	go parserManager.Process(ctx, closeCh)

	logrus.Info("starting server...")
	go func() {
		if err := http.ListenAndServe(":8088", router); err != nil {
			logrus.Fatal("failed to start the server, ", err)
		}
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quitChannel
	logrus.Info("quitting server...")
	close(closeCh)
	time.Sleep(time.Second * 3)
}
