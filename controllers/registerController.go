package controllers

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	// "golang/database"
	// "golang/data"
	"golang/utils"
)

type Person struct {
	// ID int
	Nome string
	Senha  string
}

func RegisterController(w http.ResponseWriter, r *http.Request) {
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

  if len(p.Nome) <= 0 || len(p.Senha) <= 0 {
		// Adicionar JSON com mensagem
		w.WriteHeader(400)
		io.WriteString(w, "Nome e/ou Senha mt curtos")
		return
	} 

	fmt.Println(p.Nome, p.Senha)
	fmt.Println(utils.FindUser(p.Nome, p.Senha))

	// if utils.FindUser(p.Nome, p.Senha) == false {
	// 	data.RegisterData(p.Nome, p.Senha)
	// } else {
	// 	return
	// }

	// Verificar se o usuario ja existe | eu poderia ate chamar a propria controller?
	io.WriteString(w, "prosseguir com o cadastro")


	// w.WriteHeader(200)
	// fmt.Fprintf(w, "Person: %+v", p.Nome)
	// fmt.Fprintf(w, "Person: %+v", p)

	// database.ConnectDb("RegisterUser") // ? 
	// fmt.Printf("got /user request\n")
	// io.WriteString(w, "Hello, user!\n")
}