package main

import (
	"devcode/controller"
	"devcode/internal/database"
	"devcode/repository"
	"devcode/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//runtime.GOMAXPROCS(2)

	db := database.NewMysqlConnection()
	//sentry.CaptureMessage("Migrating Database")
	database.Migrate(db)
	//sentry.CaptureMessage("Database Migrated")

	app := fiber.New(fiber.Config{
		AppName: "Devcode Todo",
		Prefork: false,
	})
	//app.Use(logger.New())

	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(&activityRepository)
	activityController := controller.NewActivityController(&activityService)
	activityController.ActivityHttpRoute(app)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(&todoRepository)
	todoController := controller.NewTodoController(&todoService)
	todoController.TodoHttpRoute(app)

	log.Fatal(app.Listen(":3030"))
}
