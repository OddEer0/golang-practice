package mongodb

import (
	"context"
	"fmt"

	resourceModel "github.com/OddEer0/golang-practice/resources/resource_model"
	resourceRepository "github.com/OddEer0/golang-practice/resources/resource_repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

// Create implements resourceRepository.User.
func (u *userRepository) Create(ctx context.Context, user *resourceModel.User) (*resourceModel.User, error) {
	res, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println("inserted id ", res.InsertedID)
	return user, nil
}

func (u *userRepository) UpdateUserLogin(ctx context.Context, id string, login string) (*resourceModel.User, error) {
	where := bson.D{{Key: "id", Value: id}}
	update := bson.D{{
		Key: "login", Value: login,
	}}
	res, err := u.collection.UpdateOne(ctx, where, update)
	if err != nil {
		return nil, err
	}
	fmt.Println("updated count ", res.ModifiedCount)
	return &resourceModel.User{}, nil
}

// CreateCopy implements resourceRepository.User.
func (u *userRepository) CreateCopy(context.Context, resourceModel.User) (resourceModel.User, error) {
	panic("unimplemented")
}

// DeleteById implements resourceRepository.User.
func (u *userRepository) DeleteById(context.Context, string) error {
	panic("unimplemented")
}

// GetById implements resourceRepository.User.
func (u *userRepository) GetById(ctx context.Context, id string) (*resourceModel.User, error) {
	result := new(resourceModel.User)
	where := bson.D{{Key: "id", Value: id}}

	err := u.collection.FindOne(ctx, where).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetByIdCopy implements resourceRepository.User.
func (u *userRepository) GetByIdCopy(context.Context, string) (resourceModel.User, error) {
	panic("unimplemented")
}

func NewUserRepository(collection *mongo.Collection) resourceRepository.User {
	return &userRepository{}
}
