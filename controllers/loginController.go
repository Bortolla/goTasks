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


func LoginController(w http.ResponseWriter, r *http.Request) {
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

	// Caso o usuario nao exista
	if utils.FindUser(p.Nome, p.Senha) == false {
		// ele nao pode fzr login
		data.RegisterData(p.Nome, p.Senha)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "401",
		})
		return

	} else {
		// usuario existe, entao pode logar
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "ok",
			"status": "200",
		})
		return
	}

}