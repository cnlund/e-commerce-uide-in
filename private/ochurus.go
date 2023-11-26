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

// Struct del formato de formulario de informacion sobre los postulantes
type Format_Post struct {
	Nombre  string
	Edad    int
	Ciudad  int
	Correo  string
	Carrera int
}

// Struct de las diversas habilidades que pueden haber en cada carrera
type Skills struct {
	Habilidad    string
	Conoce       bool
	Nivel_con    int
	Carrera_pert int
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

// handler de imagen prueba
func flechaHandler(rw http.ResponseWriter, r *http.Request) {
	fimg, error := template.ParseFiles("./public/Imagenes/Flecha.jpg")
	if error != nil {
		log.Fatal("Fallo la funcion de FLECHA ", error)
	}
	fimg.Execute(rw, nil)
}

// Handler de almacenamiento de ciudades
func ciudadesHandler(w http.ResponseWriter, r *http.Request) {
	var cnombre string
	rows, _ := db.Query("SELECT Ciudad_nombre FROM Ciudades")
	for rows.Next() {
		rows.Scan(&cnombre)
		fmt.Fprintln(w, cnombre)
	}
}

// Handler de almacenamiento de carreras
func carrerasHandler(rw http.ResponseWriter, r *http.Request) {
	var cnombre string
	rows, _ := db.Query("SELECT Carrera_nombre FROM Carreras")
	for rows.Next() {
		rows.Scan(&cnombre)
		fmt.Fprintln(rw, cnombre)
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
	mux.HandleFunc("/postular/lcarreras", carrerasHandler)
	mux.HandleFunc("/postular/imgflecha", flechaHandler)
	mux.HandleFunc("/contratar", contratantesHandler)
	//Aqui mandamos nuestra pagina web al puerto local 443
	http.ListenAndServe(":443", mux)
}
