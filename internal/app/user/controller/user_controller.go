package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/app/user/service"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/middlewares"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"
)

type managerController struct {
	managerService service.UserService
}

func InitUserController(router fiber.Router, managerService service.UserService) {
	controller := managerController{
		managerService: managerService,
	}

	jwt := jwt.Jwt

	middleware := middlewares.NewMiddleware(jwt)

	managerRoute := router.Group("/v1/user")
	managerRoute.Get("/", middleware.RequireAuth(), controller.GetUserById)
	managerRoute.Patch("/", middleware.RequireAuth(), controller.UpdateUserById)
}

func (mc *managerController) GetUserById(ctx *fiber.Ctx) error {
	managerID := ctx.Locals("claims").(jwt.Claims).UserID

	res, err := mc.managerService.GetUserById(ctx.Context(), managerID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (mc *managerController) UpdateUserById(ctx *fiber.Ctx) error {
	var requestBody dto.UpdateUserRequest
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	managerID := ctx.Locals("claims").(jwt.Claims).UserID

	_, err := mc.managerService.UpdateUserById(ctx.Context(), managerID, requestBody)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(requestBody)
}
