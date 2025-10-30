package main

import (
	"encoding/json"
	"net/http"
)

// Struct (equivalente a uma "entidade" em Java)
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Função que retorna uma lista de usuários em JSON
func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":8080", nil)
}
