package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	muxRouter := http.NewServeMux()
	muxRouter.HandleFunc("/users", getUsers)
	fmt.Println("api in on :8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "John Doe",
		},
		{
			ID:   2,
			Name: "Ozzy",
		},
	})
}
