package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Senha string `json:"senha"`
}

func ConnectDb() {

	DB, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Banco de dados conectado com sucesso.")

			results, err := DB.Query("SELECT id, nome, senha FROM users")
			if err != nil {
			    panic(err.Error()) 
			}

			for results.Next() {
			    var tag Tag
			    err = results.Scan(&tag.ID, &tag.Nome, &tag.Senha)
			    if err != nil {
			        panic(err.Error()) 
			    }
					fmt.Println(tag.ID) 
					fmt.Println(tag.Nome)
					fmt.Println(tag.Senha)
			}		
	}
}