package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/src/data"
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

// Get one from hardcoded data
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, item := range data.Todos {
		if item.ID == id {
			return c.JSON(item)
			
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Todo not found",
	})
}
