package domain

import (
	"github.com/google/uuid"
	"time"
)

// User represents a user entity in this system
type User interface {
	Id() uuid.UUID
	Name() string
	Password() string
	TokenValue() string
	TokenExpiredTime() time.Time
}

func NewUser(id uuid.UUID, name, password string) User {
	// todo: checks for valid args

	return &user {
		id: id,
		name: name,
		password: password,
	}
}

type user struct {
	id uuid.UUID
	name string
	password string
	tokenValue string
	tokenExpiredTime time.Time
}

func (u *user) Id() uuid.UUID {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Password() string {
	return u.password
}

func (u *user) TokenValue() string {
	return u.tokenValue
}

func (u *user) TokenExpiredTime() time.Time {
	return u.tokenExpiredTime
}
