package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/GoLangEcomerce/ejercicios"
)

var Ruta string = "./files/txt/tablas.txt"

func GrabaTabla() {
	var texto string = ejercicios.Tablas()
	archivo, err := os.Create(Ruta)
	if err != nil {
		fmt.Println("error al crear el archivo , ", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Tablas()
	if !Append(Ruta, texto) {
		fmt.Println("error al concatenar contenido")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(Ruta, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Se generó un error we xD ---->  ", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Se generó un error en wrtieString ---->  ", err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(Ruta)
	if err != nil {
		fmt.Println("ocurrio un error cuate--> ", err.Error())
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("* uwu * " + registro)
	}
	archivo.Close()

}
