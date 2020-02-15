package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func New() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)
	return router
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
