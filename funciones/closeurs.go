package funciones

import (
	"fmt"
)

func tablita(valor int) func() int {
	number := valor
	secuencia := 0
	return func() int {
		secuencia++
		return number * secuencia
	}

}

func LlamarClosure() {
	tabladel := 2
	Mitabla := tablita(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
