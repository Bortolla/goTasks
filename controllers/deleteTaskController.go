package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"golang/data"
)

func DeleteTaskController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "500",
		})
		fmt.Println("Erro no método - createTaskController")
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
		fmt.Println("Erro na decodificação do JSON - createTaskController")
		return
  	}

	if t.UsuarioId < 1 || len(t.NomeTarefa) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "404",
		})
		fmt.Println("UsuarioId ou NomeTarefa vazios")
		return
	} else {
		var resultado = data.DeleteTaskData(t.NomeTarefa, t.UsuarioId)
		
		if resultado == true {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{
				"msg": "ok",
				"status": "200",
			})	
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(map[string]any{
				"msg": "erro",
				"status": "500",
			})
			return
		}
	
	}
}