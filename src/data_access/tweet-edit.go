package data_access

import (
	"context"

	"github.com/AndreiBanciu/twitter-clone-api-go/src/db"
	"github.com/AndreiBanciu/twitter-clone-api-go/src/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TweetEdit(tweetId string, tweet dto.Tweet) error {
	twitterCloneCollection := db.DB.Collection("twitter_clone")

	objectId, err := primitive.ObjectIDFromHex(tweetId)
	if err != nil {
		return err
	}

	update := bson.M{
        "$set": tweet,
    }
	_, err = twitterCloneCollection.UpdateOne(context.Background(), bson.M{"_id": objectId}, update)
	if err != nil {
		return err
	}


	return nil
}
