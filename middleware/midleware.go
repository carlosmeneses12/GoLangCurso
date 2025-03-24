package middleware

import (
	"fmt"
)

func Sumar(a, b int) int {
	return a + b
}

func Restar(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Ejecutando mi middleware")

	result := operacionesMid(Sumar)(2, 5)
	fmt.Println(result)

	result = operacionesMid(Restar)(2, 5)
	fmt.Println(result)

	result = operacionesMid(Multiplicar)(2, 5)
	fmt.Println(result)
}

func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("inicio de operacion")
		return f(a, b)
	}
}
