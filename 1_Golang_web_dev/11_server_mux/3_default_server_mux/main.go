package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This si my route one")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This si my route two")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	// we can implement the following as well
	// http.Handle("/dog", http.HandlerFunc(d))
	// http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8000", nil)
}
