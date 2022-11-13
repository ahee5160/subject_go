package home_work

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := &HTTPServer{addr: ":8088"}
	server.AddRoute(http.MethodGet, "/ping", func(ctx Context) {
		fmt.Println("pong")
	})
	server.Get("/user", func(ctx Context) {
		fmt.Println("ahe")
	})
	server.Start()
}