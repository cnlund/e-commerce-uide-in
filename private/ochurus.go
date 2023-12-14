package main

import (
	"github.com/gofiber/fiber/v2",
	"github.com/pocketbase/pocketbase"
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

// Handler con la lista de las ciudades
func ciudadesHandler(c *fiber.Ctx) error {
	app := pocketbase.New()
	res, _ := app.Dao().DB().NewQuery("select * from ciudades").Execute()
	return res
}

// -----------------------------------------------------------------------------------------
func main() {
	// Zona HTML----------------------------------------------------------------------------
	// Aqui tambien le asignamos con el nombre que queremos que aparezca en la pagina
	web := fiber.New()
	web.Get("/", indexHandler)
	web.Get("/comprobacion", comprohandler)
	web.Post("/postular", postularHandler)
	web.Post("/contratar", contratarHandler)
	//creamos la conexion del puerto
	web.Listen(":403")
}
