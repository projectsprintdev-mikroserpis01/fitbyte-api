package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
)

type userController struct {
	userService contracts.UserService
}

func InitNewController(router fiber.Router, userService contracts.UserService) {
	// init
}
