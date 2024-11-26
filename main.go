package main

import (
	"fmt"
	"net/http"
	"simpledbservice/db"
	"simpledbservice/handlers"
)

func main() {
	database := db.New()
	database.New() // Initialize the database connection
	handle := handlers.New(database)
	http.HandleFunc("/", handle.HandleGetAll)
	http.HandleFunc("/get", handle.HandleGetOne)
	http.HandleFunc("/create", handle.HandlePost)
	http.HandleFunc("/update", handle.HandleUpdate)
	http.HandleFunc("/delete", handle.HandleDelete)
	fmt.Println("Server started")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
