package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	echo "github.com/omeid/echo/http"
)

var (
	host    string
	port    string
	reverse bool
)

func init() {
	flag.StringVar(&host, "host", "", "the host to listen on")
	flag.StringVar(&port, "port", "3000", "the port to listen on")
	flag.BoolVar(&reverse, "reverse", false, "whatever echo replies should be reversed.")
}

func main() {

	flag.Parse()

	addr := fmt.Sprintf("%s:%s", host, port)
	server := echo.NewServer(reverse)

	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}
