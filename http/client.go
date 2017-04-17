package echo

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/omeid/echo"
)

// NewClient returns a new http echo client.
func NewClient(host string) (echo.Echo, error) {

	return &client{
		//TODO: This needs proper handling.
		//      Reason why there is an error on
		//      return.
		echoURL:  host + "/echo",
		countURL: host + "/count",
		host:     host,
	}, nil
}

type client struct {
	host string

	echoURL  string
	countURL string
}

func (c *client) Say(what string) (string, error) {

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(echo.Message{Value: what})
	if err != nil {
		log.Fatal(err) //Should never happen.
	}

	resp, err := http.Post(c.echoURL, "application/json", &buf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	msg := &echo.Message{}

	err = json.NewDecoder(resp.Body).Decode(msg)
	if err != nil {
		return "", err
	}

	return msg.Value, nil
}

func (c *client) Count() (int64, error) {
	resp, err := http.Get(c.countURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	count := &echo.Count{}

	err = json.NewDecoder(resp.Body).Decode(count)
	if err != nil {
		log.Printf("Error %v", err)
		return 0, err
	}

	return count.Value, nil
}
