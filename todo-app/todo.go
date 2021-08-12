package main

import (
	"encoding/json"
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
)

type Todo struct {
	Id        string              `json:"id"`
	Title     string              `json:"title"`
	IsDone    bool                `json:"isDone"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
	SubTodos  map[string]*SubTodo `json:"subTodos"`
}

var todos Todos

func newTodo(t Todo) *Todo {
	return &Todo{
		Id:        t.Id,
		Title:     t.Title,
		IsDone:    t.IsDone,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		SubTodos:  t.SubTodos,
	}
}

func (t *Todo) /*Back with old friends like */ noId() /*who is tonight's director*/ interface{} {
	return struct {
		Title     string              `json:"title"`
		IsDone    bool                `json:"isDone"`
		CreatedAt time.Time           `json:"createdAt"`
		UpdatedAt time.Time           `json:"updatedAt"`
		SubTodos  map[string]*SubTodo `json:"subTodos"`
	}{
		Title:     t.Title,
		IsDone:    t.IsDone,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		SubTodos:  t.SubTodos,
	}
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

type Todos map[string]*Todo

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

func (t *Todos) retrieve(id string) *Todo {
	if v, ok := (*t)[id]; ok {
		return v
	}
	var todo Todo
	(*t)[id] = &todo
	return &todo
}

func (t *Todos) retrieveAll() *Todos {
	rows, err := r.Table("todo").Run(session)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sliceTodos []Todo
	err2 := rows.All(&sliceTodos)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}

	for _, st := range sliceTodos {
		todos[st.Id] = newTodo(st)
		if len(todos[st.Id].SubTodos) > 0 {
			for _, sst := range todos[st.Id].SubTodos {
				subTodos[sst.Id] = sst
			}
		}
	}
	return t
}

func (ts *Todos) add(newTodo *Todo) {

	newTodo.CreatedAt = time.Now()
	newTodo.SubTodos = make(map[string]*SubTodo)
	newTodo.Id = insert("todo", newTodo.noId())
	todos[newTodo.Id] = newTodo

}
