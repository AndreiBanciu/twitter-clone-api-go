package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/src/data_access"
	"github.com/gofiber/fiber/v2"
)

// Delete from db
func DeleteTweet(c *fiber.Ctx) error {
	id := c.Params("id")

	err := data_access.TweetDelete(id)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tweet deleted successfully",
	})
}
