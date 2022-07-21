package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/indigowar/todo-backend/internal/domain"
	"github.com/indigowar/todo-backend/internal/repository"
	"github.com/indigowar/todo-backend/pkg/auth"
)

type TodoService interface {
	// GetList - get list that user owns or get an error
	//
	GetList(token string, id uuid.UUID) (domain.List, error)
	// GetLists - get list of users list
	GetLists(string) ([]uuid.UUID, error)
	// CreateList - create a new list for user
	CreateList(string, string) error
	// DeleteList - delete a list by it's ID
	DeleteList(string, uuid.UUID) error
	// GetElement - get element of list
	GetElement(string, uuid.UUID, uuid.UUID) (domain.Element, error)
	AddElement(string, uuid.UUID, string) error
	DeleteElement(string, uuid.UUID, uuid.UUID) error
	ChangeElementStatus(string, uuid.UUID, uuid.UUID) error
	RenameElement(string, uuid.UUID, uuid.UUID, string) error
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

func (service *todoService) GetList(token string, ownerId uuid.UUID) (domain.List, error) {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return domain.List{}, err
	}

	list, err := service.todo.GetListByID(ownerId)
	if err != nil {
		return domain.List{}, errors.New("no access")
	}

	if list.Owner != ownerId {
		return domain.List{}, errors.New("no access")
	}

	return list, nil
}

func (service *todoService) GetLists(token string) ([]uuid.UUID, error) {
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

func (service *todoService) CreateList(token string, name string) error {
	id, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	err = service.todo.CreateList(domain.List{
		Id:    uuid.New(),
		Name:  name,
		Owner: id,
	})
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (service *todoService) DeleteList(token string, id uuid.UUID) error {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return err
	}

	list, err := service.todo.GetListByID(id)
	if err != nil {
		return err
	}
	if list.Owner != ownerId {
		return errors.New("no access")
	}

	return service.todo.DeleteList(id)
}

func (service *todoService) GetElement(token string, list, element uuid.UUID) (domain.Element, error) {
	ownerId, err := service.verifyUserAccess(token)
	if err != nil {
		return domain.Element{}, err
	}

	ownersLists, err := service.todo.GetListsByOwner(ownerId)
	if err != nil {
		return domain.Element{}, errors.New("internal error")
	}
	listFound := false
	for _, v := range ownersLists {
		if list == v {
			listFound = true
			break
		}
	}
	if !listFound {
		return domain.Element{}, errors.New("no access")
	}

	return service.todo.GetElement(list, element)
}

func (service *todoService) AddElement(token string, list uuid.UUID, value string) error {
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

	return service.todo.AddElement(list, domain.Element{
		Id:    uuid.New(),
		Value: value,
		Done:  false,
	})
}

func (service *todoService) DeleteElement(token string, list, element uuid.UUID) error {
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

func (service *todoService) ChangeElementStatus(token string, list, element uuid.UUID) error {
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

func (service *todoService) RenameElement(token string, list, element uuid.UUID, value string) error {
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
