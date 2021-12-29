package service

import (
	"devcode/entity"
	"devcode/repository"
)

type ActivityService struct {
	activityRepository repository.ActivityRepositoryInterface
}

type ActivityServiceInterface interface {
	Add(title, email string) (entity.Activity, error)
	GetAll() ([]entity.Activity, error)
	GetById(id int) (entity.ActivityWNull, error)
	DeleteById(id int) error
	UpdateById(id int, title string) (entity.ActivityWNull, error)
}

func (repo *ActivityService) UpdateById(id int, title string) (entity.ActivityWNull, error) {
	return repo.activityRepository.UpdateById(id, title)
}
func (repo *ActivityService) DeleteById(id int) error {
	return repo.activityRepository.DeleteById(id)
}

func (repo *ActivityService) Add(title, email string) (entity.Activity, error) {
	return repo.activityRepository.Add(entity.Activity{
		Title: title,
		Email: email,
	})
}

func (repo *ActivityService) GetAll() ([]entity.Activity, error) {
	return repo.activityRepository.GetAll()
}

func (repo *ActivityService) GetById(id int) (entity.ActivityWNull, error) {
	return repo.activityRepository.GetById(id)
}

func NewActivityService(repo *repository.ActivityRepositoryInterface) ActivityServiceInterface {
	return &ActivityService{
		activityRepository: *repo,
	}
}
