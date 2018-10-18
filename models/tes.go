package models

import (
	// "./shared/database"
	_ "github.com/go-sql-driver/mysql"
)

var (
	id   int
	name string
)

func HomeTes() {

	//rows, err := database.Db.Query("select USER_ID, USERNAME from user where USER_ID = ?", 1)

	// rows, err := database.db.Query("select USER_ID, USERNAME from user where USER_ID = ?", 1)
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
}
