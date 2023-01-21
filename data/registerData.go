package data

import (
	"fmt"
	"database/sql"
	// "strconv"
	// "log"
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	Nome string `json:"nome"`
	Senha   string    `json:"senha"`
}

func RegisterData(nome, senha string)  {
	db, err := sql.Open("mysql", "root:root@/golang")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var sql string = fmt.Sprintf("INSERT INTO `users` (`nome`, `senha`) VALUES('%s', '%s')", nome, senha)
	results, err := db.Exec(sql)
  if err != nil {
    panic(err.Error())
  }

	fmt.Println(results)

}