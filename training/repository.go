package training

import "github.com/jmoiron/sqlx"

type Repo struct {
	db *sqlx.DB
}

type Repository interface {
	GetById(id uint) (Training, error)
	GetByUserId(userId string) ([]Training, error)
	Create(training Training) error
}

func (r Repo) GetById(id uint) (Training, error) {
	training := Training{}
	err := r.db.Get(&training, "SELECT * FROM trainings WHERE id=$1", id)

	return training, err
}

func (r Repo) GetByUserId(userId string) ([]Training, error) {
	var trainings []Training
	err := r.db.Select(&trainings, "SELECT * FROM trainings WHERE user_id=$1", userId)

	return trainings, err
}

func (r Repo) Create(training Training) error {
	_, err := r.db.NamedExec("INSERT INTO trainings (user_id, title, type, distance, time, date) "+
		"VALUES (:user_id, :title, :type, :distance, :time, :date)", training)
	if err != nil {
		return err
	}
	return nil
}

func CreateRepository(db *sqlx.DB) Repository {
	return Repo{db: db}
}
