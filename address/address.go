package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/iocn-io/ripemd160"
	"github.com/mr-tron/base58"
)

type WalletKeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWalletKeyPair() *WalletKeyPair {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKeyRow := privateKey.PublicKey
	publicKey := append(publicKeyRow.X.Bytes(), publicKeyRow.Y.Bytes()...)
	return &WalletKeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func (w *WalletKeyPair) GetAddress() string {
	pub := w.PublicKey
	h1 := sha256.Sum256(pub)
	ri := ripemd160.New()
	ri.Write(h1[:])
	h2 := ri.Sum(nil)
	version := 0x00

	lb := append([]byte{byte(version)}, h2...)
	fir := sha256.Sum256(lb)
	se := sha256.Sum256(fir[:])
	checkRight := se[0:4]

	address := append(lb, checkRight...)

	return base58.Encode(address)

}

func IsValidAddress(address string) bool {
	de, _ := base58.Decode(address)
	r := de[:len(de)-4]
	f := sha256.Sum256(r)
	s := sha256.Sum256(f[:])

	ri := s[:4]
	li := de[len(de)-4:]
	return bytes.Equal(ri, li)
}
