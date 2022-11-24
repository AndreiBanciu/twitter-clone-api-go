package controllers

import (
	"github.com/AndreiBanciu/twitter-clone-api-go/data"
	"github.com/AndreiBanciu/twitter-clone-api-go/data_access"
	"github.com/AndreiBanciu/twitter-clone-api-go/dto"
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

// Delete from hardcoded data
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	newSlice := []dto.Todo{}

	for _, todo := range data.Todos {
		if todo.ID != id {
			newSlice = append(newSlice, todo)
		}
	}

	if len(newSlice) == len(data.Todos) {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found!")
	}

	data.Todos = newSlice
	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}
