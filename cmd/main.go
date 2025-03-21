package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/thoughtowner/microservices-first-hw/internal/service"
)

func main() {
	api := fiber.New()

	svc := service.New()

	api.Get("/", svc.RootHandler)

	api.Get("/api/contact", svc.ContactRead)
	api.Post("/api/contact", svc.ContactCreate)
	api.Put("/api/contact", svc.ContactUpdate)
	api.Delete("/api/contact", svc.ContactDelete)

	api.Get("/api/group", svc.GroupRead)
	api.Post("/api/group", svc.GroupCreate)
	api.Put("/api/group", svc.GroupUpdate)
	api.Delete("/api/group", svc.GroupDelete)

	if err := api.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}
