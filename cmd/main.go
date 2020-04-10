package main

import (
	"go.uber.org/zap"
	"github.com/thaniri/folrep/controller"
	"net/http"
	"time"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()
	logger.Info("hahaha")
	controller := controller.New(*logger)

	webApp := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      controller,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	webApp.ListenAndServe()
}
