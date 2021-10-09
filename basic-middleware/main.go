package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(rw, r)
	}
}

func foo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, r.URL.Path)
}

func bar(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, r.URL.Path)
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
