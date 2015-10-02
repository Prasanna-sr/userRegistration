package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type userArr [60]userData
type userData struct {
	emailID  string
	password string
	name     string
	city     string
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
	var emailID sql.NullString
	var name sql.NullString
	var password sql.NullString
	var city sql.NullString
	var ua userArr
	i := 0
	ud := userData{}

	rows, err := db.Query("select * from users;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&emailID, &password, &name, &city)
		if err != nil {
			log.Fatal(err)
		}
		if emailID.Valid {
			ud.emailID = emailID.String
		} else {
			ud.emailID = ""
		}
		if password.Valid {
			ud.password = password.String
		} else {
			ud.password = ""
		}
		if name.Valid {
			ud.name = name.String
		} else {
			ud.name = ""
		}
		if city.Valid {
			ud.city = city.String
		} else {
			ud.city = ""
		}
		ua[i] = ud
		i++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return ua
}
