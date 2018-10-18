package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// DBCon is the connection handle
	// for the database
	db  *sql.DB
	err error
)

// type Db struct {
// 	*sql.DB
// }

// var (
// 	id   int
// 	name string
// )

func ConnectDB() {
	//log.Println("connect")

	db, err = sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/golangtesdb") //if using password "root:password@tcp(127.0.0.1:3306)/golangtesdb")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err) // do something here
	}

	log.Println("connected")

	// rows, err := db.Query("select USER_ID, USERNAME from user where USER_ID = ?", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	defer db.Close()

}
