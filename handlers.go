package main

import "github.com/gofiber/fiber/v2"

func RootHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("hello")
}
