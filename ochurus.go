package main

import (
	"database/sql"
	"fmt"
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
	rows, err := db.Query("INSERT INTO Tabla_prueba (id, dato_prueba) VALUES (4, 'isra');")
	if err != nil {
		log.Fatal("Error al insertar: " + err.Error())
	}
	defer rows.Close()

	//Zona de HTML
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hola Mundo")
	})
	//Aqui mandamos nuestra pagina web al puerto local 3000
	//http.ListenAndServe(":3000", nil)
}
