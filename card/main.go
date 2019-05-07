package main

import (
	"card/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Index)

	http.HandleFunc("/room", handlers.Rooms)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
