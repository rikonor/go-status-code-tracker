Status Code Tracker
---

Small utility for getting the status code of a finished http request.

### Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	status "github.com/rikonor/go-status-code-tracker"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	http.Handle("/", middleware(r))
	http.ListenAndServe(":8080", nil)
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wrap the ResponseWriter with a StatusCodeTracker
		t := status.Track(w)

		h.ServeHTTP(t, r)

		// Can now get the response status code
		fmt.Println(t.Status())
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "error", 400)
}
```
