package service

import (
	"github.com/gofiber/fiber/v2"
)

type Service struct{}

func New() *Service {
	return new(Service)
}

func (s *Service) RootHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello!!!")
}

func (s *Service) ContactRead(ctx *fiber.Ctx) error {
	return ctx.SendString("Contact read")
}

func (s *Service) ContactCreate(ctx *fiber.Ctx) error {
	return ctx.SendString("Contact create")
}

func (s *Service) ContactUpdate(ctx *fiber.Ctx) error {
	return ctx.SendString("Contact update")
}

func (s *Service) ContactDelete(ctx *fiber.Ctx) error {
	return ctx.SendString("Contact delete")
}

func (s *Service) GroupRead(ctx *fiber.Ctx) error {
	return ctx.SendString("Group read")
}

func (s *Service) GroupCreate(ctx *fiber.Ctx) error {
	return ctx.SendString("Group create")
}

func (s *Service) GroupUpdate(ctx *fiber.Ctx) error {
	return ctx.SendString("Group update")
}

func (s *Service) GroupDelete(ctx *fiber.Ctx) error {
	return ctx.SendString("Group delete")
}
