package controller

import (
	"database/sql"
	"devcode/entity"
	"devcode/service"
	"github.com/gofiber/fiber/v2"
)

type ActivityController struct {
	activityService service.ActivityServiceInterface
}

func (ctrl *ActivityController) ActivityHttpRoute(app *fiber.App) {
	app.Post("/activity-groups", ctrl.Add)
	app.Get("/activity-groups-plain", ctrl.Plain)
	app.Get("/activity-groups", ctrl.GetAll)
	app.Get("/activity-groups/:id", ctrl.GetById)
	app.Delete("/activity-groups/:id", ctrl.Delete)
	app.Patch("/activity-groups/:id", ctrl.UpdateById)
}

func (ctrl *ActivityController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	err = ctrl.activityService.DeleteById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return ctx.Status(404).JSON(&entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Params("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
	}
	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    entity.EmptyObject{},
	})
}

func (ctrl *ActivityController) GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	activity, err := ctrl.activityService.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(404).JSON(&entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Params("id") + " Not Found",
				Data:    entity.EmptyObject{},
			})
		}
	}
	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (ctrl *ActivityController) GetAll(ctx *fiber.Ctx) error {
	activities, err := ctrl.activityService.GetAll()
	if err != nil {
		return ctx.Status(404).JSON(&entity.BaseApiResponse{
			Status:  "Not Found",
			Message: "Activities Not Found",
			Data:    entity.EmptyObject{},
		})
	}
	return ctx.Status(200).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}
func (ctrl *ActivityController) Plain(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("Hello")
}
func (ctrl *ActivityController) Add(ctx *fiber.Ctx) error {
	var requestData entity.ActivityCreateRequest
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

	insertedData, errInsert := ctrl.activityService.Add(requestData.Title, requestData.Email)
	if errInsert != nil {
		return ctx.Status(500).JSON(&entity.BaseApiResponse{
			Status:  "Internal Server Error",
			Message: "Internal Server Error",
			Data:    entity.EmptyObject{},
		})
	}

	return ctx.Status(201).JSON(&entity.BaseApiResponse{
		Status:  "Success",
		Message: "Success",
		Data:    insertedData,
	})

}

func (ctrl *ActivityController) UpdateById(ctx *fiber.Ctx) error {
	id, errParam := ctx.ParamsInt("id")
	if errParam != nil {
		return ctx.Status(400).JSON(&entity.BaseApiResponse{
			Status:  "Bad Request",
			Message: "Bad Request",
			Data:    entity.EmptyObject{},
		})
	}

	var requestData entity.ActivityCreateRequest
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

	updatedData, errUpdate := ctrl.activityService.UpdateById(id, requestData.Title)
	if errUpdate != nil {

		if errUpdate == sql.ErrNoRows {
			return ctx.Status(404).JSON(&entity.BaseApiResponse{
				Status:  "Not Found",
				Message: "Activity with ID " + ctx.Params("id") + " Not Found",
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
		Data:    updatedData,
	})

}

func NewActivityController(service *service.ActivityServiceInterface) *ActivityController {
	return &ActivityController{
		activityService: *service,
	}
}
