package controller

import (
	"database/sql"
	"devcode/entity"
	"devcode/service"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ActivityEchoController struct {
	activityService service.ActivityServiceInterface
}

func (ctrl *ActivityEchoController) ActivityHttpEchoRoute(app *echo.Echo) {
	app.POST("/activity-groups", ctrl.AddNew)
	app.GET("/activity-groups", ctrl.GetAll)
	app.GET("/activity-groups/:id", ctrl.GetById)
	app.DELETE("/activity-groups/:id", ctrl.Delete)
	app.PATCH("/activity-groups/:id", ctrl.UpdateById)
}

func (ctrl *ActivityEchoController) UpdateById(ctx echo.Context) error {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
	}

	requestData := new(entity.ActivityCreateRequest)
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

	updatedData, errUpdate := ctrl.activityService.UpdateById(id, requestData.Title)
	if errUpdate != nil {

		if errUpdate == sql.ErrNoRows {
			return ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Param("id") + " Not Found",
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
		Data:    updatedData,
	})
}

func (ctrl *ActivityEchoController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	err = ctrl.activityService.DeleteById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
	}
	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    entity.EmptyObject{},
	})
}

func (ctrl *ActivityEchoController) GetById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	activity, err := ctrl.activityService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
	}
	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (ctrl *ActivityEchoController) GetAll(ctx echo.Context) error {
	activities, err := ctrl.activityService.GetAll()
	if err != nil {
		return ctx.JSON(404, &entity.BaseApiResponse{
			Status:  "Not Found",
			Message: "Activities Not Found",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

func (ctrl *ActivityEchoController) AddNew(ctx echo.Context) error {
	requestData := new(entity.ActivityCreateRequest)
	if err := ctx.Bind(requestData); err != nil {
		log.Println(err)
	}

	if requestData.Title == "" {
		return ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    entity.EmptyObject{},
		})
	}

	insertedData, errInsert := ctrl.activityService.Add(requestData.Title, requestData.Email)
	if errInsert != nil {
		return ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}

	return ctx.JSON(201, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    insertedData,
	})

}

func NewActivityEchoController(service *service.ActivityServiceInterface) *ActivityEchoController {
	return &ActivityEchoController{
		activityService: *service,
	}
}
