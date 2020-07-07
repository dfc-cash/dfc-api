// Package dfccore is a client library for communicating with an
// instance of dfc-core using through the server's HTTP port.
package dfccore

import "net/http"

// SetCursorDone is the success mdfcge returned by dfc-core when a cursor
// update succeeds.
const SetCursorDone = "Done"

// HTTP represents the http client that a dfccore client uses to make http
// requests.
type HTTP interface {
	Do(req *http.Request) (*http.Response, error)
}

// confirm interface conformity
var _ HTTP = http.DefaultClient
