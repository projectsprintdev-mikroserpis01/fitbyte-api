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

	authGroup := router.Group("/v1/auth")
	authGroup.Post("/", controller.handleAuth)
}

func (mc *authController) handleAuth(ctx *fiber.Ctx) error {
	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := mc.authService.Authenticate(ctx.Context(), req)
	if err != nil {
		return err
	}

	status := fiber.StatusOK
	if req.Action == "create" {
		status = fiber.StatusCreated
	}
	return ctx.Status(status).JSON(res)

}
