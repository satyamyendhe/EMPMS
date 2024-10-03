// Web Server
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	m "vsys.empms.commons/models"
	utils "vsys.empms.commons/utils"

	websecure "vsys.empms.commons/websecure"
	"vsys.empms.web/pages"
)

var (
	restHost = os.Getenv("REST_HOST")
	restPort = os.Getenv("REST_PORT")
	webPort  = os.Getenv("WEB_PORT")
	webHost  = os.Getenv("WEB_HOST")
)

func Web() {

	if strings.TrimSpace(restHost) == "" {
		restHost = "localhost" // or the appropriate default
	}
	if strings.TrimSpace(restPort) == "" {
		restPort = "7100"
	}
	if strings.TrimSpace(webPort) == "" {
		webPort = "3200"
	}
	if strings.TrimSpace(webHost) == "" {
		webHost = "0.0.0.0"
	}

	route := mux.NewRouter()

	route.Use(websecure.CommonMiddleware)

	route.HandleFunc("/", SendToRestLogin).Methods("GET")
	route.HandleFunc("/login", SendToRestLogin)
	route.HandleFunc("/signup", SendToRestSignUp)
	route.HandleFunc("/dashboard", GetDashboard).Methods("GET")
	route.HandleFunc("/get-emps", GetEmpsRest).Methods("GET")
	route.HandleFunc("/get-emp", GetEmpRest).Methods("POST")
	route.HandleFunc("/add-emp", AddEmpRest).Methods("POST")
	route.HandleFunc("/logout", LogoutHandler).Methods("GET")
	route.HandleFunc("/update", UpdateRest).Methods("PUT")
	route.HandleFunc("/delete-emp", DeleteRest).Methods("DELETE")
	route.HandleFunc("/AddModal", AddModal).Methods("GET")
	route.HandleFunc("/EditModal", EditModal).Methods("POST")
	route.HandleFunc("/get-logs", GetLogsRest).Methods("GET")
	route.HandleFunc("/logs", GetDashboard).Methods("GET")

	// Serve static files
	staticDir := "./static"
	route.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS middleware
	handler := c.Handler(route)

	log.Printf("Listening on http://%s:%s", webHost, webPort)
	if err := http.ListenAndServe(":"+webPort, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func SendToRestLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render login page for GET request
		loginStruct := pages.GetLogin{}
		html := loginStruct.Build()
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
		fmt.Println("GET method: Rendering login page")
		return
	}

	if r.Method == http.MethodPost {
		loginStruct := pages.GetLogin{}

		r.ParseForm()
		email := r.FormValue("email")
		pass := r.FormValue("password")

		users := m.GetLogin{
			Email:    email,
			Password: pass,
		}

		// Marshal login data to JSON
		marshalData, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Error marshalling login data", http.StatusInternalServerError)
			return
		}

		// Send POST request to REST service
		url := "http://" + restHost + ":" + restPort + "/login"
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))
		if err != nil {
			http.Error(w, "Error creating request to REST service", http.StatusInternalServerError)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error sending request to REST service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check the response status from REST service
		if resp.StatusCode == http.StatusAccepted {
			// Generate JWT token
			token, time, err := utils.CreateJwtToken(users.Email)
			if err != nil {
				http.Error(w, "Error generating token", http.StatusInternalServerError)
				return
			}

			// Set token in cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    token,
				Path:     "/",
				Expires:  time,
				HttpOnly: true,
			})

			// Redirect to dashboard
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			// Invalid credentials, render login page with error message
			loginStruct.ErrorMsg = `<div class="alert alert-danger border-2 d-flex align-items-center" role="alert" style="height: 40px;">
  <p class="mb-0 flex-1">Invalid Credentials, Please try again!</p>
  <button class="btn-close" type="button" data-bs-dismiss="alert" aria-label="Close"></button>
</div>`
			html := loginStruct.Build()
			w.Header().Set("Content-Type", "text/html")
			// w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(html))
			fmt.Println("POST method: Rendering login page with error")
			return
		}
	}
}

func SendToRestSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		signupStruct := pages.GetSignUp{}
		html := signupStruct.Build()
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
		fmt.Println("GET method: Rendering Signup page")
		return
	}

	if r.Method == http.MethodPost {

		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()
		email := r.FormValue("email")
		pass := r.FormValue("password")
		name := r.FormValue("name")

		users := m.GetLogin{
			Name:     name,
			Email:    email,
			Password: pass,
		}


		marshalData, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Error marshalling login data", http.StatusInternalServerError)
			return
		}

		url := fmt.Sprintf("http://%s:%s/signup", restHost, restPort)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))
		if err != nil {
			http.Error(w, "Error creating request to REST service", http.StatusInternalServerError)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Print("Error sending request to REST service", err)
			http.Error(w, "Error sending request to REST service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusAccepted {
			token, time, err := utils.CreateJwtToken(users.Email)
			if err != nil {
				http.Error(w, "Error generating token", http.StatusInternalServerError)
				return
			}

			// Set token in cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    token,
				Path:     "/",
				Expires:  time,
				HttpOnly: true,
			})
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		} else {
			signupStruct := pages.GetSignUp{}
			signupStruct.ErrorMsg = `<div class="alert alert-danger border-2 d-flex align-items-center" role="alert" style="height: 40px;">
  <p class="mb-0 flex-1">Email already exists</p>
  <button class="btn-close" type="button" data-bs-dismiss="alert" aria-label="Close"></button>
</div>`
			html := signupStruct.Build()
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(html))
			return
		}
	}
}

// Dashboard func
func GetDashboard(w http.ResponseWriter, r *http.Request) {

	email, _ := utils.GetCookieValueEmail(r, "auth_token")
	url := "http://" + restHost + ":" + restPort + "/setMail"
	marshalData, _ := json.Marshal(email)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))
	if err != nil {
		http.Error(w, "Error creating request to REST service", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Print("Error sending Mail to REST service", err)
		http.Error(w, "Error sending Mail to REST service", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {

		dashboardStruct := pages.Dashboard{}
		html := dashboardStruct.Build(w, r)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))

	}

}

func GetEmpsRest(w http.ResponseWriter, r *http.Request) {
	url := "http://" + restHost + ":" + restPort + "/get-emps"
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Error creating request to REST service", http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, _ := client.Do(res)

	var emp []m.Employees
	json.NewDecoder(resp.Body).Decode(&emp)

	json.NewEncoder(w).Encode(emp)
}

func AddEmpRest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp m.Employees
	json.NewDecoder(r.Body).Decode(&emp)

	// --------------
	marshalData, _ := json.Marshal(emp)
	url := "http://" + restHost + ":" + restPort + "/add-emp"
	res, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))

	if err != nil {
		http.Error(w, "Error creating request to REST service", http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(res)

	if resp.StatusCode != http.StatusOK || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Update Fail")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// logout handler
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the authentication token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // Expire the cookie immediately
		HttpOnly: true,
	})

	// Redirect to login page after logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Update Employee
func UpdateRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// url := "http://" + dbHelperHost + ":" + dbHelperPort + "/update/" + EmpID

	var emp m.Employees
	json.NewDecoder(r.Body).Decode(&emp)
	marshalData, _ := json.Marshal(emp)
	url := "http://" + restHost + ":" + restPort + "/update"
	res, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(marshalData))

	client := &http.Client{}
	resp, err := client.Do(res)

	if resp.StatusCode != http.StatusOK || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Update Fail")
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Delete Employee
func DeleteRest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var empID m.ForID
	json.NewDecoder(r.Body).Decode(&empID)

	marshalData, _ := json.Marshal(empID)

	url := "http://" + restHost + ":" + restPort + "/delete"
	res, _ := http.NewRequest(http.MethodDelete, url, bytes.NewReader(marshalData))

	client := &http.Client{}
	client.Do(res)
}

func GetEmpRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var empID m.ForID
	err := json.NewDecoder(r.Body).Decode(&empID)
	if err != nil {
		log.Print("Error while decoidng data in web :", err)
	}

	marshalData, _ := json.Marshal(empID)

	url := "http://" + restHost + ":" + restPort + "/get-emp"
	res, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))

	client := &http.Client{}
	resp, _ := client.Do(res)

	var emp m.Employees
	json.NewDecoder(resp.Body).Decode(&emp)
	json.NewEncoder(w).Encode(emp)

}

// Func to get emp data by ID
func FetchEmployee(id string) (m.Employees, error) {
	var empID m.ForID
	empID.ID = id
	marshalData, _ := json.Marshal(empID)

	if strings.TrimSpace(webPort) == "" {
		webPort = "3200"
	}
	if strings.TrimSpace(webHost) == "" {
		webHost = "0.0.0.0"
	}
	url := fmt.Sprint("http://", webHost, ":", webPort, "/get-emp")
	res, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalData))

	if err != nil {
		return m.Employees{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		return m.Employees{}, err
	}
	defer resp.Body.Close()

	var emp m.Employees
	json.NewDecoder(resp.Body).Decode(&emp)
	return emp, nil
}

func EditModal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var empID m.ForID
	err := json.NewDecoder(r.Body).Decode(&empID)
	if err != nil {
		log.Print("Error while decoidng data in web :", err)
	}

	Modal := &pages.Modals{}
	Modal.Action = "Edit"

	Modal.EmpId = empID.ID
	Modal.EmpData, _ = FetchEmployee(empID.ID)

	html := Modal.Build()

	// Write the HTML content directly to the response writer
	fmt.Fprint(w, html)
}

func AddModal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	Modal := &pages.Modals{}
	Modal.Action = "Add"
	html := Modal.Build()
	// Write the HTML content directly to the response writer
	fmt.Fprint(w, html)
}

// Handler function to Get LOGS
func GetLogsRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := "http://" + restHost + ":" + restPort + "/get-logs"
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Error creating request to REST service", http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, _ := client.Do(res)

	var log []m.Log
	json.NewDecoder(resp.Body).Decode(&log)
	json.NewEncoder(w).Encode(log)
}
