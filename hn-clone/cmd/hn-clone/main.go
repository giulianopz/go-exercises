package main

import (
	"example/internal"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("../../web/assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", internal.Home)

	r.HandleFunc("/news", internal.News)

	addr := ":8080"
	fmt.Println("serving at", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
