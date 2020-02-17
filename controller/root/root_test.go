package root

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler did not return correct HTTP status code. Want %v, want %v", http.StatusOK, status)
	}
	expected := "Hello world!"
	trimmedBody := strings.TrimSpace(responseRecorder.Body.String())
	if trimmedBody != expected {
		t.Errorf("Handler did not return expected body. Want %v, got %v", expected, responseRecorder.Body.String())
	}
}
