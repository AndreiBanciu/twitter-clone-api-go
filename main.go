package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/controllers"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/db"
)

func main() {
	app := fiber.New()

	appApi := app.Group("/api")

	// comes from db
	appApi.Get("/tweets", controllers.GetTweets)
	appApi.Get("/tweets/:id", controllers.GetTweetById)
	appApi.Post("/tweets", controllers.PostTweet)
	appApi.Delete("/tweets/:id", controllers.DeleteTweet)
	appApi.Put("/tweets/:id", controllers.EditTweet)
	
	// comes from hardcoded data
	appApi.Post("/todo", controllers.PostTodo)
	appApi.Get("/todo/:id", controllers.GetTodo)
	appApi.Get("/todo", controllers.GetTodos)
	appApi.Delete("/todo/:id", controllers.DeleteTodo)
	appApi.Put("/todo/:id", controllers.EditTodo)
		
	db.ConnectToDb()
	app.Listen("localhost:5000")
}
