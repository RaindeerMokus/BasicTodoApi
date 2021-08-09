package main

type SubTodo struct {
	Id       uint64 `json:"id"`
	ParentId uint64 `json:"parentId"`
	Title    string `json:"title"`
	IsDone   bool   `json:"isDone"`
}

type SubTodos map[uint64]*SubTodo
