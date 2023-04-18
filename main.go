package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/controllers"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/db"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	appApi := app.Group("/api")

	// comes from db
	appApi.Get("/tweets", controllers.GetTweets)
	appApi.Get("/tweets/:id", controllers.GetTweetById)
	appApi.Post("/tweets", controllers.PostTweet)
	appApi.Delete("/tweets/:id", controllers.DeleteTweet)
	appApi.Put("/tweets/:id", controllers.EditTweet)
		
	db.ConnectToDb()
	app.Listen("localhost:5000")
}
