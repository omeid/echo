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

// Reverse the input string.
func Reverse(input string) string {
	//This solution comes from no one but Russ Cox himself:
	// https://groups.google.com/d/msg/golang-nuts/oPuBaYJ17t4/PCmhdAyrNVkJ

	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
