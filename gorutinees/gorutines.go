package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MinombreLento(name string, canal1 chan bool) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canal1 <- true
}
