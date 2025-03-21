package arreglosslices

import (
	"fmt"
)

var tabla [10]int = [10]int{10, 2, 3} //definiendo slices (no tiene una dimension de la cantidad de elementos)
var matrix [20][30]int

func NuestroArreglo() {
	tabla[7] = 33
	tabla[2] = 2

	tabla2 := [10]string{"Hola", "pablo", "pedro", "zapatillas", "cero ", "Cuarenta"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matrix[5][7] = 15
	fmt.Println(matrix)
}
