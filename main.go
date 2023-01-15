package main

import (
	"fmt"
	"errors"
	"net/http"
	"os"
	"golang/controllers"
	"golang/database"
)

func main() {

	database.ConnectDb()

	http.HandleFunc("/", controllers.GetRoot)

	http.HandleFunc("/user", controllers.GetUserController)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Servidor fechado")
	} else if err != nil {
		fmt.Println("Erro ao inicializar o servidor: %s\n", err)
		os.Exit(1)
	} 
}