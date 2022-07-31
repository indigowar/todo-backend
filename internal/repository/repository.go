package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/indigowar/todo-backend/internal/domain"
	"time"
)

var (
	ErrUserNotFound      = errors.New("user was not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserRepo interface {
	Get(uuid.UUID) (domain.User, error)
	GetByName(string) (uuid.UUID, error)
	Create(domain.User) error
	Delete(uuid.UUID) error

	UpdatePassword(uuid.UUID, string) error
	UpdateUserName(uuid.UUID, string) error

	SetRefresh(uuid.UUID, string, time.Time) error
	GetByRefresh(string) (uuid.UUID, time.Time, error)
}

var (
	ErrListNotFound         = errors.New("list was not found")
	ErrListAlreadyExists    = errors.New("list already exists")
	ErrElementNotFound      = errors.New("element was not found")
	ErrElementAlreadyExists = errors.New("element already exists")
)

type TodoRepo interface {
	GetListByID(uuid.UUID) (domain.List, error)
	CreateList(domain.List) error
	GetListsByOwner(uuid.UUID) ([]uuid.UUID, error)
	DeleteList(uuid.UUID) error

	GetElement(uuid.UUID, uuid.UUID) (domain.Element, error)
	AddElement(uuid.UUID, domain.Element) error
	DeleteElement(uuid.UUID, uuid.UUID) error
	ChangeStatus(uuid.UUID, uuid.UUID) error
	RenameElement(uuid.UUID, uuid.UUID, string) error
}
