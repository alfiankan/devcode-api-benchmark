package repository

import (
	"database/sql"
	"devcode/entity"
	"log"
	"time"
)

type TodoRepository struct {
	Db        *sql.DB
	lastID    int64
	todoCache []entity.Todo
}

type TodoRepositoryInterface interface {
	Add(todo entity.Todo) (entity.Todo, error)
	GetAll() ([]entity.Todo, error)
	GetFilterAll(groupId int) ([]entity.Todo, error)
	GetById(id int) (entity.Todo, error)
	DeleteById(id int) error
	UpdateById(id int, title string, isActive string) (entity.Todo, error)
}

func (repo *TodoRepository) GetFilterAll(groupId int) ([]entity.Todo, error) {
	stmt, err := repo.Db.Prepare("SELECT * FROM todos WHERE activity_group_id = ?")
	if err != nil {
		return nil, err
	}
	results, err := stmt.Query(groupId)

	// iter golang select db scan
	var todos []entity.Todo
	for results.Next() {
		var todo entity.Todo
		err = results.Scan(&todo.ID, &todo.Title, &todo.ActivityGroupId, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func (repo *TodoRepository) Add(todo entity.Todo) (entity.Todo, error) {
	repo.lastID++
	todo.ID = repo.lastID
	todo.CreatedAt = time.Now().Format(time.RFC3339)
	todo.UpdatedAt = todo.CreatedAt
	todo.DeletedAt = nil
	go func() {
		stmt, err1 := repo.Db.Prepare("INSERT INTO todos (title, activity_group_id, is_active, priority) VALUES (?,?,?,?)")
		if err1 != nil {
			log.Println(err1)
		}
		_, err := stmt.Exec(todo.Title, todo.ActivityGroupId, todo.IsActive, todo.Priority)
		if err != nil {
			log.Println(err)
		}
		repo.todoCache = append(repo.todoCache, todo)
	}()

	return todo, nil

}

func (repo *TodoRepository) GetAll() ([]entity.Todo, error) {
	if len(repo.todoCache) > 0 {
		return repo.todoCache, nil
	}
	stmt, err := repo.Db.Prepare("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	results, err := stmt.Query()

	// iter golang select db scan
	var todos []entity.Todo
	for results.Next() {
		var todo entity.Todo
		err = results.Scan(&todo.ID, &todo.Title, &todo.ActivityGroupId, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}
	results.Close()

	return todos, err
}

func (repo *TodoRepository) GetById(id int) (entity.Todo, error) {
	stmt, err := repo.Db.Prepare("SELECT * FROM todos WHERE id = ?")
	if err != nil {
		return entity.Todo{}, err
	}
	var todo entity.Todo
	err = stmt.QueryRow(id).Scan(&todo.ID, &todo.Title, &todo.ActivityGroupId, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
	return todo, err
}

func (repo *TodoRepository) DeleteById(id int) error {
	stmt, err := repo.Db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(id)
	if rowAffected, _ := result.RowsAffected(); rowAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (repo *TodoRepository) UpdateById(id int, title string, isActive string) (entity.Todo, error) {
	var err error
	if title != "" {
		stmt, err := repo.Db.Prepare("UPDATE todos SET title = ? WHERE id = ?")
		if err != nil {
			log.Println(err)
			return entity.Todo{}, err
		}
		_, err = stmt.Exec(title, id)
	}
	if isActive != "" {
		isTodoActive := 0
		if isActive == "true" {
			isTodoActive = 1
		}
		stmt, err := repo.Db.Prepare("UPDATE todos SET is_active = ? WHERE id = ?")
		if err != nil {
			log.Println(err)
			return entity.Todo{}, err
		}
		_, err = stmt.Exec(isTodoActive, id)
	}

	// get updated data
	activity, err := repo.GetById(id)

	return activity, err
}

func NewTodoRepository(db *sql.DB) TodoRepositoryInterface {
	return &TodoRepository{
		Db: db,
	}
}
