package test

import (
	"devcode/entity"
	"devcode/internal/database"
	"devcode/repository"
	"log"
	"time"

	"testing"
)

func TestTodoRepository_Add(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewTodoRepository(db)

	start := time.Now()
	res, err := repo.Add(entity.Todo{
		Title:           "test repo",
		ActivityGroupId: 1,
	})
	log.Println(err, res)
	log.Println(time.Since(start))
}

func TestTodoRepository_GetAll(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewTodoRepository(db)

	start := time.Now()
	res, err := repo.GetAll()
	log.Println(err, res)
	log.Println(time.Since(start))
}
