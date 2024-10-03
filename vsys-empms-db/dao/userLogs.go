package dao

import (
	"fmt"
	"log"

	m "vsys.empms.commons/models"
)

func GetUserName(email string) (string, error) {
	var name string

	fmt.Println(email)
	err := db.QueryRow(`SELECT name FROM "user" WHERE email = '` + email + `'; `).Scan(&name)
	if err != nil {
		log.Fatalf("Error while getting userName : %v", err)
		return "", err
	}

	return name, nil
}

func Addlog(empName, empEmail, empDeg, operation, userName string) error {
	_, err := db.Exec(`INSERT INTO logs (emp_name, emp_email, emp_designation, operation, updated_by)
		VALUES
		($1, $2, $3, $4, $5)
	`, empName, empEmail, empDeg, operation, userName)
	if err != nil {
		log.Printf("Error while adding logs : %v", err)
		return err
	}

	return nil
}

func GetLogs() ([]m.Log, error) {
	rows, err := db.Query("SELECT * FROM logs ORDER BY id DESC;")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var logs []m.Log

	for rows.Next() {
		var l m.Log
		err := rows.Scan(
			&l.ID,
			&l.Created,
			&l.EmpName,
			&l.EmpEmail,
			&l.EmpDeg,
			&l.Operation,
			&l.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}
	return logs, nil
}
