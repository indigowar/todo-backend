package repository

import (
	"time"

	"github.com/google/uuid"
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

func (u userMongoRepo) Get(uuid uuid.UUID) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userMongoRepo) GetByName(s string) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u userMongoRepo) Create(user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userMongoRepo) Delete(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u userMongoRepo) UpdatePassword(uuid uuid.UUID, s string) error {
	//TODO implement me
	panic("implement me")
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
