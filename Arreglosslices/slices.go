package arreglosslices

import (
	"fmt"
)

var TaBlaaS []int = []int{2, 5, 6}
var Arreglo [10]int = [10]int{6, 4, 3, 2, 1}

func MuestroSlice() {
	fmt.Println(TaBlaaS)

	porcion := Arreglo[3:] //slice creado desde un vector desde posicion 3

	porcion2 := Arreglo[:5] //slice creado desde el inicio hasta el 5

	porcion3 := Arreglo[6:7] // desde una posicion determinada hasta una posicion determinada

	fmt.Println(porcion, porcion2, porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, Capacidad %d", len(elementos), cap(elementos))
}
