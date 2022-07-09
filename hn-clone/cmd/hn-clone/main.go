package main

import (
	"database/sql"
	"example/internal"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "hackernews",
	}

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to db!")
	defer db.Close()

	userService := internal.NewUserService(db)

	/* router */

	r := mux.NewRouter()

	assets := http.FileServer(http.Dir("../../web/assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	subpages := http.FileServer(http.Dir("../../web/subpages/"))
	r.PathPrefix("/subpages/").Handler(http.StripPrefix("/subpages/", subpages))

	r.HandleFunc("/", internal.Top)
	r.HandleFunc("/news", internal.News)
	r.HandleFunc("/newest", internal.Newest)
	r.HandleFunc("/ask", internal.Ask)

	r.HandleFunc("/login", userService.Login).Methods("POST")
	r.HandleFunc("/user", userService.User).Methods("POST")

	addr := ":8080"
	fmt.Println("serving at", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
