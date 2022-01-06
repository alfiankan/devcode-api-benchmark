package controller

import (
	"database/sql"
	"devcode/entity"
	"devcode/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityGinController struct {
	activityService service.ActivityServiceInterface
}

func (ctrl *ActivityGinController) ActivityHttpGinRoute(app *gin.Engine) {
	app.POST("/activity-groups", ctrl.AddNew)
	app.GET("/activity-groups", ctrl.GetAll)
	app.GET("/activity-groups/:id", ctrl.GetById)
	app.DELETE("/activity-groups/:id", ctrl.Delete)
	app.PATCH("/activity-groups/:id", ctrl.UpdateById)
}

func (ctrl *ActivityGinController) UpdateById(ctx *gin.Context) {
	id, errParam := strconv.Atoi(ctx.Param("id"))
	if errParam != nil {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
		return
	}

	requestData := new(entity.ActivityCreateRequest)
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

	updatedData, errUpdate := ctrl.activityService.UpdateById(id, requestData.Title)
	if errUpdate != nil {

		if errUpdate == sql.ErrNoRows {
			ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Param("id") + " Not Found",
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
		Data:    updatedData,
	})
}

func (ctrl *ActivityGinController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	err = ctrl.activityService.DeleteById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
			return
		}
	}
	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    entity.EmptyObject{},
	})
}

func (ctrl *ActivityGinController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	activity, err := ctrl.activityService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(404, &entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Param("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
			return
		}
	}
	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (ctrl *ActivityGinController) GetAll(ctx *gin.Context) {
	activities, err := ctrl.activityService.GetAll()
	if err != nil {
		ctx.JSON(404, &entity.BaseApiResponse{
			Status:  "Not Found",
			Message: "Activities Not Found",
			Data:    entity.EmptyObject{},
		})
		return
	}
	ctx.JSON(200, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

func (ctrl *ActivityGinController) AddNew(ctx *gin.Context) {
	requestData := new(entity.ActivityCreateRequest)
	if err := ctx.Bind(requestData); err != nil {
		log.Println(err)
	}

	if requestData.Title == "" {
		ctx.JSON(400, &entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "title cannot be null",
			Data:    entity.EmptyObject{},
		})
		return
	}

	insertedData, errInsert := ctrl.activityService.Add(requestData.Title, requestData.Email)
	if errInsert != nil {
		ctx.JSON(500, &entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
		return
	}

	ctx.JSON(201, &entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    insertedData,
	})

}

func NewActivityGinController(service *service.ActivityServiceInterface) *ActivityGinController {
	return &ActivityGinController{
		activityService: *service,
	}
}
