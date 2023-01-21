package controllers

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"golang/data"
	// "golang/utils"
)

type GetTasksResponse struct {
	UsuarioId int
}

func GetTasksController(w http.ResponseWriter, r *http.Request) {
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

	var t GetTasksResponse

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

	if t.UsuarioId < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "erro",
			"status": "404",
		})
		return
	} else {
		// Buscar a tarefa referente aquele usuario
		var resultado = data.GetTasksData(t.UsuarioId)
		// fmt.Println(len(resultado))
		// fmt.Println(resultado)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]any{
			"msg": "ok",
			"status": "200",
			"tasks": resultado,
		})
		return
	}

}
