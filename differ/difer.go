package differ

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("este es un primer mensaje")
	defer fmt.Println("este es un mensaje defer")

	fmt.Println("este es un mensaje final")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
