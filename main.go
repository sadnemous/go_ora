package main

import (
	"database/sql"
	"fmt"

	"example.com/go_ora/config"
	"example.com/go_ora/ora_db"
)

func main() {
	var c config.Config
	c.GetConfig()
	fmt.Println(c)
	var db *sql.DB
	db = ora_db.ConnectToOracle(&c)

	str_time := ora_db.GetSystemTIme(db)
	fmt.Println(str_time)

	ora_db.GetDemoData(db)
	db.Close()
}
