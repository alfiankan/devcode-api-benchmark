package controller

import (
	"database/sql"
	"devcode/entity"
	"devcode/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	TodoService service.TodoServiceInterface
}

func (ctrl *TodoController) TodoHttpRoute(app *fiber.App) {
	app.Post("/todo-items", ctrl.Add)
	app.Get("/todo-items", ctrl.GetAll)
	app.Get("/todo-items/:id", ctrl.GetById)
	app.Delete("/todo-items/:id", ctrl.DeleteById)
	app.Patch("/todo-items/:id", ctrl.UpdateById)
}

func (ctrl *TodoController) UpdateById(ctx *fiber.Ctx) error {
	id, errParam := ctx.ParamsInt("id")
	if errParam != nil {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
	}

	var requestEdit entity.TodoCreateRequest

	errParseBody := ctx.BodyParser(&requestEdit)
	if errParseBody != nil {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
	}

	result, errUpdateById := ctrl.TodoService.UpdateById(id, requestEdit.Title, strconv.FormatBool(requestEdit.IsActive))
	if errUpdateById != nil {
		if errUpdateById == sql.ErrNoRows {
			return ctx.Status(404).JSON(&entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Params("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
		return ctx.Status(500).JSON(&entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (ctrl *TodoController) DeleteById(ctx *fiber.Ctx) error {
	id, errParam := ctx.ParamsInt("id")
	if errParam != nil {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
	}
	errGetById := ctrl.TodoService.DeleteById(id)
	if errGetById != nil {
		if errGetById == sql.ErrNoRows {
			return ctx.Status(404).JSON(&entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Params("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
		return ctx.Status(500).JSON(&entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    entity.EmptyObject{},
	})
}

func (ctrl *TodoController) GetById(ctx *fiber.Ctx) error {
	id, errParam := ctx.ParamsInt("id")
	if errParam != nil {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
	}
	result, errGetById := ctrl.TodoService.GetById(id)
	if errGetById != nil {
		if errGetById == sql.ErrNoRows {
			return ctx.Status(404).JSON(&entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Params("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
		return ctx.Status(500).JSON(&entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (ctrl *TodoController) GetAll(ctx *fiber.Ctx) error {
	var result []entity.Todo
	var err error
	groupId := ctx.Query("activity_group_id", "#")
	// if group id null
	if groupId != "#" {
		groupIdInt, errGroupId := strconv.Atoi(groupId)
		if errGroupId != nil {
			return ctx.Status(400).JSON(&entity.BaseApiResponse{
				Status:  "Bad Request",
				Message: "activity_group_id must be integer",
				Data:    entity.EmptyObject{},
			})
		}
		result, err = ctrl.TodoService.GetFilterAll(groupIdInt)
		if len(result) == 0 {
			return ctx.Status(200).JSON(&entity.BaseApiResponse{
				Status:  "Success",
				Message: "Success",
				Data:    []entity.EmptyObject{},
			})
		}
		return ctx.Status(200).JSON(&entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		})
	}
	// if without query
	result, err = ctrl.TodoService.GetAll()
	if err != nil {
		return ctx.Status(500).JSON(&entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}

	if len(result) == 0 {
		return ctx.Status(200).JSON(&entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    []entity.EmptyObject{},
		})
	}

	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    nil,
	})
}

func (ctrl *TodoController) Add(ctx *fiber.Ctx) error {
	var requestData entity.TodoCreateRequest
	errParseBody := ctx.BodyParser(&requestData)

	if errParseBody != nil {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
	}

	if requestData.Title == "" {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    entity.EmptyObject{},
		})
	}

	if requestData.ActivityGroupId == 0 {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "activity_group_id cannot be null",
			Data:    entity.EmptyObject{},
		})
	}

	insertedData, errInsert := ctrl.TodoService.Add(requestData.Title, requestData.ActivityGroupId)
	if errInsert != nil {
		return ctx.Status(500).JSON(&entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    errInsert.Error(),
		})
	}

	return ctx.Status(201).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    insertedData,
	})

}

func NewTodoController(service *service.TodoServiceInterface) *TodoController {
	return &TodoController{
		TodoService: *service,
	}
}
