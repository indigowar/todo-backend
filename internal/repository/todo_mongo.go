package repository

import (
	"github.com/google/uuid"
	"github.com/indigowar/todo-backend/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTodoRepo(client *mongo.Client) (TodoRepo, error) {
	return &todoMongoRepo{
		client: client,
	}, nil
}

type todoMongoRepo struct {
	client *mongo.Client
}

func (t todoMongoRepo) GetListByID(uuid uuid.UUID) (domain.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) CreateList(list domain.List) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) GetListsByOwner(uuid uuid.UUID) ([]uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) DeleteList(uuid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) GetElement(uuid uuid.UUID, uuid2 uuid.UUID) (domain.Element, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) AddElement(uuid uuid.UUID, element domain.Element) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) DeleteElement(uuid uuid.UUID, uuid2 uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) ChangeStatus(uuid uuid.UUID, uuid2 uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) RenameElement(uuid uuid.UUID, uuid2 uuid.UUID, s string) error {
	//TODO implement me
	panic("implement me")
}
