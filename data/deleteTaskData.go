package data 

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteTaskData(nomeTarefa string, usuarioId int) (bool) {
	db, err := sql.Open("mysql", "root:@/golang")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Banco de dados conectado com sucesso")
		defer db.Close()

		var sql string = fmt.Sprintf("DELETE FROM `tasks` WHERE nome = '%s' AND dono = '%d'", nomeTarefa, usuarioId)

		results, err := db.Query(sql)

		if err != nil {
			return false
		}

		return true
		fmt.Println(results)
	}
	return false
}