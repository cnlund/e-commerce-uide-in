package main

import (
	"context"
	"log"
	"text/template"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
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
	mensaje := "funciona"
	template, err := template.ParseFiles("public/index.html")
	if err != nil {
		return err
	}
	return template.Execute(c, mensaje)
}

// Handler para enviar hacia postular.html
func handlerpostular(c *fiber.Ctx) error {
	mensaje := "el backend te ha movido a postular"
	template, err := template.ParseFiles("public/postular.html")
	if err != nil {
		return err
	}
	return template.Execute(c, mensaje)
}

// Handler para ir a contratar
func handlercontratar(c *fiber.Ctx) error {
	template, _ := template.ParseFiles("public/contratar.html")
	return template.Execute(c, nil)
}

// -----------------------------------------------------------------------------------------
func main() {
	// Zona HTML----------------------------------------------------------------------------
	app := fiber.New()
	app.Post("/", handlerindex)
	app.Post("/postular", handlerpostular)
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
	app.Listen(":403")
}
