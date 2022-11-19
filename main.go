package main

import (
	"encoding/json"
	"os"
)

func main() {
	var (
		kobj = &KeyObject{}
		err  error
		enc  *json.Encoder
	)

	if kobj.AuthKeys, err = GenerateAuthkeys(); err != nil {
		panic(err)
	}

	if kobj.ExchangeKeys, err = GenerateExchangeKeys(); err != nil {
		panic(err)
	}

	kobj.Address = GenerateAddress(kobj.AuthKeys, kobj.ExchangeKeys)
	enc = json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	enc.Encode(kobj.ToRaw())
}
