// Package echo describes a simple echo server.
package echo

// Echo is a simple service used for testing conex
// it will connect to an Echo Container and does what
// an echo server does, say it back.
// You can also run any executable and get the result
// for more comperhensive testing.
type Echo interface {
	Say(what string) (string, error)
	Count() (int64, error)
}

// Message represents a message send for echo.
type Message struct {
	Value string `json:"value"`
}

// Count represents the count response.
type Count struct {
	Value int64
}
