package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ucwong/golang-kv"
)

var db kv.Bucket

func main() {
	db = kv.Badger("")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v, %v, %v\n", r.URL, r.Method, r.URL.Path)
	res, uri := "OK", r.URL.Path
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
	return string(db.Get([]byte(k)))
}

func Set(k, v string) error {
	return db.Set([]byte(k), []byte(v))
}
