package main

import (
	"fmt"
	"net/http"

	badger "github.com/dgraph-io/badger/v2"
)

var db *badger.DB

func main() {
	if bg, err := badger.Open(badger.DefaultOptions(".badger")); err == nil {
		defer bg.Close()
		db = bg

		http.HandleFunc("/", Handler)
		http.ListenAndServe(":8080", nil)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path[1:]
	q := r.URL.Query()

	var res = "OK"
	switch method {
	case "get":
		res = Get(q.Get("k"))
	case "set":
		err := Set(q.Get("k"), q.Get("v"))
		if err != nil {
			res = "ERR" //fmt.Sprintf("%v", err)
		}
	default:
		res = Default()
	}
	fmt.Fprintf(w, res)
}

func Get(k string) string {
	fmt.Println("Do get [k=" + k + "]")
	return get(k)
}

func Set(k, v string) error {
	fmt.Println("Do set [k=" + k + ",v=" + v + "]")
	return set(k, v)
}

func Default() string {
	return "method not found"
}

func set(k, v string) (err error) {
	if len(k) == 0 || len(v) == 0 {
		return
	}
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(k), []byte(v))
	})
	return
}

func get(k string) (v string) {
	if len(k) == 0 {
		return
	}
	db.View(func(txn *badger.Txn) error {
		if item, err := txn.Get([]byte(k)); err == nil {
			if val, err := item.ValueCopy(nil); err == nil {
				v = string(val)
			}
		}
		return nil
	})
	return
}
