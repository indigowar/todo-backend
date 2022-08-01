package repository

import (
	"context"
	"time"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/indigowar/todo-backend/internal/domain"
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
func (e mongoElement) Id() uuid.UUID {
	id, _ := uuid.Parse(e.ID)
	return id
}
func (e mongoElement) Value() string {
	return e.ElementValue
}
func (e mongoElement) Done() bool {
	return e.Status
}

type mongoList struct {
	ID string `bson:"_id"`
	ListName     string `bson:"name"`
	ListOwner    string `bson:"owner"`
	ElementsID []string `bson:"elements"`
}
func (l mongoList) Id() uuid.UUID {
	id, _ := uuid.Parse(l.ID)
	return id
}
func (l mongoList) Name() string {
	return l.ListName
}
func (l mongoList) Owner() uuid.UUID {
	id, _ := uuid.Parse(l.ListOwner)
	return id
}
func (l mongoList) Elements() []uuid.UUID {
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

func (t todoMongoRepo) GetListByID(id uuid.UUID) (domain.List, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id.String()}}

	var result mongoList
	if err := t.lists.FindOne(ctx, filter).Decode(&result); err != nil {
		return domain.NewList("", uuid.UUID{}), errors.New("list is not found")
	}
	return result, nil
}

func (t todoMongoRepo) CreateList(list domain.List) error {
	if _, err := t.GetListByID(list.Id()); err == nil {
		return errors.New("list already exists")
	}

	mList := mongoList {
		ID: list.Id().String(),
		ListName: list.Name(),
		ListOwner: list.Owner().String(),
	}
	for i, v := range list.Elements() {
		mList.ElementsID[i] = v.String()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if _, err := t.lists.InsertOne(ctx, mList); err != nil {
		return errors.New("internal error, failed to create a list")
	}
	return nil
}

func (t todoMongoRepo) GetListsByOwner(id uuid.UUID) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id.String()}}

	cursor, err := t.lists.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("internal error")
	}

	var mLists []mongoList
	if err = cursor.All(context.TODO(), &mLists); err != nil {
		return nil, errors.New("internal error")
	}

	result := make([]uuid.UUID, len(mLists))
	for i, v := range mLists {
		result[i] = v.Id()
	}
	return result, nil
}

func (t todoMongoRepo) DeleteList(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	filter := bson.D{{Key:"_id", Value: id.String()}}

	if _, err := t.lists.DeleteOne(ctx, filter); err != nil {
		return errors.New("user was not found")
	}
	return nil	
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
