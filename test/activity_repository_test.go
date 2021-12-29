package test

import (
	"devcode/entity"
	"devcode/internal/database"
	"devcode/repository"
	"log"
	"sync"
	"testing"
	"time"
)

func TestActivityRepository_Add(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			start := time.Now()
			repo.Add(entity.Activity{
				Title: "test repo",
				Email: "em@email.com",
			})
			log.Println(time.Since(start))
			wg.Done()
		}()

	}
	wg.Wait()
}

func TestActivityRepository_AddSingle(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)

	start := time.Now()
	res, err := repo.Add(entity.Activity{
		Title: "test repo",
		Email: "em@email.com",
	})
	log.Println(err, res)
	log.Println(time.Since(start))

}

func TestGetAll(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)

	start := time.Now()
	result, err := repo.GetAll()
	log.Println(err)
	for _, v := range result {
		log.Println(v)
	}
		
	
	log.Println(time.Since(start))
}

func TestGetAById(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)

	start := time.Now()
	result, err := repo.GetById(5)
	log.Println(err)
	log.Println(result)
	log.Println(time.Since(start))
}

func TestDeleteById(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)

	start := time.Now()
	err := repo.DeleteById(5)
	log.Println(err)
	log.Println(time.Since(start))
}