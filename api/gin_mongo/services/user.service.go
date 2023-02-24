package services

import (
	"context"
	"errors"

	"github.com/jaredmyers/gogym/api/gin_mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userCollection *mongo.Collection
	ctx            context.Context // i've learned storing ctx in struct is no good.
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserService {
	return &UserService{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

// db.collection.find({name:"name"})
func (u *UserService) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err

}
func (u *UserService) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("Documents not found")
	}

	return users, nil
}
func (u *UserService) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "name", Value: user.Name},
		bson.E{Key: "age", Value: user.Age},
		bson.E{Key: "address", Value: user.Address},
	}}}

	result, _ := u.userCollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update")
	}
	return nil
}
func (u *UserService) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document found for delete")
	}
	return nil
}
