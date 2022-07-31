package domain

import "github.com/google/uuid"

type Element interface {
	Id() uuid.UUID
	Value() string
	Done() bool
}

type List interface {
	Id() uuid.UUID
	Name() string
	Owner() uuid.UUID
	Elements() []uuid.UUID
}

func NewElement(value string) Element {
	return &element{
		id: uuid.New(),
		value: value,
		done: false,
	}
}

func NewList(name string, owner uuid.UUID) List {
	return &list{
		id: uuid.New(),
		name: name,
		owner: owner,
		elements: make([]uuid.UUID, 0),
	}
}

type element struct {
	id    uuid.UUID
	value string
	done  bool
}

func (e *element) Id() uuid.UUID {
	return e.id
}

func (e *element) Value() string {
	return e.value
}

func (e *element) Done() bool {
	return e.done
}

type list struct {
	id       uuid.UUID
	name     string
	owner    uuid.UUID
	elements []uuid.UUID
}

func (l *list) Id() uuid.UUID {
	return l.id
}

func (l *list) Name() string {
	return l.name
}

func (l *list) Owner() uuid.UUID {
	return l.owner
}

func (l *list) Elements() []uuid.UUID {
	return l.elements
}
