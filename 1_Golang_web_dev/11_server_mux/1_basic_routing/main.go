package main

import (
	"io"
	"net/http"
)

type something int

func (s something) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// this is not the best solution to handle the routing
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "wouf wouf")
	case "/cat":
		io.WriteString(res, "miu miu")
	}
}

func main() {
	var s something
	http.ListenAndServe(":3000", s)
}
