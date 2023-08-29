package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	// para omitir o campo na hora de fazer o bind, basta colocar `json:"-"`
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{123, 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	var conta2 Conta
	// caso o json não esteja com os atributos corretos, vai retornar 0, pois não vai conseguir fazer o bind
	err = json.Unmarshal(res, &conta2)
	if err != nil {
		panic(err)
	}

	fmt.Println(conta2)

	// nesse caso como tem notação no struct vai conseguir fazer o bind
	res = []byte(`{"n": 123, "s": 100}`)
	err = json.Unmarshal(res, &conta2)
	if err != nil {
		panic(err)
	}

	fmt.Println(conta2)
}
