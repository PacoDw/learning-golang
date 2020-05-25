package main

import (
	"io"
	"net/http"
)

type routeOne int

func (r routeOne) ServeHTTP(res http.ResponseWriter, requ *http.Request) {
	io.WriteString(res, "This si my route one")
}

type routeTwo int

func (r routeTwo) ServeHTTP(res http.ResponseWriter, requ *http.Request) {
	io.WriteString(res, "This si my route two")
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/dog", new(routeOne))
	mux.Handle("/cat", new(routeOne))

	http.ListenAndServe("3000", mux)
}
