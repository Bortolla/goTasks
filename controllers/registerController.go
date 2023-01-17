package controllers

import (
	// "fmt"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		return
	}

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
  	if err != nil {
    	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		return
  	}

  	if len(p.Nome) == 0 || len(p.Senha) == 0 {
		// Adicionar JSON com mensagem
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		// jsonData := []byte(`{"msg":"erro", "status": "401"}`)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		return
	} else {

		// Caso o usuario nao exista
		if utils.FindUser(p.Nome, p.Senha) == false {
			// ele eh cadastrado
			data.RegisterData(p.Nome, p.Senha)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(map[string]any{
				"msg": "ok",
				"status": "201",
			})
			return

		} else {
			// ele nao eh cadastado
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(409)
			json.NewEncoder(w).Encode(map[string]any{
				"msg": "erro",
				"status": "409",
			})
			return
		}
	}
}