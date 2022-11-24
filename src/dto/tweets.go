package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tweet struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	User  string `json:"user"`
	Value string `json:"value"`
}

type GetAllTweetsResponse struct {
	Data  []Tweet `json:"data"`
	Total int     `json:"total"`
}
