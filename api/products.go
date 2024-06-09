package api

import (
	"request_limiter/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(c *fiber.Ctx) error {

	items := []models.Item{
		{ID: 1, Name: "Item 1", Price: 100},
		{ID: 2, Name: "Item 2", Price: 200},
		{ID: 3, Name: "Item 3", Price: 300},
	}

	return c.JSON(fiber.Map{
		"success": true,
		"items":   items,
	})
}
