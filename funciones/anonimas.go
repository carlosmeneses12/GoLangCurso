package funciones

import (
	"fmt"
)

func Calculos() {

	var numero3 int = 32
	var numero4 int = 43

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}
	fmt.Println(suma(19, 23))

	suma = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 * numero4
	}
	fmt.Println(suma(19, 23))
}
