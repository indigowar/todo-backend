package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

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
	UserID       string `bson:"_id"`
	UserName     string `bson:"name"`
	UserPassword string `bson:"password"`
	Token        struct {
		Value     string    `bson:"value"`
		ExpiredAt time.Time `bson:"expired_at"`
	} `bson:"token"`
}

func (u mongoUser) Id() uuid.UUID {
	id, _ := uuid.Parse(u.UserID)
	return id
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

func (u *userMongoRepo) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
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

func (u userMongoRepo) GetByName(ctx context.Context, name string) (uuid.UUID, error) {
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

func (u userMongoRepo) Create(ctx context.Context, user domain.User) error {
	if _, err := u.Get(ctx, user.Id()); err == nil {
		return errors.New("user already exists")
	}

	if _, err := u.GetByName(ctx, user.Name()); err == nil {
		return errors.New("user already exists")
	}

	mUser := mongoUser{
		UserID:       user.Id().String(),
		UserName:     user.Name(),
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

func (u userMongoRepo) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.D{{Key: "_id", Value: id.String()}}

	_, err := u.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("failed to delete a user")
	}
	return nil
}

func (u userMongoRepo) UpdatePassword(ctx context.Context, id uuid.UUID, value string) error {
	filter := bson.D{{Key: "_id", Value: id.String()}}
	updater := bson.D{{Key: "password", Value: value}}

	_, err := u.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return errors.New("failed to update the password")
	}
	return nil
}

func (u userMongoRepo) UpdateUserName(ctx context.Context, id uuid.UUID, value string) error {
	filter := bson.D{{Key: "_id", Value: id.String()}}
	updater := bson.D{{Key: "name", Value: value}}

	_, err := u.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return errors.New("failed to update the username")
	}
	return nil
}

func (u userMongoRepo) SetRefresh(ctx context.Context, id uuid.UUID, token string, expired_at time.Time) error {
	filter := bson.D{{Key: "_id", Value: id.String()}}
	updater := bson.D{{Key: "token.value", Value: token}, {Key: "token.expired_at", Value: expired_at}}

	_, err := u.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return errors.New("failed to update token")
	}
	return nil
}

func (u userMongoRepo) GetByRefresh(ctx context.Context, token string) (uuid.UUID, time.Time, error) {
	filter := bson.D{{Key: "token.value", Value: token}}

	var result mongoUser
	if err := u.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return uuid.UUID{}, time.Time{}, errors.New("user was not found by this token")
	}

	return result.Id(), result.TokenExpiredTime(), nil
}
