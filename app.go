package main

import (
	"devcode/controller"
	"devcode/internal/database"
	"devcode/repository"
	"devcode/service"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/labstack/echo/v4"

	//"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func fiberServer() {
	runtime.GOMAXPROCS(2)

	db := database.NewMysqlConnection()
	//sentry.CaptureMessage("Migrating Database")
	database.Migrate(db)
	//sentry.CaptureMessage("Database Migrated")

	app := fiber.New(fiber.Config{
		AppName: "Devcode Todo",
		Prefork: true,
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	//app.Use(logger.New())
	app.Use(recover.New())

	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(&activityRepository)
	activityController := controller.NewActivityController(&activityService)
	activityController.ActivityHttpRoute(app)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(&todoRepository)
	todoController := controller.NewTodoController(&todoService)
	todoController.TodoHttpRoute(app)

	log.Println(app.Listen(":3030"))
}

func echoServer() {
	db := database.NewMysqlConnection()
	database.Migrate(db)
	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(&activityRepository)
	activityController := controller.NewActivityEchoController(&activityService)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(&todoRepository)
	todoController := controller.NewTodoEchoController(&todoService)

	e := echo.New()

	// Activity

	activityController.ActivityHttpEchoRoute(e)
	todoController.TodoHttpEchoRoute(e)

	e.Logger.Print(e.Start(":3030"))
}

func ginServer() {
	db := database.NewMysqlConnection()
	database.Migrate(db)
	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(&activityRepository)
	activityController := controller.NewActivityGinController(&activityService)

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(&todoRepository)
	todoController := controller.NewTodoGinController(&todoService)
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	activityController.ActivityHttpGinRoute(server)
	todoController.TodoHttpGinRoute(server)

	server.Run(":3030")

}

func main() {

    fiberServer()
    //echoServer()
    //ginServer()
}
