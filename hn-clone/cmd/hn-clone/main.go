package main

import (
	"example/internal"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("../../web/assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", internal.Home)

	//http.HandleFunc("/news", internal.News)

	addr := ":8080"
	fmt.Println("serving at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
