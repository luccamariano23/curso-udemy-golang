package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Insert
	stmt, _ := db.Prepare("insert into usuarios(nome) values (?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Lucca")

	//Pegar o ultimo id inserido
	// res.LastInsertId()
	// --------------------------
	//Pegar as linhas afetadas
	// res.RowsAffected()

	fmt.Println(res)
}
