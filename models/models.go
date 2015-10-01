package models

import (
	"database/sql"
)

var db, err = sql.Open("mysql",
	"root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

func GetAllUsers() {
	rows, err := db.Query("select * from users;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&emailID, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(emailID, name)
	}

}
