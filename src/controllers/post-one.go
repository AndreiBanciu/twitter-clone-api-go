package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/src/data"
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

// Add one to hardcoded data
func PostTodo(c *fiber.Ctx) error {
	var newTodo dto.Todo

	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Couldn't create todo",
		})
	}

	data.Todos = append(data.Todos, newTodo)
	return c.Status(fiber.StatusCreated).JSON(newTodo)
}
