package main

import (
	"fmt"
	"os"

	r "github.com/dancannon/gorethink"
)

func initDB() {
	r.SetTags("rethinkdb", "json")
	dbAdress := os.Getenv("DB_SERVER")

	session, err = r.Connect(r.ConnectOpts{
		Address:  (dbAdress + ":28015"),
		Database: "test",
	})
	if err != nil {
		fmt.Println(err)
	}

	result, err := r.DB("test").TableCreate("todo").RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}
	result, err = r.DB("test").TableCreate("subtodo").IndexCreate("parentId").RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printObj(result)
}

func insert(table string, obj interface{}) string {
	result, err := r.Table(table).Insert(obj, r.InsertOpts{ReturnChanges: true}).RunWrite(session)
	if err != nil {
		fmt.Println(err)
	}

	printObj(result)
	fmt.Println(result.GeneratedKeys)
	return result.GeneratedKeys[0]
}

func update(table string, id string, obj interface{}) {
	result, err := r.Table(table).Get(id).Update(obj).RunWrite(session)
	r.Table(table).Get(id).Update(obj).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printObj(result)
}

func replace(table string, id string, obj interface{}) {
	result, err := r.Table(table).Get(id).Replace(obj).RunWrite(session)
	r.Table(table).Get(id).Update(obj).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	printObj(result)
}
