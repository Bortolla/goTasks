package data 

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Senha string `json:"senha"`
}

func CreateTaskData(nomeTarefa string, usuarioId int) (bool) {
	db, err := sql.Open("mysql", "root:@/golang")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Banco de dados conectado com sucesso")
		defer db.Close()

		var sql string = fmt.Sprintf("INSERT INTO `tasks` (nome, dono) VALUES('%s', '%s')", nomeTarefa, usuarioId)

		results, err := db.Query(sql)

		if err != nil {
			return false
		}

		return true
		fmt.Println(results)
	}
	return false
}