package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/metrics", nil)

	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MetricsHandler)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler did not return correct HTTP status code. Want %v, want %v", http.StatusOK, status)
	}
}
