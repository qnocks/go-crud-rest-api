package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: "1", Item: "Clean room", Completed: false},
	{Id: "2", Item: "Do homework", Completed: true},
	{Id: "3", Item: "Learn GO", Completed: false},
}

func getAllTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func saveTodo(ctx *gin.Context) {
	var todo Todo
	if err := ctx.BindJSON(&todo); err != nil {
		return
	}

	todos = append(todos, todo)
	ctx.IndentedJSON(http.StatusCreated, todo)
}

func getTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo, err = findTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo, err = findTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	ctx.IndentedJSON(http.StatusOK, todo)
}

func findTodoById(id string) (*Todo, error) {
	for i, todo := range todos {
		if todo.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()

	router.GET("/todos", getAllTodos)
	router.GET("/todos/:id", getTodoById)
	router.POST("/todos", saveTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)

	router.Run("localhost:8080")
}
