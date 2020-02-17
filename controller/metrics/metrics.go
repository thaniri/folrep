package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func MetricsHandler(writer http.ResponseWriter, request *http.Request) {
	promhttp.Handler().ServeHTTP(writer, request)
}
