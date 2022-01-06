package controller

import (
	"database/sql"
	"devcode/entity"
	"devcode/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoEchoController struct {
	TodoService service.TodoServiceInterface
}

func (ctrl *TodoEchoController) TodoHttpEchoRoute(app *echo.Echo) {
	app.POST("/todo-items", ctrl.Add)
	app.GET("/todo-items", ctrl.GetAll)
	app.GET("/todo-items/:id", ctrl.GetById)
	app.DELETE("/todo-items/:id", ctrl.DeleteById)
	app.PATCH("/todo-items/:id", ctrl.UpdateById)
}

func (ctrl *TodoEchoController) UpdateById(ctx echo.Context) error {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
	}


	requestEdit := new(entity.TodoCreateRequest)
	if err := ctx.Bind(requestEdit); err != nil {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
	}


	result, errUpdateById := ctrl.TodoService.UpdateById(id, requestEdit.Title, strconv.FormatBool(requestEdit.IsActive))
	if errUpdateById != nil {
		if errUpdateById == sql.ErrNoRows {
			return ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
		return ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (ctrl *TodoEchoController) DeleteById(ctx echo.Context) error {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
	}
	errGetById := ctrl.TodoService.DeleteById(id)
	if errGetById != nil {
		if errGetById == sql.ErrNoRows {
			return ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
		return ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    entity.EmptyObject{},
	})
}

func (ctrl *TodoEchoController) GetById(ctx echo.Context) error {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
	}
	result, errGetById := ctrl.TodoService.GetById(id)
	if errGetById != nil {
		if errGetById == sql.ErrNoRows {
			return ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
		return ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (ctrl *TodoEchoController) GetAll(ctx echo.Context) error {
	var result []entity.Todo
	var err error
	groupId := ctx.QueryParam("activity_group_id")
	// if group id null
	if groupId != "" {
		groupIdInt, errGroupId := strconv.Atoi(groupId)
		if errGroupId != nil {
			return ctx.JSON(400, &entity.BaseApiResponse{
				Status:  "Bad Request",
				Message: "activity_group_id must be integer",
				Data:    entity.EmptyObject{},
			})
		}
		result, err = ctrl.TodoService.GetFilterAll(groupIdInt)
		if len(result) == 0 {
			return ctx.JSON(200, &entity.BaseApiResponse{
				Status:  "Success",
				Message: "Success",
				Data:    []entity.EmptyObject{},
			})
		}
		return ctx.JSON(200, &entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		})
	}
	// if without query
	result, err = ctrl.TodoService.GetAll()
	if err != nil {
		return ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}

	if len(result) == 0 {
		return ctx.JSON(200, &entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    []entity.EmptyObject{},
		})
	}

	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    nil,
	})
}

func (ctrl *TodoEchoController) Add(ctx echo.Context) error {

	requestData := new(entity.TodoCreateRequest)
	if err := ctx.Bind(requestData); err != nil {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
	}

	if requestData.Title == "" {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    entity.EmptyObject{},
		})
	}

	if requestData.ActivityGroupId == 0 {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "activity_group_id cannot be null",
			Data:    entity.EmptyObject{},
		})
	}

	insertedData, errInsert := ctrl.TodoService.Add(requestData.Title, requestData.ActivityGroupId)
	if errInsert != nil {
		return ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    errInsert.Error(),
		})
	}

	return ctx.JSON(201, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    insertedData,
	})

}

func NewTodoEchoController(service *service.TodoServiceInterface) *TodoEchoController {
	return &TodoEchoController{
		TodoService: *service,
	}
}
