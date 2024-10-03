package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	DB "vsys.empms.dbhelper/db"
)

var (
	db *sql.DB = DB.DBconn()
)

// Func to signup user -> add his data to DB.
func AddUser(name, email, password string) error {
	var isAdmin = true
	var currentDate = time.Now()
	fmt.Println(name, ":::", email, ":::", password)
	_, err := db.Exec(`INSERT INTO "user" (name, email, pass, isAdmin, createdOn) VALUES ($1, $2, $3, $4 ,$5)`, name, email, password, isAdmin, currentDate)
	if err != nil {
		log.Printf("Error while inserting data into table in  userHandler : %v", err)
		return err
	}
	return nil

}

func LoginUser(email, password string) error {
	var dbPass string
	var isAdmin bool

	err := db.QueryRow(`SELECT pass , isadmin from "user" WHERE  email = $1`, email).Scan(&dbPass, &isAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found")
			return err
		}
		log.Printf("Error while getting loign details in userHandler")
		return err
	}

	fmt.Println(dbPass)
	if dbPass != password {
		return fmt.Errorf("invalid pass") //--> we are expecting error in return that why we format string to error using fmt.Errorf
	}

	if !isAdmin {
		return fmt.Errorf("not an admin")
	}
	return nil
}
