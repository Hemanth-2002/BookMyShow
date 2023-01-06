package database

import (
	"bms/model"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func SetupDBClient() (*gorm.DB, sqlmock.Sqlmock, error) {
	var (
		Db  *sql.DB
		err error
	)

	Db, mock, _ := sqlmock.New()

	db, err := gorm.Open("postgres", Db)
	return db, mock, err
}
func TestCreateUser(t *testing.T) {

	db, mock, err := SetupDBClient()
	if err != nil {
		t.Fatalf("failed to create mock db client: %v", err)
	}
	defer db.Close()
	mockClient := DBClient{
		Db: db,
	}
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" (.+)`).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1),
	).WillReturnError(err)
	mock.ExpectCommit()
	id, err := mockClient.CreateUser(model.User{})

	if id != 1 || err != nil {
		t.Errorf("failed to create user")
	}

}
