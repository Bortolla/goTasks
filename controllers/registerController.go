package controllers

import (
	"fmt"
	"io"
	"net/http"
	"golang/database"
	// "golang/data"
)

//  nao implementado
func RegisterController(w http.ResponseWriter, r *http.Request) {
	// Pegar do corpo da requisição: 
	// Nome e senha do usuario

	database.ConnectDb("RegisterUser") // ? 
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}