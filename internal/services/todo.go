package services

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/todo-backend/internal/domain"
	"github.com/indigowar/todo-backend/internal/repository"
	"github.com/indigowar/todo-backend/pkg/auth"
)

type TodoService interface {
	// GetList - get list that user owns or get an error
	GetList(context.Context, string, uuid.UUID) (domain.List, error)
	// GetLists - get list of users list
	GetLists(context.Context, string) ([]uuid.UUID, error)
	// CreateList - create a new list for user
	CreateList(context.Context, string, string) error
	// DeleteList - delete a list by it's ID
	DeleteList(context.Context, string, uuid.UUID) error
	// GetElement - get element of list
	GetElement(context.Context, string, uuid.UUID, uuid.UUID) (domain.Element, error)
	AddElement(context.Context, string, uuid.UUID, string) error
	DeleteElement(context.Context, string, uuid.UUID, uuid.UUID) error
	ChangeElementStatus(context.Context, string, uuid.UUID, uuid.UUID) error
	RenameElement(context.Context, string, uuid.UUID, uuid.UUID, string) error
}

type todoService struct {
	users repository.UserRepo
	todo  repository.TodoRepo
	auth  auth.TokenManager
}

func NewTodoService(users repository.UserRepo, todo repository.TodoRepo, manager auth.TokenManager) TodoService {
	return &todoService{
		users: users,
		todo:  todo,
		auth:  manager,
	}
}

func (service *todoService) GetList(ctx context.Context, token string, ownerId uuid.UUID) (domain.List, error) {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return domain.NewList("", uuid.UUID{}), err
	}

	list, err := service.todo.GetListByID(ownerId)
	if err != nil {
		return domain.NewList("", uuid.UUID{}), errors.New("no access")
	}

	if list.Owner() != ownerId {
		return domain.NewList("", uuid.UUID{}), errors.New("no access")
	}

	return list, nil
}

func (service *todoService) GetLists(ctx context.Context, token string) ([]uuid.UUID, error) {
	id, err := service.verifyUserAccess(token)
	if err != nil {
		return nil, err
	}

	lists, err := service.todo.GetListsByOwner(id)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return lists, nil
}

func (service *todoService) CreateList(ctx context.Context, token string, name string) error {
	id, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	err = service.todo.CreateList(domain.NewList(name, id))
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (service *todoService) DeleteList(ctx context.Context, token string, id uuid.UUID) error {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	list, err := service.todo.GetListByID(id)
	if err != nil {
		return err
	}
	if list.Owner() != ownerId {
		return errors.New("no access")
	}

	return service.todo.DeleteList(id)
}

func (service *todoService) GetElement(ctx context.Context, token string, list, element uuid.UUID) (domain.Element, error) {
	l, err := service.GetList(ctx, token, list)
	if err != nil {
		return domain.NewElement(""), err
	}

	containsElement := false
	for _, v := range l.Elements() {
		if v == element {
			containsElement = true
			break
		}
	}

	if !containsElement {
		return domain.NewElement(""), errors.New("element was not found")
	}

	return service.todo.GetElement(element)
}

func (service *todoService) AddElement(ctx context.Context, token string, list uuid.UUID, value string) error {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	ownersLists, err := service.todo.GetListsByOwner(ownerId)
	if err != nil {
		return errors.New("internal error")
	}
	listFound := false
	for _, v := range ownersLists {
		if list == v {
			listFound = true
			break
		}
	}
	if !listFound {
		return errors.New("no access")
	}

	return service.todo.AddElement(list, domain.NewElement(value))
}

func (service *todoService) DeleteElement(ctx context.Context, token string, list, element uuid.UUID) error {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	ownersLists, err := service.todo.GetListsByOwner(ownerId)
	if err != nil {
		return errors.New("internal error")
	}
	listFound := false
	for _, v := range ownersLists {
		if list == v {
			listFound = true
			break
		}
	}

	if !listFound {
		return errors.New("no access")
	}

	return service.todo.DeleteElement(list, element)
}

func (service *todoService) ChangeElementStatus(ctx context.Context, token string, list, element uuid.UUID) error {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	ownersLists, err := service.todo.GetListsByOwner(ownerId)
	if err != nil {
		return errors.New("internal error")
	}
	listFound := false
	for _, v := range ownersLists {
		if list == v {
			listFound = true
			break
		}
	}

	if !listFound {
		return errors.New("no access")
	}

	return service.todo.ChangeStatus(list, element)
}

func (service *todoService) RenameElement(ctx context.Context, token string, list, element uuid.UUID, value string) error {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	ownersLists, err := service.todo.GetListsByOwner(ownerId)
	if err != nil {
		return errors.New("internal error")
	}
	listFound := false
	for _, v := range ownersLists {
		if list == v {
			listFound = true
			break
		}
	}

	if !listFound {
		return errors.New("no access")
	}

	return service.todo.RenameElement(list, element, value)
}

func (service *todoService) verifyUserAccess(token string) (uuid.UUID, error) {
	id, available, err := service.auth.Verify(token)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid token")
	}
	if !available {
		return uuid.UUID{}, errors.New("token is expired")
	}

	return id, nil
}
