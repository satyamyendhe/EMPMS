package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	m "vsys.empms.commons/models"
	restUtils "vsys.empms.rest/services"

	"github.com/gorilla/mux"
)

var dbHelperHost string
var dbHelperPort string
var restPort string
var restHost string

// Routes all the url
func Web() {

	dbHelperHost = os.Getenv("DBHELPER_HOST")
	dbHelperPort = os.Getenv("DBHELPER_PORT")
	restPort = os.Getenv("RESTPORT")
	restHost = os.Getenv("REST_HOST")

	if strings.TrimSpace(restHost) == "" {
		restHost = "0.0.0.0"
	}
	if strings.TrimSpace(restPort) == "" {
		restPort = "7100"
	}
	if strings.TrimSpace(dbHelperHost) == "" {
		dbHelperHost = "0.0.0.0"
	}
	if strings.TrimSpace(dbHelperPort) == "" {
		dbHelperPort = "7200"
	}

	router := mux.NewRouter()

	router.HandleFunc("/login", SendToDBLogin)
	router.HandleFunc("/signup", SendToDBSignUp)
	router.HandleFunc("/get-emps", GetEmpsDB).Methods("GET")
	router.HandleFunc("/get-emp", GetEmpDB).Methods("POST")
	router.HandleFunc("/delete", DeleteEmpDB).Methods("DELETE")
	router.HandleFunc("/add-emp", AddEmpDB).Methods("POST")
	router.HandleFunc("/update", UpdateDataDB).Methods("PUT")
	router.HandleFunc("/setMail", sendMail).Methods("POST")
	router.HandleFunc("/get-logs", GetLogsDB).Methods("GET")

	log.Printf("Listening on  http://%v:%v", restHost, restPort)
	http.ListenAndServe(":"+restPort, router)
}

// Login Handler Func
func SendToDBLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user m.GetLogin
	json.NewDecoder(r.Body).Decode(&user)
	marshalData, _ := json.Marshal(user)

	// making request to DBhelper
	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/login"
	res, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Send the request to DBHelper
	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		http.Error(w, "Error sending request to DBHelper", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusAccepted {
		fmt.Println("Verify Rest")
		// Handle success response
		w.WriteHeader(http.StatusAccepted)
	} else {
		// Handle failure response
		w.WriteHeader(http.StatusUnauthorized)
	}

}

// Signup hander Func
func SendToDBSignUp(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Access-Control-Allow-Origin", "*")                  // Allow all origins
	// w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS") // Allow methods
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")      // Allow header

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")

	var user m.GetLogin
	json.NewDecoder(r.Body).Decode(&user)
	marshalData, _ := json.Marshal(user)

	// making request to DB
	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/signup"
	res, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))
	client := &http.Client{}
	resp, _ := client.Do(res)

	if resp.StatusCode == http.StatusAccepted {
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// Functon to get all employees data form db
func GetEmpsDB(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers
	// w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow methods
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow headers

	w.Header().Set("Content-Type", "application/json")

	// Make HTTP GET request to DB
	resp, err := http.Get("http://" + dbHelperHost + ":" + dbHelperPort + "/get-emps")
	if err != nil {
		log.Printf("Error making POST request: %v", err)
		http.Error(w, "Failed to get user information i am in restserv", http.StatusForbidden)
		return
	}
	defer resp.Body.Close()

	// checking if we get proper respons or not
	// Use RespondFailure function to handle non-successful responses
	restUtils.RespondFailure(resp, w)

	// decode json response
	var emp []m.Employees
	json.NewDecoder(resp.Body).Decode(&emp)

	// //  access the respons on web
	w.WriteHeader(http.StatusOK)

	// w.Write(responseData)
	json.NewEncoder(w).Encode(emp)

}

func AddEmpDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var emp m.Employees
	json.NewDecoder(r.Body).Decode(&emp)
	marshalData, _ := json.Marshal(emp)

	resp, err := http.Post("http://"+dbHelperHost+":"+dbHelperPort+"/add-emp", "application/json", bytes.NewReader(marshalData))
	if err != nil {
		log.Printf("Error making Post request: %v", err)
		http.Error(w, "Failed to add employee information", http.StatusForbidden)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteEmpDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empID m.ForID
	json.NewDecoder(r.Body).Decode(&empID)

	marshalData, _ := json.Marshal(empID)

	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/delete"
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewReader(marshalData))
	if err != nil {
		log.Printf("Error making Delete request: %v", err)
		http.Error(w, "Failed to delete employee information", http.StatusForbidden)
		return
	}

	// Idk much about below two line
	client := &http.Client{}    //aapni NewRequest ko execute karega
	reqs, err := client.Do(req) //Uss request se  response lega

	if err != nil {
		log.Printf("Error sending DELETE request: %v", err)
		http.Error(w, "Failed to send DELETE request", http.StatusInternalServerError)
		return
	}
	defer reqs.Body.Close()
	w.WriteHeader(http.StatusOK)
}

func UpdateDataDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/update"

	var emp m.Employees
	json.NewDecoder(r.Body).Decode(&emp)

	marshalData, _ := json.Marshal(emp)
	resp, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(marshalData))

	Client := &http.Client{}
	rep, _ := Client.Do(resp)

	// var emps m.Employees
	// json.NewDecoder(rep.Body).Decode(&emps)

	if rep.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func GetEmpDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empID m.ForID
	err := json.NewDecoder(r.Body).Decode(&empID)
	if err != nil {
		log.Print("Error while decoidng data in rest :", err)
	}

	marshalData, _ := json.Marshal(empID)

	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/get-emp"
	res, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))

	client := &http.Client{}
	resp, _ := client.Do(res)

	var emp m.Employees
	json.NewDecoder(resp.Body).Decode(&emp)
	json.NewEncoder(w).Encode(emp)

}

func sendMail(w http.ResponseWriter, r *http.Request) {
	var email string
	json.NewDecoder(r.Body).Decode(&email)
	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/setMail"
	marshalData, _ := json.Marshal(email)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))
	if err != nil {
		http.Error(w, "Error creating request to DB service", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Print("Error sending Mail to DB service", err)
		http.Error(w, "Error sending Mail to DB service", http.StatusInternalServerError)
		return
	}
}

func GetLogsDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := "http://" + dbHelperHost + ":" + dbHelperPort + "/get-logs"
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Error creating request to DB service", http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, _ := client.Do(res)

	var log []m.Log
	json.NewDecoder(resp.Body).Decode(&log)
	json.NewEncoder(w).Encode(log)
}
