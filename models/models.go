package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type userArr [60]userData
type userData struct {
	emailID string
	name    string
}
type error interface {
	Error() string
}

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("mysql", "mysql@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("mysql connection ok")
	}
	// defer db.Close()
}

func FetchAllUsers() userArr {
	var emailID string
	// var name string
	var ua userArr
	i := 0
	ud := userData{"", ""}

	rows, err := db.Query("select emailID from users;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&emailID)
		if err != nil {
			log.Fatal(err)
		}
		ud.emailID = emailID
		ua[i] = ud
		i++
	}
	return ua
}
