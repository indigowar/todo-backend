package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Token    struct {
		Value     string    `json:"value"`
		ExpiredAt time.Time `json:"expired_at"`
	} `json:"token"`
}
