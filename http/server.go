package echo

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/omeid/echo"
)

// NewServer returns a new http echo server.
func NewServer(reverse bool) http.Handler {
	return &echoServer{reverse: reverse}
}

type echoServer struct {
	reverse bool
	count   int64
}

func (e *echoServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path, method := req.URL.Path, req.Method
	switch {
	case path == "/echo" && method == http.MethodPost:
		e.echoHandler(w, req)
	case path == "/count" && method == http.MethodGet:
		e.countHandler(w, req)
	default:
		http.NotFound(w, req)
	}
}

func (e *echoServer) echoHandler(w http.ResponseWriter, req *http.Request) {
	msg := &echo.Message{}
	err := decode(w, req, msg)

	if err != nil {
		return
	}

	if e.reverse {
		msg.Value = echo.Reverse(msg.Value)
	}

	reply(w, msg)
	atomic.AddInt64(&e.count, 1)
}

func (e *echoServer) countHandler(w http.ResponseWriter, req *http.Request) {
	reply(w, &echo.Count{
		Value: atomic.LoadInt64(&e.count),
	})
}

func decode(w http.ResponseWriter, req *http.Request, target interface{}) error {
	err := json.NewDecoder(req.Body).Decode(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return err
}

func reply(w http.ResponseWriter, msg interface{}) {
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
}
