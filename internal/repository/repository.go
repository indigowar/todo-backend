package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/indigowar/todo-backend/internal/domain"
)

var (
	ErrUserNotFound      = errors.New("user was not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserRepo interface {
	Get(context.Context, uuid.UUID) (domain.User, error)
	GetByName(context.Context, string) (uuid.UUID, error)
	Create(context.Context, domain.User) error
	Delete(context.Context, uuid.UUID) error

	UpdatePassword(context.Context, uuid.UUID, string) error
	UpdateUserName(context.Context, uuid.UUID, string) error

	SetRefresh(context.Context, uuid.UUID, string, time.Time) error
	GetByRefresh(context.Context, string) (uuid.UUID, time.Time, error)
}

var (
	ErrListNotFound         = errors.New("list was not found")
	ErrListAlreadyExists    = errors.New("list already exists")
	ErrElementNotFound      = errors.New("element was not found")
	ErrElementAlreadyExists = errors.New("element already exists")
)

type TodoRepo interface {
	GetListByID(context.Context, uuid.UUID) (domain.List, error)
	CreateList(context.Context, domain.List) error
	GetListsByOwner(context.Context, uuid.UUID) ([]uuid.UUID, error)
	DeleteList(context.Context, uuid.UUID) error

	GetElement(context.Context, uuid.UUID) (domain.Element, error)
	AddElement(context.Context, uuid.UUID, domain.Element) error
	DeleteElement(context.Context, uuid.UUID, uuid.UUID) error
	ChangeStatus(context.Context, uuid.UUID, uuid.UUID) error
	RenameElement(context.Context, uuid.UUID, uuid.UUID, string) error
}
