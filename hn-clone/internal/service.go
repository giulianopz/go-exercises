package internal

import (
	"database/sql"
	"encoding/json"
	"example/internal/db"
	"example/internal/mapper"
	"example/internal/util"
	"fmt"
	"net/http"
)

type UserService struct {
	userRepo *db.UserRepo
}

func NewUserService(database *sql.DB) *UserService {

	return &UserService{userRepo: db.NewUserRepo(database)}
}

func (s *UserService) User(w http.ResponseWriter, r *http.Request) {

	usr := r.FormValue("username")
	pwd := r.FormValue("password")
	hashed, _ := util.HashPassword(pwd)
	userId := s.userRepo.Create(usr, hashed)
	bytes, _ := json.Marshal(userId)

	//TODO redirect logged-in user to main page within a new session
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {

	pwd := r.FormValue("password")
	hash, err := util.HashPassword(pwd)
	if err != nil {
		mapper.ErrToISE(w, err)
	}

	user, err := s.userRepo.UserByHash(hash)
	//TODO if user is present redirect logged-in user to home
	fmt.Println(user)
}
