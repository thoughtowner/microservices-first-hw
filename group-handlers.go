package main

import (
	"github.com/gofiber/fiber/v2"
)

type GroupHandler struct{}

func NewGroupHandler() CRUDHandler {
	return new(GroupHandler)
}

func (u *GroupHandler) Read() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Group read")
	}
}

func (u *GroupHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Group create")
	}
}

func (u *GroupHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Group update")
	}
}

func (u *GroupHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Group delete")
	}
}
