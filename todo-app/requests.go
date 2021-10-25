package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/todos", getTodos)
	router.POST("/newTodo", postTodo)
	router.POST("/newSubTodo", postSubTodo)
	router.PUT("/renameTodo", putRenameTodo)
	router.PUT("/renameSubTodo", putRenameSubTodo)
	router.PUT("/tickTodo", putTickTodo)
	router.PUT("/tickSubTodo", putTickSubTodo)
	router.PUT("/moveSubTodo", putMoveSubTodo)

	router.Run("0.0.0.0:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos.retrieveAll())
}

func postTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todos.add(&newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func postSubTodo(c *gin.Context) {
	var newSubTodo SubTodo

	if err := c.BindJSON(&newSubTodo); err != nil {
		return
	}
	subTodos.add(&newSubTodo)
	c.IndentedJSON(http.StatusCreated, newSubTodo)
}

func putRenameTodo(c *gin.Context) {
	var newNameTodo Todo

	if err := c.BindJSON(&newNameTodo); err != nil {
		return
	}

	todos[newNameTodo.Id].Title = newNameTodo.Title
	todos[newNameTodo.Id].UpdatedAt = time.Now()

	update("todo", newNameTodo.Id, todos[newNameTodo.Id])

	c.IndentedJSON(http.StatusCreated, todos[newNameTodo.Id])
}

func putRenameSubTodo(c *gin.Context) {
	var newNameSubTodo SubTodo

	if err := c.BindJSON(&newNameSubTodo); err != nil {
		return
	}

	subTodos[newNameSubTodo.Id].Title = newNameSubTodo.Title
	update("subtodo", newNameSubTodo.Id, subTodos[newNameSubTodo.Id])
	update("todo", subTodos[newNameSubTodo.Id].ParentId, todos[subTodos[newNameSubTodo.Id].ParentId])

	c.IndentedJSON(http.StatusCreated, subTodos[newNameSubTodo.Id])
}

func putTickTodo(c *gin.Context) {
	var tickTodo Todo

	if err := c.BindJSON(&tickTodo); err != nil {
		return
	}

	todos[tickTodo.Id].IsDone = !todos[tickTodo.Id].IsDone
	todos[tickTodo.Id].UpdatedAt = time.Now()

	update("todo", tickTodo.Id, todos[tickTodo.Id])

	c.IndentedJSON(http.StatusCreated, todos[tickTodo.Id])
}

func putTickSubTodo(c *gin.Context) {
	var tickSubTodo SubTodo

	if err := c.BindJSON(&tickSubTodo); err != nil {
		return
	}

	subTodos[tickSubTodo.Id].IsDone = !subTodos[tickSubTodo.Id].IsDone

	update("subtodo", tickSubTodo.Id, subTodos[tickSubTodo.Id])
	update("todo", subTodos[tickSubTodo.Id].ParentId, todos[subTodos[tickSubTodo.Id].ParentId])

	c.IndentedJSON(http.StatusCreated, subTodos[tickSubTodo.Id])
}

func putMoveSubTodo(c *gin.Context) {
	var moveSubTodo SubTodo

	if err := c.BindJSON(&moveSubTodo); err != nil {
		return
	}

	oldParentTodo := todos[subTodos[moveSubTodo.Id].ParentId]
	delete(oldParentTodo.SubTodos, moveSubTodo.Id)

	subTodos[moveSubTodo.Id].ParentId = moveSubTodo.ParentId
	todos[moveSubTodo.ParentId].SubTodos[moveSubTodo.Id] = subTodos[moveSubTodo.Id]

	printObj(oldParentTodo)
	printObj(todos[moveSubTodo.ParentId])

	update("subtodo", moveSubTodo.Id, subTodos[moveSubTodo.Id])
	replace("todo", oldParentTodo.Id, oldParentTodo)
	update("todo", moveSubTodo.ParentId, todos[moveSubTodo.ParentId])

	c.IndentedJSON(http.StatusCreated, subTodos[moveSubTodo.Id])
}
