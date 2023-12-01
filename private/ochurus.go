package main

import (
	"context"

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
	return c.Render("index", ".html")
}

// Handler para enviar hacia postular.html
func handlerpostular(c *fiber.Ctx) error {
	return c.Render("postulacion", ".html")
}

// -----------------------------------------------------------------------------------------
func main() {
	// Zona HTML----------------------------------------------------------------------------
	app := fiber.New()
	app.Post("/", handlerindex)
	app.Post("/postular", handlerpostular)
	//conexion con la api
	opt := option.WithCredentialsFile("servicekey.json")
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return
	}
	app.Listen(":3433")
}
