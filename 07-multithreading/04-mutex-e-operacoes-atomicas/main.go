package main

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

var number uint64 = 0

// Temos formas de resolver o problema com concorrência:
// 1 - via mutex dando o lock e unlock
// 2 - No modo de desenvolvimento, da para usar o race, que vai dar alertas no console de onde está o problema,
// isso seria mais para ver o problema acontecer e não para resolver de fato. Para rodar o programa dessa segunda
// forma, basta executar o comando: go run --race main.go
// 3 - Usando o atomic, que é algo nativo do Go e gerência a concorrência de forma mais simples internamente.
func main() {
	// Forma com mutex:
	//m := sync.Mutex{}
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	m.Lock()
	//	number++
	//	m.Unlock()
	//	w.Write([]byte("Você teve " + strconv.FormatUint(number, 10) + " acessos!"))
	//})
	//http.ListenAndServe(":3000", nil)

	// Forma com Atomic:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		w.Write([]byte("Você teve " + strconv.FormatUint(number, 10) + " acessos!"))
	})
	http.ListenAndServe(":3000", nil)
}
