package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Zona de html
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hola Mundo")
	})
	//Aqui mandamos nuestra pagina web al puerto local 3000
	http.ListenAndServe(":3000", nil)
}
