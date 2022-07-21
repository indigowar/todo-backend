package domain

import "github.com/google/uuid"

type Element struct {
	Id    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Done  bool      `json:"done"`
}

type List struct {
	Id       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Owner    uuid.UUID   `json:"owner"`
	Elements []uuid.UUID `json:"elements"`
}
