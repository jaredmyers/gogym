// Golang
// Server using Gin

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Music", Completed: false},
}

func getToDos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)
}

func addToDo(context *gin.Context) {

	var newToDo todo
	if err := context.BindJSON(&newToDo); err != nil {
		return
	}

	todos = append(todos, newToDo)
	context.IndentedJSON(http.StatusCreated, newToDo)
}

func getToDo(context *gin.Context) {

	id := context.Param("id")
	todo, err := getById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}
func getById(id string) (*todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Not Found")
}

func toggleStatus(context *gin.Context) {

	id := context.Param("id")
	todo, err := getById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func notnotmain() {

	router := gin.Default()
	router.GET("/todos", getToDos)
	router.GET("/todos/:id", getToDo)
	router.PATCH("/todos/:id", toggleStatus)
	router.POST("/todos", addToDo)
	router.Run("localhost:8000")
}
