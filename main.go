package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)

	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/set", SetHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get invoke, %s!", r.URL.Path[1:])
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Set invoke %s!", r.URL.Path[1:])
}
