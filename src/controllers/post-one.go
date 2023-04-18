package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/src/data_access"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/dto"
	"github.com/gofiber/fiber/v2"
)

// Add one to db
func PostTweet(c *fiber.Ctx) error {
	var tweet dto.Tweet
	err := c.BodyParser(&tweet)

	newTweet, _ := data_access.TweetCreate(tweet)
	
	if err != nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newTweet)
}
