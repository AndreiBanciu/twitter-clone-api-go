package api

import (
	"net/http"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/data"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/dto"
	"github.com/gin-gonic/gin"
)

func PostTweet(c *gin.Context){
	var newTweet dto.Tweet

	if err := c.BindJSON(&newTweet); err != nil {
		return
	}

	data.Tweets = append(data.Tweets, newTweet)
	c.IndentedJSON(http.StatusCreated, newTweet)
}
