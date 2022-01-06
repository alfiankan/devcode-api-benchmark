package controller

import (
	"database/sql"
	"devcode/entity"
	"devcode/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoGinController struct {
	TodoService service.TodoServiceInterface
}

func (ctrl *TodoGinController) TodoHttpGinRoute(app *gin.Engine) {
	app.POST("/todo-items", ctrl.Add)
	app.GET("/todo-items", ctrl.GetAll)
	app.GET("/todo-items/:id", ctrl.GetById)
	app.DELETE("/todo-items/:id", ctrl.DeleteById)
	app.PATCH("/todo-items/:id", ctrl.UpdateById)
}

func (ctrl *TodoGinController) UpdateById(ctx *gin.Context) {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
		return
	}

	requestEdit := new(entity.TodoCreateRequest)
	if err := ctx.Bind(requestEdit); err != nil {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
		return
	}

	result, errUpdateById := ctrl.TodoService.UpdateById(id, requestEdit.Title, strconv.FormatBool(requestEdit.IsActive))
	if errUpdateById != nil {
		if errUpdateById == sql.ErrNoRows {
			ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
			return
		}
		ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
		return
	}
	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (ctrl *TodoGinController) DeleteById(ctx *gin.Context) {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
		return
	}
	errGetById := ctrl.TodoService.DeleteById(id)
	if errGetById != nil {
		if errGetById == sql.ErrNoRows {
			ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
			return
		}
		ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
		return
	}
	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    entity.EmptyObject{},
	})
}

func (ctrl *TodoGinController) GetById(ctx *gin.Context) {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "id must be integer",
			Data:    entity.EmptyObject{},
		})
		return
	}
	result, errGetById := ctrl.TodoService.GetById(id)
	if errGetById != nil {
		if errGetById == sql.ErrNoRows {
			ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Todo with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
			return
		}
		ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
		return
	}
	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    result,
	})
}

func (ctrl *TodoGinController) GetAll(ctx *gin.Context) {
	var result []entity.Todo
	var err error
	groupId := ctx.Query("activity_group_id")
	// if group id null
	if groupId != "" {
		groupIdInt, errGroupId := strconv.Atoi(groupId)
		if errGroupId != nil {
			ctx.JSON(400, &entity.BaseApiResponse{
				Status:  "Bad Request",
				Message: "activity_group_id must be integer",
				Data:    entity.EmptyObject{},
			})
			return
		}
		result, err = ctrl.TodoService.GetFilterAll(groupIdInt)
		if len(result) == 0 {
			ctx.JSON(200, &entity.BaseApiResponse{
				Status:  "Success",
				Message: "Success",
				Data:    []entity.EmptyObject{},
			})
			return
		}
		ctx.JSON(200, &entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    result,
		})
		return
	}
	// if without query
	result, err = ctrl.TodoService.GetAll()
	if err != nil {
		ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
		return
	}

	if len(result) == 0 {
		ctx.JSON(200, &entity.BaseApiResponse{
			Status:  "Success",
			Message: "Success",
			Data:    []entity.EmptyObject{},
		})
		return
	}

	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    nil,
	})
}

func (ctrl *TodoGinController) Add(ctx *gin.Context) {

	requestData := new(entity.TodoCreateRequest)
	if err := ctx.Bind(requestData); err != nil {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
		return
	}

	if requestData.Title == "" {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    entity.EmptyObject{},
		})
		return
	}

	if requestData.ActivityGroupId == 0 {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "activity_group_id cannot be null",
			Data:    entity.EmptyObject{},
		})
		return
	}

	insertedData, errInsert := ctrl.TodoService.Add(requestData.Title, requestData.ActivityGroupId)
	if errInsert != nil {
		ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    errInsert.Error(),
		})
		return
	}

	ctx.JSON(201, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    insertedData,
	})

}

func NewTodoGinController(service *service.TodoServiceInterface) *TodoGinController {
	return &TodoGinController{
		TodoService: *service,
	}
}
