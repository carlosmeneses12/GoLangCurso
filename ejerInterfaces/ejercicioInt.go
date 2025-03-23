package ejerinterfaces

import (
	"fmt"

	"github.com/GoLangEcomerce/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un %s, y estoy respirando  \n", hu.Sexo())
}
