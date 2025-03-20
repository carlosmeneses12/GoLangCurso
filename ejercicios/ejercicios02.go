package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func Multiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa un numero y te dare sus 10 primeras tablas de multiplicar, :  ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El valor ingresado genero el error, :  " + err.Error())
		}
		for n := 0; n <= 10; n++ {
			fmt.Println(numero, "*", n, "Es igual a = ", numero*n)
		}
	}

}
