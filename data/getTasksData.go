package data

import (
	"fmt"
	// "log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Nome string `json:"nome"`
}

func GetTasksData(usuarioId int) ([]string) {
	array := make([]string, 0)
	db, err := sql.Open("mysql", "root:@/golang")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Banco de dados conectado com sucesso")
		defer db.Close()

		var sql string = fmt.Sprintf("SELECT nome FROM `tasks` WHERE dono = '%d'", usuarioId)

		results, err := db.Query(sql)

		if err != nil {
			fmt.Println("erro")		
			return array
		}

		for results.Next() {
			var task Task

			err = results.Scan(&task.Nome)

			if err != nil {
				panic(err.Error())
			}

			array = append(array, task.Nome)
		}
		return array
	}
	return array
}