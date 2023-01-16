package database

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

func ConnectDb(operation string, nome string, senha string) (string) {

	// "mysql", "usuario:senha@/nomeBancoDeDados"
	DB, err := sql.Open("mysql", "root:@/golang")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Banco de dados conectado com sucesso.")
		
		if (operation == "GetUser") {
			results, err := DB.Query("SELECT id, nome, senha FROM users")
			if err != nil {
					panic(err.Error()) 
			}

			for results.Next() {
					var user User
					err = results.Scan(&user.ID, &user.Nome, &user.Senha)
					if err != nil {
							panic(err.Error()) 
					}
					fmt.Println(user.ID) 
					fmt.Println(user.Nome)
					fmt.Println(user.Senha)
			}		
		}
	
	return "fim"

	}
}