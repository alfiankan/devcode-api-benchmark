package service

import (
	"devcode/entity"
	"devcode/repository"
)

type TodoService struct {
	TodoRepository repository.TodoRepositoryInterface
	todoCache      []entity.Todo
}

type TodoServiceInterface interface {
	Add(title string, activity_group_id int64) (entity.Todo, error)
	GetAll() ([]entity.Todo, error)
	GetById(id int) (entity.Todo, error)
	GetFilterAll(groupId int) ([]entity.Todo, error)
	DeleteById(id int) error
	UpdateById(id int, title string, isActive string) (entity.Todo, error)
}

func (repo *TodoService) UpdateById(id int, title string, isActive string) (entity.Todo, error) {
	return repo.TodoRepository.UpdateById(id, title, isActive)
}

func (repo *TodoService) DeleteById(id int) error {
	return repo.TodoRepository.DeleteById(id)
}

func (repo *TodoService) GetById(id int) (entity.Todo, error) {
	return repo.TodoRepository.GetById(id)
}

func (repo *TodoService) GetFilterAll(groupId int) ([]entity.Todo, error) {
	return repo.TodoRepository.GetFilterAll(groupId)
}

func (repo *TodoService) GetAll() ([]entity.Todo, error) {
	return repo.TodoRepository.GetAll()
}

func (repo *TodoService) Add(title string, activity_group_id int64) (entity.Todo, error) {

	return repo.TodoRepository.Add(entity.Todo{
		Title:           title,
		ActivityGroupId: activity_group_id,
		Priority:        "very-high",
		IsActive:        true,
	})
}

func NewTodoService(repo *repository.TodoRepositoryInterface) TodoServiceInterface {
	return &TodoService{
		TodoRepository: *repo,
	}
}
