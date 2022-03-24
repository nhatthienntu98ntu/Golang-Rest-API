package services

import (
	"context"
	"errors"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var ctx = context.TODO()

// Get All User
func GetUsers() ([]*models.UserResponse, error) {
	var users []*models.UserResponse
	// Get all user
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	// Reading data
	for result.Next(ctx) {
		var user models.UserResponse
		if err = result.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = result.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("document not found")
	}
	return users, nil
}

//Get User
func GetUser(id *primitive.ObjectID) (*models.UserResponse, error) {
	var user *models.UserResponse
	fmt.Println("get user by id: ", id)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	fmt.Println("User: ", user)
	return user, err
}

// Create New User
func AddUser(user *models.User) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, user)
	return result, err
}

// Update User
func UpdateUser(id primitive.ObjectID, user *models.User) (*models.UserResponse, error) {
	// Update user
	update := bson.M{"username": user.Username, "email": user.Email, "phone": user.Phone, "dateOfBirth": user.DateOfBirth}
	result, _ := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if result.MatchedCount != 1 {
		return nil, errors.New("no matched document found for update")
	}

	// Find User
	var updateUser *models.UserResponse
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&updateUser)
	if err != nil {
		return nil, err
	}

	return updateUser, err
}

// Delete User
func DeleteUser(id primitive.ObjectID) error {
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return err
}
