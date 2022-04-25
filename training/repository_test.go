package training

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepo_GetById(t *testing.T) {
	// given
	repo, mock := createMockedRepo()
	defer repo.db.Close()
	expectedTraining := createTestTraining()
	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "distance", "time", "date"})
	rows.AddRow(
		expectedTraining.Id,
		expectedTraining.UserId,
		expectedTraining.Title,
		expectedTraining.Type,
		expectedTraining.Distance,
		expectedTraining.Time,
		expectedTraining.Date)
	mock.ExpectQuery("SELECT . FROM trainings WHERE id=").WithArgs(expectedTraining.Id).WillReturnRows(rows)

	// when
	result, err := repo.GetById(expectedTraining.Id)

	// then
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Equal(t, expectedTraining, result)
}

func TestRepo_GetByUserId(t *testing.T) {
	// given
	repo, mock := createMockedRepo()
	defer repo.db.Close()
	training := createTestTraining()
	expectedResult := []Training{}
	expectedResult = append(expectedResult, training)
	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "type", "distance", "time", "date"})
	rows.AddRow(
		training.Id,
		training.UserId,
		training.Title,
		training.Type,
		training.Distance,
		training.Time,
		training.Date)
	mock.ExpectQuery("SELECT . FROM trainings WHERE user_id=").WithArgs(training.UserId).WillReturnRows(rows)

	// when
	result, err := repo.GetByUserId(training.UserId)

	// then
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Equal(t, expectedResult, result)
}

func createMockedRepo() (Repo, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	return Repo{
		db: sqlxDB,
	}, mock
}

func createTestTraining() Training {
	date, _ := time.Parse(time.RFC3339, "2022-04-11T15:58:01.511Z")
	return Training{
		Id:       123,
		UserId:   "r32njo",
		Title:    "Test Run",
		Type:     "run",
		Distance: 2131,
		Time:     360,
		Date:     date,
	}
}
