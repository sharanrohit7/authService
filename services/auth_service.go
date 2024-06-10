package services

import (
	"context"

	"github.com/sharanrohit7/authService/config"
	"github.com/sharanrohit7/authService/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user *models.User) error {
	collection := config.DB.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	collection := config.DB.Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return &user, err
}
