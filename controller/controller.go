package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/thaniri/folrep/controller/metrics"
	"github.com/thaniri/folrep/controller/root"
)

func New() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", metrics.MetricsHandler)

	router.HandleFunc("/", root.IndexHandler)
	return router
}

