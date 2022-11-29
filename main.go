package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

var todos = []todo{
	{ID: "1", Title: "todo 1", IsDone: false},
	{ID: "2", Title: "todo 2", IsDone: true},
	{ID: "3", Title: "todo 3", IsDone: false},
	{ID: "4", Title: "todo 4", IsDone: false},
}

func getTodos(ctx *gin.Context) {
	fmt.Println("123")
	ctx.IndentedJSON(http.StatusOK, todos)
}

func createNewTodo(ctx *gin.Context) {
	var new_todo todo

	if err := ctx.BindJSON(&new_todo); err != nil {
		return
	}
	todos = append(todos, new_todo)
	ctx.IndentedJSON(http.StatusCreated, new_todo)
}

func getTodoById(id string) (*todo, error) {
	fmt.Println(id)
	for i, item := range todos {
		if item.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found !")
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todo", createNewTodo)
	router.GET("/todo/:id", getTodo)
	router.Run("localhost:" + util.GodotEnv("GO_PORT"))
}
