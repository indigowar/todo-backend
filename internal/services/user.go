package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/indigowar/todo-backend/internal/domain"
	"github.com/indigowar/todo-backend/internal/repository"
	"github.com/indigowar/todo-backend/pkg/auth"
	"time"
)

type UserService interface {
	// CreateUser - create a new user.
	// Method returns the refresh token and an error.
	CreateUser(name, password string) (string, error)
	// DeleteUser - deletes user from service,
	// gets access token and returns an error.
	DeleteUser(token string) error
	// GetName - returns the name of token's owner or an error
	GetName(token string) (string, error)
	// UpdatePassword - updates user's password by user's token,
	// Returns an error
	UpdatePassword(string, string) error
	// Login - login in the account, get access and refresh tokens
	Login(string, string) (string, string, error)
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

func (service *userService) CreateUser(name, password string) (string, error) {
	id := uuid.New()
	refresh, expTime, err := service.auth.NewLongLive(id)
	if err != nil {
		return "", errors.New("internal server error")
	}

	user := domain.User{
		Id:       id,
		Name:     name,
		Password: password,
	}
	user.Token.Value = refresh
	user.Token.ExpiredAt = expTime

	if err := service.user.Create(user); err != nil {
		return "", err
	}

	return user.Token.Value, nil
}

func (service *userService) DeleteUser(token string) error {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return errors.New("invalid token")
	}

	if !available {
		return errors.New("token is expired")
	}

	lists, err := service.todo.GetListsByOwner(id)
	if err == nil {
		for _, v := range lists {
			_ = service.todo.DeleteList(v)
		}
	}

	return service.user.Delete(id)
}

func (service *userService) GetName(token string) (string, error) {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return "", errors.New("invalid token")
	}

	if !available {
		return "", errors.New("token is expired")
	}

	user, err := service.user.Get(id)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func (service *userService) UpdatePassword(token, password string) error {
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

	return service.user.UpdatePassword(id, password)
}

func (service *userService) UpdateName(token, name string) error {
	id, available, err := service.auth.Verify(token)

	if err != nil {
		return errors.New("invalid token")
	}

	if !available {
		return errors.New("token is expired")
	}

	return service.user.UpdateUserName(id, name)
}

func (service *userService) NewAccessToken(refresh string) (string, error) {
	//id, available, err := service.auth.Verify(refresh)
	//
	//if err != nil {
	//	return "", errors.New("invalid token")
	//}
	//
	//if !available {
	//	return "", errors.New("token is expired")
	//}
	//
	//if token, err := service.auth.NewJWT(id); err != nil {
	//	return "", errors.New("server internal error")
	//} else {
	//	return token, nil
	//}
	if refresh == "" {
		return "", errors.New("invalid token")
	}

	id, expired, err := service.user.GetByRefresh(refresh)

	if err != nil {
		return "", errors.New("invalid token")
	}

	if expired.After(time.Now()) {
		return "", errors.New("token is expired")
	}

	return service.auth.NewJWT(id)
}

func (service *userService) Login(name, password string) (string, string, error) {
	id, err := service.user.GetByName(name)
	if err != nil {
		return "", "", errors.New("login error")
	}
	user, _ := service.user.Get(id)

	if user.Password != password {
		return "", "", errors.New("login error")
	}

	refresh, expiredAt, err := service.auth.NewLongLive(user.Id)
	if err != nil {
		return "", "", errors.New("internal server error")
	}
	jwt, err := service.auth.NewJWT(user.Id)
	if err != nil {
		return "", "", errors.New("internal server error")
	}

	if service.user.SetRefresh(user.Id, refresh, expiredAt) != nil {
		return "", "", errors.New("internal server error")
	}

	return jwt, refresh, nil
}
