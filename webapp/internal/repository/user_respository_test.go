package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Waelson/go-tests/internal/model"
	"github.com/Waelson/go-tests/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAll_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do DB: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).AddRow(1, "waelson", "123456", "test@test.com")
	mock.ExpectQuery(repository.SELECT_ALL_USERS).WillReturnRows(rows)

	repo := repository.NewUserRepository(db)
	users, err := repo.FindAll()

	assert.Equal(t, 1, len(users))
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindAll_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do DB: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(repository.SELECT_ALL_USERS).WillReturnError(assert.AnError)

	repo := repository.NewUserRepository(db)
	users, err := repo.FindAll()

	assert.Nil(t, users)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestSave_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do DB: %v", err)
	}
	defer db.Close()

	user := model.User{Username: "waelson", Password: "123456", Email: "test@test.com"}
	mock.ExpectExec("INSERT INTO users").WithArgs(user.Username, user.Password, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewUserRepository(db)
	savedUser, err := repo.Save(user)

	assert.Equal(t, user, savedUser)
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestSave_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock do DB: %v", err)
	}
	defer db.Close()

	user := model.User{Username: "waelson", Password: "123456", Email: "test@test.com"}
	mock.ExpectExec("INSERT INTO users").WithArgs(user.Username, user.Password, user.Email).WillReturnError(assert.AnError)

	repo := repository.NewUserRepository(db)
	savedUser, err := repo.Save(user)

	assert.Equal(t, model.User{}, savedUser)
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
