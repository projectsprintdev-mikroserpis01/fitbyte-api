package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/internal/middlewares"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"
)

type activityController struct {
	activityService contracts.ActivityService
}

func InitActivityController(router fiber.Router, activityService contracts.ActivityService, middleware *middlewares.Middleware) {
	activityController := &activityController{
		activityService: activityService,
	}

	activityRouter := router.Group("/activity")
	activityRouter.Post("/", middleware.RequireAuth(), activityController.handleCreateActivity)
	activityRouter.Get("/", middleware.RequireAuth(), activityController.handleGetActivity)
	activityRouter.Patch("/:activityID", middleware.RequireAuth(), activityController.handleUpdateActivity)
	activityRouter.Patch("/", middleware.RequireAuth(), func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Activity ID not found",
		})
	})
	activityRouter.Delete("/:activityID", middleware.RequireAuth(), activityController.handleDeleteActivity)
	activityRouter.Delete("/", middleware.RequireAuth(), func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Activity ID not found",
		})
	})
}

func (ac *activityController) handleCreateActivity(ctx *fiber.Ctx) error {
	var req dto.CreateActivityRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userID := ctx.Locals("claims").(jwt.Claims).UserID
	req.UserID = userID

	res, err := ac.activityService.CreateActivity(ctx.Context(), req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (ac *activityController) handleGetActivity(ctx *fiber.Ctx) error {
	var req dto.GetActivityRequest
	if err := ctx.QueryParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := ac.activityService.GetActivity(ctx.Context(), req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (ac *activityController) handleUpdateActivity(ctx *fiber.Ctx) error {
	var req dto.UpdateActivityRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	activityID := ctx.Params("activityID")
	id, err := uuid.Parse(activityID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid activity ID",
		})
	}
	req.ActivityID = id

	userID := ctx.Locals("claims").(jwt.Claims).UserID
	req.UserID = userID

	res, err := ac.activityService.UpdateActivity(ctx.Context(), req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (ac *activityController) handleDeleteActivity(ctx *fiber.Ctx) error {
	var req dto.DeleteActivityRequest

	activityID := ctx.Params("activityID")
	id, err := uuid.Parse(activityID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid activity ID",
		})
	}
	req.ActivityID = id

	err = ac.activityService.DeleteActivity(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Activity deleted successfully",
	})
}
