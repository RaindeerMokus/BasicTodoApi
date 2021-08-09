package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var todos Todos
var maxID uint64
var subTodos SubTodos
var maxsubID uint64

func main() {
	todos = make(Todos)
	subTodos = make(SubTodos)

	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/newTodo", postTodo)
	router.POST("/newSubTodo", postSubTodo)
	router.PUT("/renameTodo", putRenameTodo)
	router.PUT("/renameSubTodo", putRenameSubTodo)
	router.PUT("/tickTodo", putTickTodo)
	router.PUT("/tickSubTodo", putTickSubTodo)
	router.PUT("/moveSubTodo", putMoveSubTodo)

	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo.Id = maxID
	maxID++
	newTodo.CreatedAt = time.Now()
	newTodo.SubTodos = make(map[uint64]*SubTodo)

	todos[newTodo.Id] = &newTodo
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func postSubTodo(c *gin.Context) {
	var newSubTodo SubTodo

	if err := c.BindJSON(&newSubTodo); err != nil {
		return
	}
	newSubTodo.Id = maxsubID
	maxsubID++

	subTodos[newSubTodo.Id] = &newSubTodo
	todos[newSubTodo.ParentId].SubTodos[newSubTodo.Id] = &newSubTodo
	c.IndentedJSON(http.StatusCreated, newSubTodo)
}

func putRenameTodo(c *gin.Context) {
	var newNameTodo Todo

	if err := c.BindJSON(&newNameTodo); err != nil {
		return
	}

	todos[newNameTodo.Id].Title = newNameTodo.Title
	todos[newNameTodo.Id].UpdatedAt = time.Now()

	c.IndentedJSON(http.StatusCreated, todos[newNameTodo.Id])
}

func putRenameSubTodo(c *gin.Context) {
	var newNameSubTodo SubTodo

	if err := c.BindJSON(&newNameSubTodo); err != nil {
		return
	}

	subTodos[newNameSubTodo.Id].Title = newNameSubTodo.Title

	c.IndentedJSON(http.StatusCreated, subTodos[newNameSubTodo.Id])
}

func putTickTodo(c *gin.Context) {
	var tickTodo Todo

	if err := c.BindJSON(&tickTodo); err != nil {
		return
	}

	todos[tickTodo.Id].IsDone = !todos[tickTodo.Id].IsDone
	todos[tickTodo.Id].UpdatedAt = time.Now()

	c.IndentedJSON(http.StatusCreated, todos[tickTodo.Id])
}

func putTickSubTodo(c *gin.Context) {
	var tickSubTodo SubTodo

	if err := c.BindJSON(&tickSubTodo); err != nil {
		return
	}

	subTodos[tickSubTodo.Id].IsDone = !subTodos[tickSubTodo.Id].IsDone

	c.IndentedJSON(http.StatusCreated, subTodos[tickSubTodo.Id])
}

func putMoveSubTodo(c *gin.Context) {
	var moveSubTodo SubTodo

	if err := c.BindJSON(&moveSubTodo); err != nil {
		return
	}

	delete(todos[subTodos[moveSubTodo.Id].ParentId].SubTodos, moveSubTodo.Id)

	subTodos[moveSubTodo.Id].ParentId = moveSubTodo.ParentId
	todos[moveSubTodo.ParentId].SubTodos[moveSubTodo.Id] = subTodos[moveSubTodo.Id]

	c.IndentedJSON(http.StatusCreated, subTodos[moveSubTodo.Id])
}
