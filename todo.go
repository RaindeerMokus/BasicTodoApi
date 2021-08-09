package main

import (
	"encoding/json"
	"time"
)

type Todo struct {
	Id        uint64              `json:"id"`
	Title     string              `json:"title"`
	IsDone    bool                `json:"isDone"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
	SubTodos  map[uint64]*SubTodo `json:"-"`
}

func (t Todo) MarshalJSON() ([]byte, error) {
	type Alias Todo

	v := []*SubTodo{}

	for _, value := range t.SubTodos {
		v = append(v, value)
	}

	return json.Marshal(&struct {
		Alias
		SubTodos []*SubTodo `json:"subTodos"`
	}{
		Alias:    (Alias)(t),
		SubTodos: v,
	})
}

type Todos map[uint64]*Todo

func (t Todos) MarshalJSON() ([]byte, error) {

	v := []*Todo{}

	for _, value := range t {
		v = append(v, value)
	}

	return json.Marshal(&struct {
		Todos []*Todo `json:"Todos"`
	}{
		Todos: v,
	})
}
