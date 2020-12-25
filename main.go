package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path[1:]
	q := r.URL.Query()

	var res = "suc"
	switch method {
	case "get":
		res = Get(q.Get("k"))
	case "set":
		err := Set(q.Get("k"), q.Get("v"))
		if err != nil {
			res = "failed"
		}
	default:
		res = Default()
	}

	fmt.Fprintf(w, res)
}

func Get(k string) string {
	fmt.Println("Do get key=" + k)

	//TODO

	return "Get some value with key=" + k
}

func Set(k, v string) error {
	fmt.Println("Do set key=" + k + ", value=" + v)

	//TODO

	return nil
}

func Default() string {
	return "method not found"
}
