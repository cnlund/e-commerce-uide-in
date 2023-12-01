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

// Struct tipografias para el html
type Tipografia struct {
	ttf   string
	otf   string
	woff  string
	woff2 string
}

// Zona de los Handlers------------------------------------------------------------------

// Handler para enviar el index.html
func handlerindex(c *fiber.Ctx) error {
	return c.Render("index", ".html")
}

// Handler para enviar la tipografia monday rain
func handlerMondayrain(c *fiber.Ctx) error {
	mr := Tipografia{
		ttf:   "public/Monday_rain/Monday_Rain.ttf",
		otf:   "public/Monday_rain/Monday-Rain.otf",
		woff:  "public/Monday_rain/Monday-Rain.woff",
		woff2: "public/Monday_rain/Monday-Rain.woff2",
	}
	error := c.BodyParser(mr)
	if error != nil {
		return error
	}
	return c.Status(fiber.StatusOK).JSON(mr)
}

// -----------------------------------------------------------------------------------------
func main() {
	// Zona HTML----------------------------------------------------------------------------
	app := fiber.New()
	app.Post("/", handlerindex, handlerMondayrain)
	//conexion con la api
	opt := option.WithCredentialsFile("servicekey.json")
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return
	}
	app.Listen(":3433")
}
