package main

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
)

type SubTodo struct {
	Id       string `json:"id"`
	ParentId string `json:"parentId"`
	Title    string `json:"title"`
	IsDone   bool   `json:"isDone"`
}

var subTodos SubTodos

type SubTodos map[string]*SubTodo

func newSubTodo(t SubTodo) *SubTodo {
	return &SubTodo{
		Id:       t.Id,
		ParentId: t.ParentId,
		Title:    t.Title,
		IsDone:   t.IsDone,
	}
}

func (t *SubTodo) noId() interface{} {
	return struct {
		Title  string `json:"title"`
		IsDone bool   `json:"isDone"`
	}{
		Title:  t.Title,
		IsDone: t.IsDone,
	}
}

func (t *SubTodos) retrieve(id string) *SubTodo {
	if v, ok := (*t)[id]; ok {
		return v
	}
	var subTodo SubTodo
	(*t)[id] = &subTodo
	return &subTodo
}

func (t *SubTodos) retrieveAll() *SubTodos {
	rows, err := r.Table("subtodo").Run(session)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sliceSubTodos []SubTodo
	err2 := rows.All(&sliceSubTodos)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}

	for _, st := range sliceSubTodos {
		subTodos[st.Id] = newSubTodo(st)
	}
	return t
}

func (ts *SubTodos) add(newSubTodo *SubTodo) {
	newSubTodo.Id = insert("subtodo", newSubTodo.noId())

	subTodos[newSubTodo.Id] = newSubTodo
	todos[(newSubTodo.ParentId)].SubTodos[(newSubTodo.Id)] = newSubTodo
	todos[(newSubTodo.ParentId)].UpdatedAt = time.Now()

	update("todo", newSubTodo.ParentId, todos[newSubTodo.ParentId])
}
