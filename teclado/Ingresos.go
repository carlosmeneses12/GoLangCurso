package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoDeNumeros() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer numero: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado vale pico, " + err.Error())
		}
	}

	fmt.Println("Ingrese el segundo numero: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("el dato ingresado vale pico, " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda  : ")
	if scanner.Scan() {
		leyenda = scanner.Text()

	}
	fmt.Println(leyenda, numero1*numero2)

}
