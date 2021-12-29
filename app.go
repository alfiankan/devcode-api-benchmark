package main

import (
	"devcode/controller"
	"devcode/internal/database"
	"devcode/repository"
	"devcode/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	//err := sentry.Init(sentry.ClientOptions{
	//	// Either set your DSN here or set the SENTRY_DSN environment variable.
	//	Dsn: "https://5ff9dbeada024f7baf2fe358c041da62@o913414.ingest.sentry.io/6125852",
	//	// Either set environment and release here or set the SENTRY_ENVIRONMENT
	//	// and SENTRY_RELEASE environment variables.
	//	Environment: "",
	//	Release:     "my-project-name@1.0.0",
	//	// Enable printing of SDK debug messages.
	//	// Useful when getting started or trying to figure something out.
	//	Debug: false,
	//})
	//if err != nil {
	//	log.Fatalf("sentry.Init: %s", err)
	//}

	db := database.NewMysqlConnection()
	//sentry.CaptureMessage("Migrating Database")
	database.Migrate(db)
	//sentry.CaptureMessage("Database Migrated")

	app := fiber.New(fiber.Config{
		AppName: "Devcode Todo",
		Prefork: true,
	})

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
