package controllers

import (
	"fmt"
	"io"
	"net/http"
	"golang/database"
	// "golang/data"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	// Pegar do corpo da requisição: 
	// Nome e senha do usuário

	database.ConnectDb("GetUser")
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}