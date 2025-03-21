package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D .F . M"
	paises["Argentina"] = " Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"]) //al llamar al elemento por su clave solo muestra el valor "Buenos aires"

	campeonato := map[string]int{
		"Barcelona ": 39,
		"U de chile": 37,
		"colo-colo":  -3,
		"Catolica":   -6,
	}

	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}*/

	delete(campeonato, "Barcelona ")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["U de chile"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t", puntaje, existe)
}
