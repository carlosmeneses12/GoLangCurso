package webserver

import (
	"fmt"
	"net/http"
)

func IniciarWeb() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
	fmt.Fprintf(w, "Hola Mundo")
}
