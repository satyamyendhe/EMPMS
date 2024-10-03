package dao

import (
	// "database/sql"

	"fmt"
	"log"
	"strconv"
	"strings"

	models "vsys.empms.commons/models"
)

func GetEmps() ([]models.Employees, error) {
	rows, err := db.Query("SELECT * FROM employee ORDER BY id;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	emps := []models.Employees{}
	for rows.Next() {
		var e models.Employees
		err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Department,
			&e.Designation,
			&e.Joindate,
			&e.Birthdate,
			&e.Gender,
			&e.Email,
			&e.Address,
			&e.Mobile,
			&e.Salary,
		)
		if err != nil {
			return nil, err
		}
		// experimantal code
		date := strings.Split(e.Joindate, "T")
		e.Joindate = date[0]

		date = strings.Split(e.Birthdate, "T")
		e.Birthdate = date[0]
		emps = append(emps, e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return emps, nil
}

func GetEmp(id int) (models.Employees, error) {
	var emp models.Employees
	db.QueryRow("SELECT * FROM employee WHERE id = $1", id).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Department,
		&emp.Designation,
		&emp.Joindate,
		&emp.Birthdate,
		&emp.Gender,
		&emp.Email,
		&emp.Address,
		&emp.Mobile,
		&emp.Salary,
	)

	// fmt.Println(emp)
	return emp, nil
}

func DeleteEmp(id int) error {
	// var emp models.Employees
	_, err := db.Query(`DELETE FROM "employee" WHERE id= $1;`, id)
	if err != nil {
		log.Printf("Error while deleting employee : %v", err)
		return err
	}
	fmt.Printf("Deleted Successfuly with id : %v\n", id)
	return nil
}

func AddEmp(emp models.Employees) error {
	_, err := db.Exec(`INSERT INTO employee (name, department, designation, dob, doj, gender, email, address, mobile, salary)
		VALUES
		($1, $2, $3, $4, $5, $6,$7, $8, $9,$10)
	`, emp.Name, emp.Department, emp.Designation, emp.Birthdate, emp.Joindate, emp.Gender, emp.Email, emp.Address, emp.Mobile, emp.Salary)
	if err != nil {
		log.Printf("Error while adding employee : %v", err)
		return err
	}

	return nil
}

func UpdateEmp(emp models.Employees) error {

	empIdInt, _ := strconv.Atoi(emp.ID)
	_, err := db.Exec(`UPDATE employee SET name=$1, department=$2, designation=$3, dob=$4, doj=$5, gender=$6, email=$7, address=$8, mobile=$9, salary=$10 WHERE id = $11;`,
		emp.Name,
		emp.Department,
		emp.Designation,
		emp.Birthdate,
		emp.Joindate,
		emp.Gender,
		emp.Email,
		emp.Address,
		emp.Mobile,
		emp.Salary,
		empIdInt,
	)
	if err != nil {
		log.Printf("Error while updating employee ")
		return err
	}
	return nil
}
