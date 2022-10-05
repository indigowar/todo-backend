package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/indigowar/todo-backend/internal/domain"
	"github.com/indigowar/todo-backend/internal/repository"
	"github.com/indigowar/todo-backend/pkg/auth"
)

type UserService interface {
	// CreateUser - create a new user.
	// Method returns the refresh token and an error.
	CreateUser(ctx context.Context, name, password string) (string, error)
	// DeleteUser - deletes user from service,
	// gets access token and returns an error.
	DeleteUser(ctx context.Context, token string) error
	// GetName - returns the name of token's owner or an error
	GetName(ctx context.Context, token string) (string, error)
	// UpdatePassword - updates user's password by user's token,
	// Returns an error
	UpdatePassword(ctx context.Context, token string, password string) error
	// Login - login in the account, get access and refresh tokens
	Login(ctx context.Context, name string, password string) (string, string, error)
}

type userService struct {
	user repository.UserRepo
	todo repository.TodoRepo
	auth auth.TokenManager
}

func NewUserService(r repository.UserRepo, todo repository.TodoRepo, a auth.TokenManager) UserService {
	return &userService{
		user: r,
		todo: todo,
		auth: a,
	}
}

func (service *userService) CreateUser(ctx context.Context, name, password string) (string, error) {
	// prepare user entity
	id := uuid.New()
	user := domain.NewUser(id, name, password)

	// create user in storage
	if err := service.user.Create(ctx, user); err != nil {
		return "", err
	}

	_, refresh, err := service.Login(ctx, name, password)
	if err != nil {
		return "", err
	}

	return refresh, nil
}

func (service *userService) DeleteUser(ctx context.Context, token string) error {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return errors.New("invalid token")
	}

	if !available {
		return errors.New("token is expired")
	}

	lists, err := service.todo.GetListsByOwner(ctx, id)
	if err == nil {
		for _, v := range lists {
			_ = service.todo.DeleteList(ctx, v)
		}
	}

	return service.user.Delete(ctx, id)
}

func (service *userService) GetName(ctx context.Context, token string) (string, error) {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return "", errors.New("invalid token")
	}

	if !available {
		return "", errors.New("token is expired")
	}

	user, err := service.user.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return user.Name(), nil
}

func (service *userService) UpdatePassword(ctx context.Context, token, password string) error {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return errors.New("invalid token")
	}

	if !available {
		return errors.New("token is expired")
	}

	if len(password) < 6 {
		return errors.New("password too simple")
	}

	return service.user.UpdatePassword(ctx, id, password)
}

func (service *userService) UpdateName(ctx context.Context, token, name string) error {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return errors.New("invalid token")
	}

	if !available {
		return errors.New("token is expired")
	}

	return service.user.UpdateUserName(ctx, id, name)
}

func (service *userService) NewAccessToken(ctx context.Context, refresh string) (string, error) {
	if refresh == "" {
		return "", errors.New("invalid token")
	}

	id, expired, err := service.user.GetByRefresh(ctx, refresh)

	if err != nil {
		return "", errors.New("invalid token")
	}

	if expired.After(time.Now()) {
		return "", errors.New("token is expired")
	}

	return service.auth.NewJWT(id)
}

func (service *userService) Login(ctx context.Context, name, password string) (string, string, error) {
	id, err := service.user.GetByName(ctx, name)
	if err != nil {
		return "", "", errors.New("login error")
	}
	user, _ := service.user.Get(ctx, id)

	if user.Password() != password {
		return "", "", errors.New("login error")
	}

	refresh, expiredAt, err := service.auth.NewLongLive(user.Id())
	if err != nil {
		return "", "", errors.New("internal server error")
	}
	jwt, err := service.auth.NewJWT(user.Id())
	if err != nil {
		return "", "", errors.New("internal server error")
	}

	if service.user.SetRefresh(ctx, user.Id(), refresh, expiredAt) != nil {
		return "", "", errors.New("internal server error")
	}

	return jwt, refresh, nil
}
