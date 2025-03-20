package ejercicios

import (
	"strconv"
)

func ConvertiraEntero(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un Error" + err.Error()
	}
	if num >= 100 {
		return num, "el numero es mayor a 100"
	} else {
		return num, "el numero es menor a 100"
	}
}
