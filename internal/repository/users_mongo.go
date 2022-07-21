package repository

import (
	"github.com/google/uuid"
	"github.com/indigowar/todo-backend/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func NewUserMongoRepository(client *mongo.Client) (UserRepo, error) {
	return &userMongoRepo{
		client: client,
	}, nil
}

type userMongoRepo struct {
	client *mongo.Client
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
