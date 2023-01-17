package controllers

import (
	"encoding/json"
	"net/http"
	"golang/data"
	"golang/utils"
)

type Task struct {
	NomeUsuario string
	NomeTarefa string
}

func GetTasksController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		return
	}

	var t Task

	err := json.NewDecoder(r.Body).Decode(&t)
  	if err != nil {
    	w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		return
  	}

	if len(t.NomeTarefa) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		return
	} else {
		// Buscar a tarefa referente aquele usuario
	}

}







