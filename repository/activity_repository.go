package repository

import (
	"database/sql"
	"devcode/entity"
	"time"
)

type ActivityRepository struct {
	Db            *sql.DB
	lastID        int64
	activityCache []entity.Activity
}

type ActivityRepositoryInterface interface {
	Add(activity entity.Activity) (entity.Activity, error)
	GetAll() ([]entity.Activity, error)
	GetById(id int) (entity.ActivityWNull, error)
	DeleteById(id int) error
	UpdateById(id int, title string) (entity.ActivityWNull, error)
}

func (repo *ActivityRepository) UpdateById(id int, title string) (entity.ActivityWNull, error) {
	stmt, err := repo.Db.Prepare("UPDATE activity SET title = ? WHERE id = ?")
	if err != nil {
		return entity.ActivityWNull{}, err
	}
	_, err = stmt.Exec(title, id)

	// get updated data
	activity, err := repo.GetById(id)

	return activity, err
}

func (repo *ActivityRepository) DeleteById(id int) error {
	stmt, err := repo.Db.Prepare("DELETE FROM activity WHERE id = ?")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(id)
	if rowAffected, _ := result.RowsAffected(); rowAffected == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (repo *ActivityRepository) Add(activity entity.Activity) (entity.Activity, error) {
	repo.lastID++
	activity.ID = repo.lastID
	activity.CreatedAt = time.Now().Format(time.RFC3339)
	activity.UpdatedAt = activity.CreatedAt
	activity.DeletedAt = nil
	go func() {
		stmt, _ := repo.Db.Prepare("INSERT INTO activity (title, email) VALUES (?,?)")

		stmt.Exec(activity.Title, activity.Email)
		repo.activityCache = append(repo.activityCache, activity)
	}()

	return activity, nil
}

func (repo *ActivityRepository) GetAll() ([]entity.Activity, error) {
	if len(repo.activityCache) > 0 {
		return repo.activityCache, nil
	}
	//return repo.Db.Table("activity").Exec("INSERT INTO activity (title, email) VALUES (?,?)", activity.Title, activity.Email).Error
	stmt, err := repo.Db.Prepare("SELECT * FROM activity")
	if err != nil {
		return nil, err
	}
	results, err := stmt.Query()

	// iter golang select db scan
	var activities []entity.Activity
	for results.Next() {
		var activity entity.Activity
		err = results.Scan(&activity.ID, &activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return activities, err
}

func (repo *ActivityRepository) GetById(id int) (entity.ActivityWNull, error) {
	//return repo.Db.Table("activity").Exec("INSERT INTO activity (title, email) VALUES (?,?)", activity.Title, activity.Email).Error
	stmt, err := repo.Db.Prepare("SELECT * FROM activity WHERE id = ?")
	if err != nil {
		return entity.ActivityWNull{}, err
	}
	var activity entity.ActivityWNull
	err = stmt.QueryRow(id).Scan(&activity.ID, &activity.Email, &activity.Title, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
	return activity, err
}

func NewActivityRepository(db *sql.DB) ActivityRepositoryInterface {
	return &ActivityRepository{
		Db: db,
	}
}
