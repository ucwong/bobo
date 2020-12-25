package main

import (
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	badger "github.com/dgraph-io/badger/v2"
)

var db *badger.DB

func main() {
	bg, err := badger.Open(badger.DefaultOptions("badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer bg.Close()

	bak, err := ioutil.TempFile("badger", "badger.bak")
	_, err = bg.Backup(bak, 0)
	defer bak.Close()

	db = bg

	fmt.Println("Badger started")

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

	return get(k)
}

func Set(k, v string) error {
	fmt.Println("Do set key=" + k + ", value=" + v)

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
		item, err := txn.Get([]byte(k))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		v = string(val)
		return nil
	})

	return
}
