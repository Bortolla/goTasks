package controllers

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"golang/database"
	// "golang/data"
)

type Person struct {
	Nome string
	Senha  string
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	// Pegar do corpo da requisição: 
	// 

	if r.Method != "POST" {
		http.Error(w, "Método não suportado.", http.StatusNotFound)
    return
	}

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // Do something with the Person struct...
	w.WriteHeader(200)
  fmt.Fprintf(w, "Person: %+v", p.Nome)
	// fmt.Fprintf(w, "Person: %+v", p)

	database.ConnectDb("GetUser")
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}