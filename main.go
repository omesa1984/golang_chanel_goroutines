/*
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/

package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	//for para lanzar as goroutines
	Envia(c)

	Mostra(c)
}

func Envia(c chan int) {
	for i := 0; i < 10; i++ {
		//func para gerar e enviar os números para o canal
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
}

func Mostra(c chan int) {
	for k := 0; k < 100; k++ {
		fmt.Println(k, "\t", <-c)
	}
}
