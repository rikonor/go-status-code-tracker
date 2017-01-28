package status

import (
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestStatusCodeTracker(t *testing.T) {
	codes := []int{200, 500}

	// setup some test routes
	for _, code := range codes {
		hfn := middleware(t, code, handleWithCode(code))

		http.Handle("/"+strconv.Itoa(code), hfn)
	}

	// start server
	go http.ListenAndServe(":12345", nil) // assuming 12345 is available

	// wait for server to start listening
	time.Sleep(10 * time.Millisecond)

	// trigger the routes and their respective checks
	for _, code := range codes {
		_, err := http.Get("http://localhost:12345/" + strconv.Itoa(code))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func handleWithCode(code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "error", code)
	}
}

func middleware(t *testing.T, expected int, hfn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// wrap `w` with a status code tracker
		tracker := Track(w)

		// call original handler
		hfn(tracker, r)

		// check if the tracker got the expected status code
		if tracker.Status() != expected {
			t.Fatalf("wrong status: %d, expected %d", tracker.Status(), expected)
		}
	}
}
