package main

// Thread 1
func main() {
	forever := make(chan bool) // canal está vazio

	// Thread 2
	go func() {
		for i := 0; i < 10; i++ {
			println(i) // canal está cheio
		}
		forever <- true // canal está cheio
	}()

	<-forever // esperando o canal ficar cheio para esvaziar, em tese segura o processo

}
