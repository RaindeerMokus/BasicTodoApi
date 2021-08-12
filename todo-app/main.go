package main

import (
	"encoding/json"
	"fmt"

	r "github.com/dancannon/gorethink"
)

var err error

var session *r.Session

func main() {
	printStr("#################################################")
	printStr("#################################################")
	printStr("#################################################")

	initDB()

	todos = make(Todos)
	subTodos = make(SubTodos)
	todos.retrieveAll()
	subTodos.retrieveAll()

	initRouter()
}

func printStr(v string) {
	fmt.Println(v)
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
