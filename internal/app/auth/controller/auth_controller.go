package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
)

type authController struct {
	authService contracts.AuthService
}

func InitAuthController(router fiber.Router, authService contracts.AuthService) {
	controller := authController{
		authService: authService,
	}

	v1 := router.Group("/v1")
	v1.Post("/login", controller.handleLogin)
	v1.Post("/register", controller.handleRegister)
}

func (mc *authController) handleLogin(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := mc.authService.Login(ctx.Context(), req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res)

}

func (mc *authController) handleRegister(ctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := mc.authService.Register(ctx.Context(), req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)

}
