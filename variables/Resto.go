package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Numero int
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Carlos"
	Estado = true
	Numero = 10
	Sueldo = 100.5
	Fecha = time.Now()

	fmt.Println("Nombre = ", Nombre)
	fmt.Println("Estado = ", Estado)
	fmt.Println("Numero = ", Numero)
	fmt.Println("Sueldo = ", Sueldo)
	fmt.Println("Fecha = ", Fecha)

}

func ConviertoaTexto(numero int) (bool, string) {

	texto := strconv.Itoa(numero)
	return true, texto

}
