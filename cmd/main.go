package main

import (
	"github.com/thaniri/folrep/controller"
	"net/http"
	"time"
)

func main() {
	controller := controller.New()

	webApp := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      controller,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	webApp.ListenAndServe()
}
