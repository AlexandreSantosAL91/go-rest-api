package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("Iniciando o servidor Rest com Go")

	log.Fatal(http.ListenAndServe(":8080", nil))
	/*
		Endereço do localhost: http://localhost:8080/users
	*/
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "João",
		},

		{
			ID:   1,
			Name: "Carlos",
		},
		{
			ID:   3,
			Name: "Davi",
		},
	})
}
