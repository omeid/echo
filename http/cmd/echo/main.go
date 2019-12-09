package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	echo "github.com/omeid/echo/http"
)

var (
	host    string
	port    string
	reverse bool
	wait    time.Duration
)

func init() {
	flag.StringVar(&host, "host", "", "the host to listen on")
	flag.StringVar(&port, "port", "3000", "the port to listen on")
	flag.BoolVar(&reverse, "reverse", false, "whatever echo replies should be reversed.")
	flag.DurationVar(&wait, "wait", 0, "Wait before listening")
}

func main() {

	flag.Parse()

	time.Sleep(wait)

	addr := fmt.Sprintf("%s:%s", host, port)
	server := echo.NewServer(reverse)

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}
