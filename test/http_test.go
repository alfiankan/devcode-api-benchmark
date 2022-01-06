package test

import (
	"devcode/entity"
	"devcode/internal/database"
	"devcode/repository"
	"devcode/service"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/labstack/echo/v4"
)

func TestAddActivityHttp(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)
	serv := service.NewActivityService(&repo)
	log.Println(serv)
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/activity-groups", func(ctx *fiber.Ctx) error {
		var requestData entity.ActivityCreateRequest
		ctx.BodyParser(&requestData)

		start := time.Now()
		serv.Add(requestData.Title, requestData.Email)
		log.Println("Elapsed : ", time.Since(start))

		return ctx.Status(201).JSON(&entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    requestData,
		})

	})


	app.Listen(":3030")

	

}


func TestEcho(t *testing.T) {
	db := database.NewMysqlConnection()
	repo := repository.NewActivityRepository(db)
	serv := service.NewActivityService(&repo)
	e := echo.New()
	e.POST("/activity-groups", func(c echo.Context) error {

		serv.Add("ddd","dddd")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3030"))
}