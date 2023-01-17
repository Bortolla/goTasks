package controllers

import (
	"fmt"
	// "io"
	"encoding/json"
	"net/http"
	// "golang/database"
	"golang/data"
	"golang/utils"
)

type Person struct {
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

  if len(p.Nome) == 0 || len(p.Senha) == 0 {
		// Adicionar JSON com mensagem
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		// jsonData := []byte(`{"msg":"erro", "status": "401"}`)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "401",
		})
		return
	} else {

		if utils.FindUser(p.Nome, p.Senha) == false {
			data.RegisterData(p.Nome, p.Senha)
			fmt.Println("usuario cadastrado")
		} else {
			fmt.Println("usuario já cadastrado")
			return
		}
	}
}