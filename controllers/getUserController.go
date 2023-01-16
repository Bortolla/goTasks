package controllers

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"golang/database"
	// "golang/data"
)

// Nao implementado ainda
// Pegar um usuario baseado no seu nome e senha
func GetUserController(w http.ResponseWriter, r *http.Request) {
 
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



	database.ConnectDb("GetUser", "vitor", "123")
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}