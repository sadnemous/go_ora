package ora_db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"example.com/go_ora/config"
	_ "github.com/sijms/go-ora/v2"
)

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func ConnectToOracle(Conf *config.Config) *sql.DB {
	connStr := Conf.Driver + "://" + Conf.Username + ":" + Conf.Password + "@" + Conf.Server + ":" + strconv.Itoa(int(Conf.Port)) + "/" + Conf.Service
	// fmt.Println("Conn String: " + connStr)
	// connStr = "oracle://demo:demo@localhost:1521/XE"
	// fmt.Println("Conn String: " + connStr)

	//open connection
	db, err := sql.Open("oracle", connStr)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	// if we keep following code, it will be executed once function
	// scope dies, therefore db will be returned but closed
	// ref: https://stackoverflow.com/questions/57867635/mysql-database-gets-error-in-golang-sql-database-is-closed
	//
	//   defer func() {
	//   	err = db.Close()
	//   	if err != nil {
	//   		fmt.Println("Cannot close connection: ", err)
	//   	}
	//   }()
	//

	//Check connection by pinging
	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}
	return db
}

func GetSystemTIme(db *sql.DB) string {
	var queryResultColumnOne string
	row := db.QueryRow("SELECT systimestamp FROM dual")
	err2 := row.Scan(&queryResultColumnOne)
	if err2 != nil {
		panic(fmt.Errorf("error scanning db: %w", err2))
	}
	outStr := fmt.Sprintf("The time in the database %s", queryResultColumnOne)
	return outStr
}

func GetDemoData(db *sql.DB) {

	// sys_time := GetSystemTIme(db)
	// fmt.Println(sys_time)
	// fetching multiple rows

	theRows, err := db.Query("select bookid, bname from Library")
	handleError("Query for multiple rows", err)
	defer theRows.Close()

	var (
		bookid string
		bname  string
	)

	fmt.Printf("bookid : bname\n")
	for theRows.Next() {
		err := theRows.Scan(&bookid, &bname)
		handleError("next row in multiple rows", err)
		fmt.Printf("%s : %s \n", bookid, bname)
	}
	err = theRows.Err()
	handleError("next row in multiple rows", err)
}
