package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/denisenkom/go-mssqldb"
)

// variable de la base de datos
var db *sql.DB

// Zona de los Structs-------------------------------------------------------------------
// Struct para el buscador
type Ciudad struct {
	Nombre string
}

// Struct del formato de formulario de informacion sobre los postulantes
type Format_Post struct {
	Nombre   string
	Edad     int
	Ciudad   string
	Correo   string
	Egresado Carrera
}

// Struct de las diversas habilidades que pueden haber en cada carrera
type Skills struct {
	Habilidad    string
	Conoce       bool
	Nivel_con    int
	Carrera_pert Carrera
}

// Struct de las carreras con sus datos relevantes
type Carrera struct {
	Nombrec string
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

// Handler de almacenamiento de ciudades
func ciudadesHandler(rw http.ResponseWriter, r *http.Request) {
	var cid int
	rows2, _ := db.Query("SELECT Ciudad_id FROM Ciudades")
	for rows2.Next() {
		rows2.Scan(&cid)
	}
	var cnombre string
	rows, _ := db.Query("SELECT Ciudad_nombre FROM Ciudades")
	for rows.Next() {
		rows.Scan(&cnombre)
		cid := Ciudad{cnombre}
		fmt.Fprintln(rw, cid.Nombre)
	}
}

// Handler de la pagina de CONTRATANTES
func contratantesHandler(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("./public/contratar.html")
	template.Execute(rw, nil)
}

// -----------------------------------------------------------------------------------------
func main() {
	//Creamos el server mux
	mux := http.NewServeMux()
	//Zona de BDD---------------------------------------------------------------------------
	var err error
	//Hacemos la conexion
	db, err = sql.Open("sqlserver", "server=CNLUNDPC;port=1433; database=Ochurus_DB")
	if err != nil {
		log.Fatal("Error al conectarse con la BDD: " + err.Error())
	}
	log.Printf("Se conecto!!!")
	defer db.Close()
	// Zona HTML----------------------------------------------------------------------------
	//Creamos los handlerfuncs
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/postular", postulanteHandler)
	mux.HandleFunc("/postular/lciudades", ciudadesHandler)
	mux.HandleFunc("/contratar", contratantesHandler)
	//Aqui mandamos nuestra pagina web al puerto local 443
	http.ListenAndServe(":443", mux)
}
