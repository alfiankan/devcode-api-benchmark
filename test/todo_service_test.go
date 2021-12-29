package test

import (
	"devcode/internal/database"
	"devcode/repository"
	"devcode/service"
	"log"
	"testing"
)

func TestTodoService_Add(t *testing.T) {
	db := database.NewMysqlConnection()
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(&todoRepository)
	todo, err := todoService.Add("test", 1)
	log.Println(err, todo)

}
