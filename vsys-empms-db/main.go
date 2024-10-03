package main

import (
	"fmt"
	"log"

	db "vsys.empms.dbhelper/db"
	srv "vsys.empms.dbhelper/server"
)

func main() {
	conn := db.DBconn()
	if conn != nil {
		fmt.Println("DB connected...\nDBHelper running")
		srv.Web()
	}
	log.Fatal("Error while connecting to DB")
}
