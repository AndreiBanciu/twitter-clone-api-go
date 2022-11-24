package data_access

import (
	"context"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/db"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/dto"
)

func TweetCreate(tweet dto.Tweet) (dto.Tweet, error) {
	twitterCloneCollection := db.DB.Collection("twitter_clone")
	
	_, err := twitterCloneCollection.InsertOne(context.Background(), tweet)

	if err != nil {
		return dto.Tweet{}, err
	}

	return tweet, nil
}
