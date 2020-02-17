package arbitrary

import (
	"fmt"
	"net/http"
)

func ArbitraryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Arbitrary!")
}
