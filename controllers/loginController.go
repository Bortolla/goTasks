package controllers

import (
	"encoding/json"
	"net/http"
	// "golang/data"
	"golang/utils"
)


func LoginController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	var usuarioId = utils.FindUser(p.Nome, p.Senha)
	if usuarioId < 1 {
		// ele nao pode fzr login
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
			"usuarioId": usuarioId, 
		})
		return
	}

}