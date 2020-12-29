package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/CortexFoundation/CortexTheseus/common"
	"github.com/CortexFoundation/CortexTheseus/common/hexutil"
	"github.com/CortexFoundation/CortexTheseus/common/math"
	"github.com/CortexFoundation/CortexTheseus/crypto"
	"github.com/CortexFoundation/CortexTheseus/crypto/secp256k1"

	"golang.org/x/crypto/sha3"

	badger "github.com/dgraph-io/badger/v2"
)

type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

var (
	db            *badger.DB
	secp256k1N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
)

const DigestLength = 32

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
	q := r.URL.Query()
	switch r.Method {
	case "GET":
		res = Get(uri)
	case "POST":
		if reqBody, err := ioutil.ReadAll(r.Body); err == nil {
			u := strings.Split(uri, "/")
			if len(u) >= 0 {
				if err := Set(uri, string(reqBody), u[len(u)-1], q.Get("msg"), q.Get("sig")); err != nil {
					res = "ERROR" //fmt.Sprintf("%v", err)
				}
			}
		}
	default:
		res = "Method not found"
	}
	fmt.Fprintf(w, res)
}

func Get(k string) string {
	return get(k)
}

var (
//msg = hexutil.MustDecode("0xce0677bb30baa8cf067c88db9811f4333d131bf8bcf12fe7065d211dce971008")
//testsig    = hexutil.MustDecode("0x90f27b8b488db00b00606796d2987f6a5f59ae62ea05effe84fef5b8b0e549984a691139ad57a3f0b906637673aa2f63d1f55cb1a69199d4009eea23ceaddc9301")
//testpubkey = hexutil.MustDecode("0x04e32df42865e97135acfb65f3bae71bdc86f4d49150ad6a440b6f15878109880a0a2b2667f7e725ceea70c673093bf67663e0312623c8e091b13cf2c0f11ef652")
)

func Set(k, v, addr, msg, sig string) error {

	//m := hexutil.MustDecode(msg)
	m := Keccak256([]byte(msg))
	//p := hexutil.MustDecode(pub)
	s := hexutil.MustDecode(sig)

	if len(m) == 0 || len(s) == 0 {
		return errors.New("Hex decode failed")
	}

	//	fmt.Println(hexutil.Encode(m[:]))

	recoveredPub, err := Ecrecover(m, s)
	if err != nil {
		return errors.New("Ecrecover failed")
	}

	pubKey, _ := UnmarshalPubkey(recoveredPub)
	recoveredAddr := PubkeyToAddress(*pubKey)
	if common.HexToAddress(addr) != recoveredAddr {
		fmt.Printf("Address mismatch: want: %v have: %v\n", addr, recoveredAddr.Hex())
		return errors.New("Key mismatched")
	}

	if !VerifySignature(recoveredPub, m, s[:len(s)-1]) {
		fmt.Println("Signature unpassed")
		return errors.New("Signature failed")
	}

	fmt.Println("signature passed")

	//ss , _ := SignHex("Hello", "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	// sig 0xee78eaa27526b412d0e970b85f47c96aa0aa67ed1c06f577ffe712a91284659a0a38529194a53891c84919369e09bf7e08d1655544cb044671461e210ddad1eb00
	// pub
	//fmt.Println(hexutil.Encode(ss[:]))
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

func EcRecover(data, sig hexutil.Bytes) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Cortex signature (V is not 27 or 28)")
	}
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	hash, _ := SignHash(data)
	rpk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}

// SignHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Cortex Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func SignHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Cortex Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg)), msg
}

func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := FromECDSAPub(&p)
	return common.BytesToAddress(Keccak256(pubBytes[1:])[12:])
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(S256(), pub.X, pub.Y)
}

func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := sha3.NewLegacyKeccak256().(KeccakState)
	for _, b := range data {
		d.Write(b)
	}
	d.Read(b)
	return b
}

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return secp256k1.RecoverPubkey(hash, sig)
}

func UnmarshalPubkey(pub []byte) (*ecdsa.PublicKey, error) {
	x, y := elliptic.Unmarshal(S256(), pub)
	if x == nil {
		return nil, errors.New("invalid key")
	}
	return &ecdsa.PublicKey{Curve: S256(), X: x, Y: y}, nil
}

func S256() elliptic.Curve {
	return secp256k1.S256()
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

//pri := "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
//msg := "Hello"
func SignHex(msg, pri string) (sig []byte, err error) {
	k0, _ := HexToECDSA(pri)
	msg0 := Keccak256([]byte(msg))
	fmt.Println("msg0 " + hexutil.Encode(msg0[:]))
	return Sign(msg0, k0)
}

// Sign calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given hash cannot be chosen by an adversery. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
func Sign(hash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	if len(hash) != DigestLength {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}
	seckey := math.PaddedBigBytes(prv.D, prv.Params().BitSize/8)
	defer zeroBytes(seckey)
	return secp256k1.Sign(hash, seckey)
}

// HexToECDSA parses a secp256k1 private key.
func HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	b, err := hex.DecodeString(hexkey)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in private key", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for private key")
	}
	return ToECDSA(b)
}

// ToECDSA creates a private key with the given D value.
func ToECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	return toECDSA(d, true)
}

// toECDSA creates a private key with the given D value. The strict parameter
// controls whether the key's length should be enforced at the curve size or
// it can also accept legacy encodings (0 prefixes).
func toECDSA(d []byte, strict bool) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = S256()
	if strict && 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	if priv.D.Cmp(secp256k1N) >= 0 {
		return nil, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}
