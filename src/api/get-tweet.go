package api

import (
	"net/http"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/data"
	"github.com/gin-gonic/gin"
)

func GetTweetById(c *gin.Context){
	id := c.Param("id")

	for _, tweet := range data.Tweets {
		if tweet.ID == id {
			c.IndentedJSON(http.StatusOK, tweet)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tweet not found"})
}
