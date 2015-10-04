package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type userArr []userData
type userData struct {
	EmailID  string
	Password string
	Name     string
	City     string
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
			ud.EmailID = emailID.String
		}
		if password.Valid {
			ud.Password = password.String
		}
		if name.Valid {
			ud.Name = name.String
		}
		if city.Valid {
			ud.City = city.String
		}
		ua = append(ua, ud)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return ua
}
