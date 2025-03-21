package main

import (
	"github.com/gofiber/fiber/v2"
)

type ContactHandler struct{}

func NewContactHandler() CRUDHandler {
	return new(ContactHandler)
}

func (u *ContactHandler) Read() fiber.Handler {
	return func(ctx *fiber.Ctx) error { return ctx.SendString("Contact read") }
}

func (u *ContactHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error { return ctx.SendString("Contact create") }
}

func (u *ContactHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error { return ctx.SendString("Contact update") }
}

func (u *ContactHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error { return ctx.SendString("Contact delete") }
}
