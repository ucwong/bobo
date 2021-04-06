package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ucwong/golang-kv"
)

var db kv.Bucket

func main() {
	db = kv.Badger(".badger")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v, %v, %v\n", r.URL, r.Method, r.URL.Path)
	res := "OK"
	uri := r.URL.Path
	switch r.Method {
	case "GET":
		res = Get(uri)
	case "POST":
		if reqBody, err := ioutil.ReadAll(r.Body); err == nil {
			if err := Set(uri, string(reqBody)); err != nil {
				res = "ERROR" //fmt.Sprintf("%v", err)
			}
		}
	default:
		res = "method not found"
	}
	fmt.Fprintf(w, res)
}

func Get(k string) string {
	return get(k)
}

func Set(k, v string) error {
	return set(k, v)
}

func get(k string) (v string) {
	if len(k) == 0 {
		return
	}
	v = string(db.Get([]byte(k)))
	return
}

func set(k, v string) (err error) {
	if len(k) == 0 || len(v) == 0 {
		return
	}

	err = db.Set([]byte(k), []byte(v))

	return
}
