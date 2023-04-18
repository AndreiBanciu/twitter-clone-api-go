package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/src/data_access"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Edit one from db
func EditTweet(c *fiber.Ctx) error {
	var tweet dto.Tweet
	err := c.BodyParser(&tweet)
	id := c.Params("id")

	data_access.TweetEdit(id, tweet)
	
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Tweet updated successfully",
	})
}
