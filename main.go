package main

import (
	"fmt"

	"github.com/GoLangEcomerce/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(5943)
	fmt.Println(estado)
	fmt.Println(texto)
}
