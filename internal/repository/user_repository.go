package repository

import (
	"database/sql"
	"github.com/Waelson/go-tests/internal/model"
)

const SELECT_ALL_USERS = "SELECT id, username, password, email FROM users"

type UserRepository interface {
	FindAll() ([]model.User, error)
	Save(user model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) FindAll() ([]model.User, error) {
	rows, err := r.db.Query(SELECT_ALL_USERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Save(user model.User) (model.User, error) {
	_, err := r.db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
