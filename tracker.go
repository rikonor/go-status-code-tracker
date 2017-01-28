package status

import "net/http"

// StatusCodeTracker wraps a ResponseWriter
// and keeps track of the returned status code
type StatusCodeTracker struct {
	w      http.ResponseWriter
	status int
}

// Status returns the saved status code
// it should only be called once the handler has finished returned
func (t *StatusCodeTracker) Status() int {
	return t.status
}

// Track a ResponseWriter for its StatusCode
func Track(w http.ResponseWriter) *StatusCodeTracker {
	return &StatusCodeTracker{
		w: w,

		// A status code of 200 is assumed unless WriteHeader is called
		status: 200,
	}
}

// Header calls the underlying Header method
func (t *StatusCodeTracker) Header() http.Header {
	return t.w.Header()
}

// Write calls the underlying Write method
func (t *StatusCodeTracker) Write(p []byte) (n int, err error) {
	return t.w.Write(p)
}

// WriteHeader calls the underlying WriteHeader method
// but also saves the status code of the response
func (t *StatusCodeTracker) WriteHeader(code int) {
	t.w.WriteHeader(code)
	t.status = code
}
