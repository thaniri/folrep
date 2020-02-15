package root

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello world!")
}
