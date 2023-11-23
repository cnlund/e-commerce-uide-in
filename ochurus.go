package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func main() {
	//Zona de BDD
	var err error
	//Hacemos la conexion
	db, err = sql.Open("sqlserver", "server=CNLUNDPC;port=1433; database=Prueba")
	if err != nil {
		log.Fatal("Error al conectarse con la BDD: " + err.Error())
	}
	log.Printf("Se conecto!!!")
	defer db.Close()

	//Zona de HTML
	//Este es el handler con el que se ejecuta el HTML
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(rw, nil)
		}
	})
	//Aqui mandamos nuestra pagina web al puerto local 3000
	http.ListenAndServe(":3000", nil)
}
