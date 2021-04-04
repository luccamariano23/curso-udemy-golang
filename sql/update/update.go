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

	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id=?")
	stmt.Exec("Teste Update", 1)

	//Delete
	stmt2, _ := db.Prepare("delete usuarios where id=?")
	stmt2.Exec(1)
}
