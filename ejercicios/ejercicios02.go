package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*func Multiplicar() {
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
}*/

func Tablas() string {
	var numero int
	var err error
	var texto string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingresa un numero! : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}

	fmt.Println(texto)
	return texto
}
