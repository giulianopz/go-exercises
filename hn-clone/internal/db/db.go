package db

import (
	"database/sql"
	"example/internal/model"
	"fmt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) Create(username string, password string) int64 {

	result, err := ur.db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		fmt.Println("insert failed:", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("insert failed:", err)
	}
	return id
}

func (ur *UserRepo) UserByHash(hash string) (model.User, error) {

	var user model.User
	row := ur.db.QueryRow("SELECT * FROM user WHERE password = ?", hash)

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("Cannot find user by hash: %s", hash)
		}
		return user, fmt.Errorf("Cannot find user by hash: %s", hash)
	}
	return user, nil
}
