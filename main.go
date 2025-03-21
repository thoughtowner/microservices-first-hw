package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type CRUDHandler interface {
	Create() fiber.Handler
	Read() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

func main() {
	app := fiber.New()

	initCRUDHandler(app, "/api/v1/contact", NewContactHandler())
	initCRUDHandler(app, "/api/v1/group", NewGroupHandler())

	if err := app.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}

func initCRUDHandler(app *fiber.App, path string, h CRUDHandler) {
	app.Get(path, h.Read())
	app.Post(path, h.Create())
	app.Put(path, h.Update())
	app.Delete(path, h.Delete())
}
