package controller

import (
	"go.uber.org/zap"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/thaniri/folrep/controller/metrics"
	"github.com/thaniri/folrep/controller/root"
)

func New(logger zap.SugaredLogger) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", metrics.MetricsHandler)

	router.HandleFunc("/", root.IndexHandler)
	//myLoggingHandler := LoggingHandler(logger, root.IndexHandler) <- this cannot work
	router.HandleFunc("/", root.IndexHandler)
	router.Use(Middleware)
	return router
}

type loggingHandler struct {
	logger zap.SugaredLogger
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("test")
}

func LoggingHandler(logger zap.SugaredLogger, h http.Handler) http.Handler {
	return loggingHandler{logger, h}
}

func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("middleware", r.URL)
        h.ServeHTTP(w, r)
    })
}
