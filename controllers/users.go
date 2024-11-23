package controllers

import (
	"github.com/gofiber/fiber/v2"
	"startup/models"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := models.FindAllUsers()
	if err != nil {
		// TODO: Handle Error properly by updating UI
		return err
	}

	return c.Render("users", fiber.Map{"Users": &users})
}
