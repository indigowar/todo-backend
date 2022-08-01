package repository

import (
	"github.com/google/uuid"
	"github.com/indigowar/todo-backend/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTodoRepo(database *mongo.Database) (TodoRepo, error) {
	return &todoMongoRepo{
		lists: database.Collection("lists"),
		elements: database.Collection("elements"),
	}, nil
}

type mongoElement struct {
	ID string `bson:"_id"`
	ElementValue string `bson:"value"`
	Status  bool `bson:"done"`
}
func (e *mongoElement) Id() uuid.UUID {
	id, _ := uuid.Parse(e.ID)
	return id
}
func (e *mongoElement) Value() string {
	return e.ElementValue
}
func (e *mongoElement) Done() bool {
	return e.Status
}

type mongoList struct {
	ID string `bson:"_id"`
	ListName     string `bson:"name"`
	ListOwner    string `bson:"owner"`
	ElementsID []string `bson:"elements"`
}
func (l *mongoList) Id() uuid.UUID {
	id, _ := uuid.Parse(l.ID)
	return id
}
func (l *mongoList) Name() string {
	return l.ListName
}
func (l *mongoList) Owner() uuid.UUID {
	id, _ := uuid.Parse(l.ListOwner)
	return id
}
func (l *mongoList) Elements() []uuid.UUID {
	result := make([]uuid.UUID, len(l.ElementsID))
	for i, v := range l.ElementsID {
		id, _ := uuid.Parse(v)
		result[i] = id
	}
	return result
}

type todoMongoRepo struct {
	lists *mongo.Collection
	elements *mongo.Collection
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
