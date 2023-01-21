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

func GetTasksData(usuarioId int) {
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
			return
		}

		//fmt.Println(results)

		for results.Next() {
			var task Task

			err = results.Scan(&task.Nome)

			if err != nil {
				panic(err.Error())
			}

			fmt.Println(task.Nome)

		}
	
		//fmt.Println(results)
		return 
	}
	return
	
}