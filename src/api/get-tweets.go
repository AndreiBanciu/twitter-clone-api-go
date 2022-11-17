package api

import (
	"net/http"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/data"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/dto"
	"github.com/gin-gonic/gin"
)

func GetTweets(c *gin.Context) {
	response := dto.Response{
		Data: data.Tweets,
		Total: len(data.Tweets),
	}
	
    c.IndentedJSON(http.StatusOK, response)
}
