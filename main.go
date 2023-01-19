package main

import (
	"fmt"
	"errors"
	"net/http"
	"os"
	"golang/controllers"
)

func main() {

	http.HandleFunc("/register", controllers.RegisterController)
	
	http.HandleFunc("/login", controllers.LoginController)

	http.HandleFunc("/createTask", controllers.CreateTaskController)

	// http.HandleFunc("/getTasks", controllers.GetTasksController)

	// http.HandleFunc("/deleteTask", controllers.DeleteTaskController)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Servidor fechado")
	} else if err != nil {
		fmt.Println("Erro ao inicializar o servidor: %s\n", err)
		os.Exit(1)
	} 
}