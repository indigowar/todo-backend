package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/indigowar/todo-backend/internal/domain"
)

func NewTodoRepo(database *mongo.Database) (TodoRepo, error) {
	return &todoMongoRepo{
		lists:    database.Collection("lists"),
		elements: database.Collection("elements"),
	}, nil
}

type mongoElement struct {
	ID           string `bson:"_id"`
	ElementValue string `bson:"value"`
	Status       bool   `bson:"done"`
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
	ID         string   `bson:"_id"`
	ListName   string   `bson:"name"`
	ListOwner  string   `bson:"owner"`
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
	lists    *mongo.Collection
	elements *mongo.Collection
}

func (t todoMongoRepo) GetListByID(ctx context.Context, id uuid.UUID) (domain.List, error) {
	filter := bson.D{{Key: "_id", Value: id.String()}}

	var result mongoList
	if err := t.lists.FindOne(ctx, filter).Decode(&result); err != nil {
		return domain.NewList("", uuid.UUID{}), errors.New("list is not found")
	}
	return result, nil
}

func (t todoMongoRepo) CreateList(ctx context.Context, list domain.List) error {
	if _, err := t.GetListByID(ctx, list.Id()); err == nil {
		return errors.New("list already exists")
	}

	mList := mongoList{
		ID:        list.Id().String(),
		ListName:  list.Name(),
		ListOwner: list.Owner().String(),
	}
	for i, v := range list.Elements() {
		mList.ElementsID[i] = v.String()
	}

	if _, err := t.lists.InsertOne(ctx, mList); err != nil {
		return errors.New("internal error, failed to create a list")
	}
	return nil
}

func (t todoMongoRepo) GetListsByOwner(ctx context.Context, id uuid.UUID) ([]uuid.UUID, error) {
	filter := bson.D{{Key: "_id", Value: id.String()}}

	cursor, err := t.lists.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("internal error")
	}

	var mLists []mongoList
	if err = cursor.All(ctx, &mLists); err != nil {
		return nil, errors.New("internal error")
	}

	result := make([]uuid.UUID, len(mLists))
	for i, v := range mLists {
		result[i] = v.Id()
	}
	return result, nil
}

func (t todoMongoRepo) DeleteList(ctx context.Context, id uuid.UUID) error {
	filter := bson.D{{Key: "_id", Value: id.String()}}

	if _, err := t.lists.DeleteOne(ctx, filter); err != nil {
		return errors.New("user was not found")
	}
	return nil
}

func (t todoMongoRepo) GetElement(ctx context.Context, id uuid.UUID) (domain.Element, error) {
	filter := bson.D{{Key: "_id", Value: id.String()}}
	var result mongoElement

	if err := t.elements.FindOne(ctx, filter).Decode(&result); err != nil {
		return domain.NewElement(""), errors.New("element was not found")
	}

	return result, nil
}

func (t todoMongoRepo) AddElement(ctx context.Context, listId uuid.UUID, element domain.Element) error {
	list, err := t.GetListByID(ctx, listId)
	if err == nil {
		return errors.New("list does not exist")
	}

	var mElement = mongoElement{
		ID:           element.Id().String(),
		ElementValue: element.Value(),
		Status:       element.Done(),
	}

	elements := make([]string, len(list.Elements())+1)
	for i, v := range list.Elements() {
		elements[i] = v.String()
	}
	elements[len(elements)-1] = element.Id().String()

	filter := bson.D{{Key: "_id", Value: listId.String()}}
	updater := bson.D{{Key: "elements", Value: elements}}

	var callback = func(sessCtx mongo.SessionContext) (interface{}, error) {
		if _, err := t.elements.InsertOne(sessCtx, mElement); err != nil {
			return nil, errors.New("failed to insert element")
		}
		if _, err := t.lists.UpdateOne(sessCtx, filter, updater); err != nil {
			return nil, errors.New("failed to update list")
		}
		return nil, nil
	}

	session, err := t.elements.Database().Client().StartSession()
	if err != nil {
		return errors.New("failed to get into database")
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return errors.New("failed to add element")
	}
	return nil
}

func (t todoMongoRepo) DeleteElement(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) ChangeStatus(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (t todoMongoRepo) RenameElement(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID, s string) error {
	//TODO implement me
	panic("implement me")
}
