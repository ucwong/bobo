package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CortexFoundation/CortexTheseus/common/hexutil"
	"github.com/CortexFoundation/CortexTheseus/crypto/secp256k1"

	badger "github.com/dgraph-io/badger/v2"
)

var db *badger.DB

func main() {
	if bg, err := badger.Open(badger.DefaultOptions(".badger")); err == nil {
		defer bg.Close()
		db = bg
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	}
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

var (
	msg        = hexutil.MustDecode("0xce0677bb30baa8cf067c88db9811f4333d131bf8bcf12fe7065d211dce971008")

	testsig    = hexutil.MustDecode("0x90f27b8b488db00b00606796d2987f6a5f59ae62ea05effe84fef5b8b0e549984a691139ad57a3f0b906637673aa2f63d1f55cb1a69199d4009eea23ceaddc9301")
	testpubkey = hexutil.MustDecode("0x04e32df42865e97135acfb65f3bae71bdc86f4d49150ad6a440b6f15878109880a0a2b2667f7e725ceea70c673093bf67663e0312623c8e091b13cf2c0f11ef652")
)

func Set(k, v string) error {
	if !VerifySignature(testpubkey, msg, testsig[:len(testsig)-1]) {
		fmt.Println("signature unpassed")
		return errors.New("signature failed")
	}
	fmt.Println("signature passed")
	return set(k, v)
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

func set(k, v string) (err error) {
	if len(k) == 0 || len(v) == 0 {
		return
	}
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(k), []byte(v))
	})
	return
}

func VerifySignature(pubkey, hash, signature []byte) bool {
	return secp256k1.VerifySignature(pubkey, hash, signature)
}
