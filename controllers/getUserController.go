package controllers

import (
	"fmt"
	"io"
	"net/http"
	"golang/database"
	// "golang/data"
)

func GetUserController(w http.ResponseWriter, r *http.Request) {
	database.ConnectDb("GetUser")
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}