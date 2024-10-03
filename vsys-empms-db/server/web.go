package server

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"net/http"
	"strconv"

	// "net/url"

	"github.com/gorilla/mux"
	m "vsys.empms.commons/models"
	dao "vsys.empms.dbhelper/dao"
)

var dbHelperHost string
var dbHelperPort string
var UName string

func Web() {

	dbHelperHost = os.Getenv("DBHELPER_HOST")
	dbHelperPort = os.Getenv("DBHELPER_PORT")
	if strings.TrimSpace(dbHelperHost) == "" {
		dbHelperHost = "0.0.0.0"
	}
	if strings.TrimSpace(dbHelperPort) == "" {
		dbHelperPort = "7200"
	}

	router := mux.NewRouter()

	router.HandleFunc("/login", Login)
	router.HandleFunc("/signup", Signup)
	router.HandleFunc("/get-emps", GetEmps).Methods("GET")
	router.HandleFunc("/get-emp", GetEmp).Methods("POST")
	router.HandleFunc("/delete", DeleteEmp).Methods("DELETE")
	router.HandleFunc("/add-emp", AddEmp).Methods("POST")
	router.HandleFunc("/update", UpdateData).Methods("PUT")
	router.HandleFunc("/setMail", UserName).Methods("POST")
	router.HandleFunc("/get-logs", GetLogs).Methods("GET")

	http.ListenAndServe(":"+dbHelperPort, router)
}

// Login Verification
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user m.GetLogin
	json.NewDecoder(r.Body).Decode(&user)

	err := dao.LoginUser(user.Email, user.Password)
	if err != nil {
		log.Println("login error")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("Verify login")
	w.WriteHeader(http.StatusAccepted)
}

// signup handler
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user m.GetLogin
	json.NewDecoder(r.Body).Decode(&user)

	err := dao.AddUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Some DB code here
	w.WriteHeader(http.StatusAccepted)

}

// handler func to get all employees to get from  DB
func GetEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	emp, err := dao.GetEmps()
	if err != nil {
		http.Error(w, "Error while getting emp", http.StatusInternalServerError)
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(emp)
	if err != nil {
		http.Error(w, "Error whille encoding data", http.StatusInternalServerError)
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusOK)
}

// Handler function to delete data form employee
func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empID m.ForID
	json.NewDecoder(r.Body).Decode(&empID)
	// Convert the ID to an integer
	empIDInt, err := strconv.Atoi(empID.ID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	emp, _ := dao.GetEmp(empIDInt)
	dao.Addlog(emp.Name, emp.Email, emp.Designation, "Deleted", UName)
	// // Call the DAO function to delete the employee
	err = dao.DeleteEmp(empIDInt)
	if err != nil {
		http.Error(w, "Error while deleting employee", http.StatusInternalServerError)
		return
	}
	// // Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted with id " + empID.ID)
}

func AddEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp m.Employees
	json.NewDecoder(r.Body).Decode(&emp)

	// Convert date strings to time.Time

	err := dao.AddEmp(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dao.Addlog(emp.Name, emp.Email, emp.Designation, "Added", UName)
	log.Println("Employee details added")
	w.WriteHeader(http.StatusOK)

}

// // Handler function using unmarshal to add employees

// func AddEmp(w http.ResponseWriter, r *http.Request) {
// 	var emp temp.Employee
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	json.Unmarshal(reqBody, &emp)

// 	test = append(test, emp)
// 	marshalData, _ := json.Marshal(test)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(marshalData)
// }

// Handler function to update data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var emp m.Employees
	json.NewDecoder(r.Body).Decode(&emp)

	err := dao.UpdateEmp(emp)
	if err != nil {
		http.Error(w, "Error while adding employee", http.StatusInternalServerError)
		return
	}
	dao.Addlog(emp.Name, emp.Email, emp.Designation, "Updated", UName)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("updated")
}

func GetEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empID m.ForID
	err := json.NewDecoder(r.Body).Decode(&empID)
	if err != nil {
		log.Print("Error while decoidng data in DB :", err)
	}
	// Convert the ID to an integer
	empIDInt, err := strconv.Atoi(empID.ID)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// // Call the DAO function to delete the employee
	emp, err := dao.GetEmp(empIDInt)
	if err != nil {
		http.Error(w, "Error while deleting employee", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

func UserName(w http.ResponseWriter, r *http.Request) {
	var email string
	json.NewDecoder(r.Body).Decode(&email)

	name, _ := dao.GetUserName(email)
	UName = name
}

// handler func to get all logs from  DB
func GetLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log, err := dao.GetLogs()
	if err != nil {
		http.Error(w, "Error while getting emp", http.StatusInternalServerError)
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(log)
	if err != nil {
		http.Error(w, "Error whille encoding data", http.StatusInternalServerError)
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// w.WriteHeader(http.StatusOK)
}
