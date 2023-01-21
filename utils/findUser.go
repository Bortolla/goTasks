package utils 

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

func FindUser(nome, senha string) (int) {
	db, err := sql.Open("mysql", "root:root@/golang")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Banco de dados conectado com sucesso")
		defer db.Close()

		var sql string = fmt.Sprintf("SELECT * FROM `users` WHERE nome = '%s' AND senha = '%s'", nome, senha)

		results, err := db.Query(sql)

		if err != nil {
			panic(err.Error())
		}

		for results.Next() {
			var user User
			err = results.Scan(&user.ID, &user.Nome, &user.Senha)
		
			if err != nil {
				panic(err.Error())
			}

			if user.ID >= 1 {
				return user.ID
			} else {
				return 0
			}
		}
		
	}

	return 0	
}