package repository

import (
	"time"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/indigowar/todo-backend/internal/domain"
)

func NewUserMongoRepository(database *mongo.Database) (UserRepo, error) {
	return &userMongoRepo{
		collection: database.Collection("users"),
	}, nil
}

type userMongoRepo struct {
	collection *mongo.Collection
}

// implements User interface
type mongoUser struct {
	UserID uuid.UUID `bson:"_id"`
	UserName string `bson:"name"`
	UserPassword string `bson:"password"`
	Token struct {
		Value string `bson:"value"`
		ExpiredAt time.Time `bson:"expired_at"`
	} `bson:"token"`
}

func (u mongoUser) Id() uuid.UUID {
	return u.UserID
}
func (u mongoUser) Name() string {
	return u.UserName
}
func (u mongoUser) Password() string {
	return u.UserPassword
}
func (u mongoUser) TokenValue() string {
	return u.Token.Value
}
func (u mongoUser) TokenExpiredTime() time.Time {
	return u.Token.ExpiredAt
}

func (u *userMongoRepo) Get(id uuid.UUID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: id.String()}}
	var result mongoUser
	if err := u.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.NewUser(uuid.UUID{}, "", ""), errors.New("user was not found")
		}
		return domain.NewUser(uuid.UUID{}, "", ""), errors.New("internal error")
	}
	return result, nil
}

func (u userMongoRepo) GetByName(name string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	filter := bson.D{{Key: "name", Value: name}}

	var result mongoUser
	if err := u.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return uuid.UUID{}, errors.New("user was not found")
		}
		return uuid.UUID{}, errors.New("internal error")
	}
	return result.Id(), nil
}

func (u userMongoRepo) Create(user domain.User) error {
	if _, err := u.Get(user.Id()); err == nil {
		return errors.New("user already exists")
	}

	if _, err := u.GetByName(user.Name()); err == nil {
		return errors.New("user already exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	mUser := mongoUser{
		UserID: user.Id(),
		UserName: user.Name(),
		UserPassword: user.Password(),
	}
	mUser.Token.Value = user.TokenValue()
	mUser.Token.ExpiredAt = user.TokenExpiredTime()

	_, err := u.collection.InsertOne(ctx, mUser)
	if err != nil {
		return errors.New("internal error, can't insert user")
	}
	return nil
}

func (u userMongoRepo) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id.String()}}

	_, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("failed to delete a user")
	}
	return nil
}

func (u userMongoRepo) UpdatePassword(id uuid.UUID, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id.String()}}
	updater := bson.D{{Key: "password", Value: value}}

	_, err := u.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return errors.New("failed to update the password")
	}
	return nil
}

func (u userMongoRepo) UpdateUserName(uuid uuid.UUID, s string) error {
	//TODO implement me
	panic("implement me")
}

func (u userMongoRepo) SetRefresh(uuid uuid.UUID, s string, time time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (u userMongoRepo) GetByRefresh(s string) (uuid.UUID, time.Time, error) {
	//TODO implement me
	panic("implement me")
}
