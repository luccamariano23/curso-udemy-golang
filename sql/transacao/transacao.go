package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id,nome) values (?,?)")
	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Teste")

	_, err = stmt.Exec(1, "Thiago") //chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	} else {
		tx.Commit()
	}

}
