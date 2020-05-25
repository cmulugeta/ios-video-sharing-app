package feed

import (
	"context"
	"errors"
	"github.com/cmulugeta/ios-video-sharing-app/windmill-backend/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserFeed(collection *mongo.Collection, ctx context.Context, userId string) ([]models.Post, error) {
	var user models.User
	collection.FindOne(ctx, bson.M{"userid":userId}).Decode(&user)
	if user.UserId == "" {
		return nil, errors.New("could't get user feed")
	}
	posts := user.Posts
	return posts, nil
}
