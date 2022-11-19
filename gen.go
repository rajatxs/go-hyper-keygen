package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

// Returns keypair of authentication keys
func GenerateAuthkeys() (kp *KeyPair, err error) {
	var gen *ecdsa.PrivateKey

	if gen, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		return nil, err
	}

	kp = &KeyPair{}
	kp.PrivateKey = gen.D.Bytes()
	kp.PublicKey = elliptic.Marshal(elliptic.P256(), gen.X, gen.Y)
	return kp, nil
}

// Returns keypair of exchange keys
func GenerateExchangeKeys() (kp *KeyPair, err error) {
	var gen *ecdsa.PrivateKey

	if gen, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader); err != nil {
		return nil, err
	}

	kp = &KeyPair{}
	kp.PrivateKey = gen.D.Bytes()
	kp.PublicKey = elliptic.Marshal(elliptic.P384(), gen.X, gen.Y)
	return kp, nil
}

// Returns public address from auth and exchange keys
func GenerateAddress(authKey *KeyPair, exchKey *KeyPair) []byte {
	h := sha256.New224()
	h.Write(append(authKey.PublicKey, exchKey.PublicKey...))
	return h.Sum(nil)
}
