package main

import (
	"context"
	"html/template"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/rewrite"
	"google.golang.org/api/option"
)

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

// Handler para enviar el index.html
func handlerindex(c *fiber.Ctx) error {
	template, _ := template.ParseFiles("index.html")
	return template.Execute(c, nil)
}

// Handler para enviar hacia postular.html
func handlerpostular(c *fiber.Ctx) error {
	template, _ := template.ParseFiles("postulacion.html")
	return template.Execute(c, nil)
}

// Handler para ir a contratar
func handlercontratar(c *fiber.Ctx) error {
	template, _ := template.ParseFiles("contratar.html")
	return template.Execute(c, nil)
}

// -----------------------------------------------------------------------------------------
func main() {
	// Zona HTML----------------------------------------------------------------------------
	app := fiber.New()
	app.Use(rewrite.New(rewrite.Config{
		Rules: map[string]string{
			"/postulacion.html": "/postular",
		},
	}))
	app.Get("/", handlerindex)
	app.Get("/postular", handlerpostular)
	app.Get("/contratar", handlercontratar)
	//conexion con la api
	ctx := context.Background()
	sa := option.WithCredentialsFile("servicekey.json")
	api, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := api.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	//creamos la conexion del puerto
	app.Listen(":3000")
}
