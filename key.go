package main

import "encoding/base64"

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

type RawKeyPair struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type KeyObject struct {
	AuthKeys     *KeyPair
	ExchangeKeys *KeyPair
	Address      []byte
}

type RawKeyObject struct {
	AuthKeys     *RawKeyPair `json:"authKeys"`
	ExchangeKeys *RawKeyPair `json:"exchangeKeys"`
	Address      string      `json:"address"`
}

// Returns constructed RawKeyPair
func (kp *KeyPair) ToRaw() (rkp *RawKeyPair) {
	rkp = &RawKeyPair{}
	rkp.PrivateKey = base64.RawURLEncoding.EncodeToString(kp.PrivateKey)
	rkp.PublicKey = base64.RawURLEncoding.EncodeToString(kp.PublicKey)
	return rkp
}

// Returns constructed RawKeyObject
func (ko *KeyObject) ToRaw() (rko *RawKeyObject) {
	rko = &RawKeyObject{}
	rko.Address = base64.RawURLEncoding.EncodeToString(ko.Address)

	rko.AuthKeys = ko.AuthKeys.ToRaw()
	rko.ExchangeKeys = ko.ExchangeKeys.ToRaw()
	return rko
}
