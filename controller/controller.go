package controller

import (
	"github.com/gorilla/mux"
	"github.com/thaniri/folrep/controller/arbitrary"
	"github.com/thaniri/folrep/controller/metrics"
	"github.com/thaniri/folrep/controller/root"
	"net/http"
)

func New() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", metrics.MetricsHandler).Methods(http.MethodGet)
	router.HandleFunc("/arbitrary", arbitrary.ArbitraryHandler).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/", root.IndexHandler).Methods(http.MethodGet, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))
	return router
}
