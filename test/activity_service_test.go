package test

import (
	"devcode/internal/database"
	"devcode/repository"
	"devcode/service"
	"log"
	"testing"
	"time"

)

func TestActivityService_AddSingle(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)
	serv := service.NewActivityService(&repo)

	start := time.Now()
	res, err := serv.Add("hello","hello")
	log.Println(err, res)
	log.Println("GET TIME ELAPSED",time.Since(start).Milliseconds(),"ms")
	time.Sleep(time.Second * 5)
}

func TestGetDetail(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)

	start := time.Now()
	res, err := repo.UpdateById(2, "upadfafafafdte")
	log.Println(err, res)
	log.Println("UPDATE TIME ELAPSED",time.Since(start).Milliseconds(),"ms")
	time.Sleep(time.Second * 5)
}
