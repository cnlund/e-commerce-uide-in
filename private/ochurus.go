package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
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
func indexHandler(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
}

// Handler de comprobacion
func comprohandler(c *fiber.Ctx) error {
	return c.SendString("El backend esta encendido")
}

// Handler que lleva hacia postular
func postularHandler(c *fiber.Ctx) error {
	return c.SendFile("public/postular.html")
}

// Handler que lleva hacia contratar
func contratarHandler(c *fiber.Ctx) error {
	return c.SendFile("public/contratar.html")
}

// Creamos una lista de las ciudades
func lciudades(ctx context.Context, client *firestore.Client) error {
	return client.Collection("cities").Documents(ctx)
}

// -----------------------------------------------------------------------------------------
func main() {
	// Zona de la BDD ----------------------------------------------------------------------
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "ochurus-a8fb4"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	// Zona HTML----------------------------------------------------------------------------
	web := fiber.New()
	web.Get("/", indexHandler)
	web.Get("/comprobacion", comprohandler)
	web.Post("/postular", postularHandler)
	web.Post("/contratar", contratarHandler)
	web.Static("/ciudades")
	//creamos la conexion del puerto
	web.Listen(":403")
}
