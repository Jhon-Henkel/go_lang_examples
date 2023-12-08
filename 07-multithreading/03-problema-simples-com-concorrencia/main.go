package main

import (
	"net/http"
	"strconv"
)

var number uint64 = 0

// - O problema que pode acontecer é o processo ser executado em mais de uma thread e a variável
// number ser acessada ao mesmo tempo por mais de uma thread, o que pode causar inconsistência no valor da variável.
// - O apache ab é uma ferramenta que pode ser usada para testar a aplicação.
// - Para testar a aplicação com 1000 requisições, 100 de cada vez, execute o comando: ab -n 1000 -c 100 http://localhost:3000/
// - A solução para o problema de concorrência será apresentada no próximo exemplo.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte("Você teve " + strconv.FormatUint(number, 10) + " acessos!"))
	})
	http.ListenAndServe(":3000", nil)
}
