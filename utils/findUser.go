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

func FindUser(nome, senha string) (bool) {
	db, err := sql.Open("mysql", "root:@/golang")

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

			if len(user.Nome) != 0 {
				return true
			} else {
				return false
			}

		}
		
	}

	return false
	

	// Fazendo a query 
	// var sql string = fmt.Sprintf("SELECT * FROM `users` WHERE nome = '%s' AND senha = '%s'", nome, senha)

	// row := db.Query(sql)

  // // if row != nil {
	// // 	// if row != sql.ErrNoRows {
  // //   // 	panic(err.Error())
	// // 	// }
	// // 	return false
  // // }

	// temp := ""
	// row.Scan(&temp)

	// fmt.Println(temp)

	// if temp != "" {
	// 	fmt.Println(temp)
	// 	return true
	// }

	// return false

	// if results != nil {
	// 	fmt.Println("Usuario ja existe")
	// 	return "sim"
	// } else {
	// 	return "nao"
	// }
	
	// return "nada"
}