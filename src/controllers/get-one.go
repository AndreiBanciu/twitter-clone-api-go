package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/src/data_access"
	"github.com/gofiber/fiber/v2"
)

// Get one from db
func GetTweetById(c *fiber.Ctx) error {
	id := c.Params("id")

	tweet, err := data_access.TweetGet(id)

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(tweet)
}
