package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

// variable de la base de datos
var db *sql.DB

// Zona de los Structs-------------------------------------------------------------------
type Filtro struct {
	Ciudad string
}
type Format_Post struct {
	Nombre  string
	Edad    int
	Ciudad  string
	Correo  string
	Carrera string
}

// Zona de los Handlers------------------------------------------------------------------
// Handler de la pagina de INICIO
func indexHandler(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./public/index.html")
	if err != nil {
		log.Fatal("Fallo la funcion de INDEX ", err)
	} else {
		template.Execute(rw, nil)
	}
}

// Handler de la pagina de POSTULANTES
func postulanteHandler(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./public/postulacion.html")
	template.Execute(rw, nil)
}

// Handler de la pagina de CONTRATANTES
func contratantesHandler(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./public/contratar.html")
	template.Execute(rw, nil)
}

// -----------------------------------------------------------------------------------------
func main() {
	//Zona de BDD---------------------------------------------------------------------------
	var err error
	//Hacemos la conexion
	db, err = sql.Open("sqlserver", "server=CNLUNDPC;port=1433; database=Prueba")
	if err != nil {
		log.Fatal("Error al conectarse con la BDD: " + err.Error())
	}
	log.Printf("Se conecto!!!")
	defer db.Close()
	// Zona HTML----------------------------------------------------------------------------
	//Creamos los handlerfuncs
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/postular", postulanteHandler)
	http.HandleFunc("/contratar", contratantesHandler)
	//Aqui mandamos nuestra pagina web al puerto local 3000
	http.ListenAndServe(":443", nil)
}
