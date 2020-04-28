package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/cmulugeta/ios-video-sharing-app/windmill-backend/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/oauth2/v2"

	//"github.com/google/uuid"
	//"net/http"
)


func CheckHashedPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckUserExists(collection *mongo.Collection, ctx context.Context, username string) (bool, map[string]interface{}) {

	err := collection.FindOne(ctx, bson.M{"username":username})
	if err.Err() == nil {
		return true, map[string]interface{}{
			"result":"username has been taken",
			"error":err,
		}
	}

	return false, map[string]interface{}{
		"message":"credentials available!",
	}
}

func SignUpUser(collection *mongo.Collection, ctx context.Context, data *models.User) (bool, *mongo.InsertOneResult) {
	user := models.User{
		UserId:    uuid.New(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Email:     data.Email,
		Relations: models.Relationships{},
	}
	res, _ := collection.InsertOne(ctx, user)
	return true, res
}


func GetUser(collection *mongo.Collection, ctx context.Context, token models.GoogleToken, info *oauth2.Tokeninfo) (models.User, string){
	var user models.User
	collection.FindOne(ctx, bson.M{"email":info.Email}).Decode(&user)
	if len(user.Username) == 0 {
		return models.User{
			UserId:    uuid.New(),
			UserToken: token,
			FirstName: "",
			LastName:  "",
			Username:  "",
			Email:     info.Email,
			Verified:  false,
			Relations: models.Relationships{},
		}, "redirecting to username creation..."
	}
	return user, ""
}






