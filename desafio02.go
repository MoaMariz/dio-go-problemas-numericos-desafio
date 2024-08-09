package main

import (
	"fmt"
	"strconv"
)

var lista []int
var listaPinPan []string

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			listaPinPan = append(listaPinPan, "PinPan")
		} else if i%3 == 0 {
			lista = append(lista, i)
			listaPinPan = append(listaPinPan, "Pin")
		} else if i%5 == 0 {
			listaPinPan = append(listaPinPan, "Pan")
		} else {
			var x = strconv.Itoa(i)
			listaPinPan = append(listaPinPan, x)
		}
	}
	fmt.Println("Desafio 1: Os números que são divisíveis por 3 no intervalo de 0 a 100 são:", lista)
	fmt.Println("Desafio 2:", listaPinPan)
}
