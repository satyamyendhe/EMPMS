package pages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	m "vsys.empms.commons/models"
	u "vsys.empms.commons/utils"
)

type Dashboard struct {
	Sidebar Sidebar
	Tabels  Tabel
	Modal   Modals
	Logs    Logs
}

var (
	webHost = os.Getenv("WEB_HOST")
	webPort = os.Getenv("WEB_PORT")
)

func (l *Dashboard) Build(w http.ResponseWriter, r *http.Request) string {

	var MainContentHtml string
	var PageName string

	switch r.URL.Path {
	case "/dashboard":
		PageName = "Dashboard"
		MainContentHtml += l.Tabels.Build()
		l.Sidebar.Button = `<button id="addEmpModal" class="btn btn-primary me-3" type="button">Add Employees</button>`
	case "/logs":
		PageName = "Logs"
		MainContentHtml += l.Logs.Build()
	}

	return u.JoinStr(`
  <!DOCTYPE html>
  <html lang="en-US" dir="ltr">
  
  <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  
  
  <!-- ===============================================-->
  <!--    Document Title-->
  <!-- ===============================================-->
  <title>`, PageName, `</title>
  
  
  <!-- ===============================================-->
  <!--    Favicons-->
  <!-- ===============================================-->
  
  <!-- <link rel="manifest" href="../../assets/img/favicons/manifest.json"> -->
  <script src="../static/falcon/public/assets/js/config.js"></script>
  <script src="../static/falcon/public/vendors/simplebar/simplebar.min.js"></script>
  
  
  <link rel="stylesheet" href="../static/falcon/public/vendors/fontawesome/all.min.js">

      <!-- BOX ICONS CSS -->
      <link rel="stylesheet" href="../static/boxicons/css/boxicons.css">
  <!-- ===============================================-->
  <!--    Stylesheets-->
  <!-- ===============================================-->

  <link href="../static/falcon/public/assets/css/theme.min.css" rel="stylesheet" id="style-default">
  
  
  <!------------------->
  <!-- custom styles -->
  <style>
  
  
  #container{
    padding: 0 2%; 
    width: 100%;
    }
    
    
    
  table thead tr{
    font-size: medium;
  }

  table tbody tr{
    font-size: 1rem;
  }
    .fa-trash-alt{
      color: red !important; 
      
      }
      .fa-edit{
      color: #2C7BE5 !important;
    }

   .modal-dialog {
    max-width: 70vh !important;
    min-width: 50vh !important;
  }
  
    
</style>
</head>


<body>

<!-- ===============================================-->
<!--    Main Content-->
<!-- ===============================================-->
<body class="main" id="top">

        <!--------------------->
        <!-------modal--------->
        <!--------------------->
        <div id="ModalHere">
        </div>
          
 

  <div id="container">

        <!--------------------->
        <!----side and nav ---->
        <!--------------------->    
        `, l.Sidebar.Build(), `
        
        <!--------------------->
        <!-- Tabels for EMp  -->
        <!--------------------->
          `, MainContentHtml, `
      </div>
    </div>
    <!-- ===============================================-->
    <!--    JavaScripts-->
    <!-- ===============================================-->

    
    
    <script src="../static/falcon/public/vendors/popper/popper.min.js"></script>
    <script src="../static/falcon/public/vendors/bootstrap/bootstrap.min.js"></script>
    <script src="../static/falcon/public/vendors/anchorjs/anchor.min.js"></script>
    <script src="../static/falcon/public/vendors/is/is.min.js"></script>
    <script src="../static/falcon/public/vendors/prism/prism.js"></script>
    <script src="../static/falcon/public/vendors/fontawesome/all.min.js"></script>
    <script src="../static/falcon/public/vendors/lodash/lodash.min.js"></script>
    <script src="../static/falcon/public/vendors/list.js/list.min.js"></script>
    <script src="../static/falcon/public/assets/js/theme.js"></script>
    <script src="../static/assets/index.js"></script>
    </body>
    
    </html>
    
    
    `)
}

//------------------------------------------------------------------------------------------------------------------
// -----------------------------------------MAIN CODE ABOVE---------------------------------------------------------
//------------------------------------------------------------------------------------------------------------------

// -------------------------
// ------- Sidebar ---------
// -------------------------
type Sidebar struct {
	Name     string
	NavItems []map[string]string
	Button   string
}

// Build method to construct the sidebar
func (s *Sidebar) Build() string {

	// Side bar data
	s.Name = "EMPMS"
	// In the MAP key is name of nav item and value is there classes
	s.NavItems = []map[string]string{
		{"Dashboard": "bx bxs-dashboard",
			"href": "/dashboard"},
		{"Logs": "bx bx-history",
			"href": "/logs"},
		{"Logout": "bx bx-log-out-circle",
			"href": "/logout"},
	}

	var NavItems string
	for _, item := range s.NavItems {
		for key, value := range item {
			if key == "href" {
				continue
			}
			NavItems += fmt.Sprintf(`
                  <!-- parent pages--><a class="nav-link" href="%s" role="button">
                    <div class="d-flex align-items-center mt-2">
                      <span class="nav-link-icon"><span class="%s fa-lg"></span></span>
                      <span class="nav-link-text h5 p1-2">%s</span>
                    </div>
                  </a>
`, item["href"], value, key)
		}
	}
	return u.JoinStr(`
  <nav class="navbar navbar-light navbar-vertical navbar-expand-xl">
          <div class="d-flex align-items-center">
            <div class="toggle-icon-wrapper">
              <button class="btn navbar-toggler-humburger-icon navbar-vertical-toggle" data-bs-toggle="tooltip"
                data-bs-placement="left" title="Toggle Navigation"><span class="navbar-toggle-icon"><span
                    class="toggle-line"></span></span></button>
            </div>
            <a class="navbar-brand" href="">
              <div class="d-flex align-items-center py-3"><span class="font-sans-serif">`, s.Name, `</span></div>
            </a>
          </div>
          <div class="collapse navbar-collapse" id="navbarVerticalCollapse">
            <div class="navbar-vertical-content scrollbar">
              <ul class="navbar-nav flex-column mb-3" id="navbarVerticalNav">
                <li class="nav-item">
                  <!-- label-->
                  <div class="row navbar-vertical-label-wrapper mt-3 mb-2">
                    <div class="col-auto navbar-vertical-label">
                    </div>
                    <div class="col ps-0">
                      <hr class="mb-0 navbar-vertical-divider" />
                    </div>
                  </div>
                    `, NavItems, `
                </li>
              </ul>
            </div>
          </div>
        </nav>
        <div class="content">
          <nav class="navbar navbar-light navbar-glass navbar-top navbar-expand">
            <button class="btn navbar-toggler-humburger-icon navbar-toggler me-1 me-sm-3" type="button"
              data-bs-toggle="collapse" data-bs-target="#navbarVerticalCollapse" aria-controls="navbarVerticalCollapse"
              aria-expanded="false" aria-label="Toggle Navigation"><span class="navbar-toggle-icon"><span
                  class="toggle-line"></span></span></button>
            <a class="navbar-brand me-1 me-sm-3" href="">
              <div class="d-flex align-items-center"><span class="font-sans-serif">`, s.Name, `</span></div>
            </a>
            <ul class="navbar-nav navbar-nav-icons ms-auto flex-row align-items-center">
              <li class="nav-item">
                <div class="theme-control-toggle fa-icon-wait d-flex px-2">
                  `, s.Button, `
                  <input class="form-check-input ms-0 theme-control-toggle-input" id="themeControlToggle" type="checkbox"
                    data-theme-control="theme" value="dark" />
                  <label class="mb-0 theme-control-toggle-label theme-control-toggle-light" for="themeControlToggle"
                    data-bs-toggle="tooltip" data-bs-placement="left" title="Switch to light theme"><span
                      class="fas fa-sun fs-0"></span></label>
                  <label class="mb-0 theme-control-toggle-label theme-control-toggle-dark" for="themeControlToggle"
                    data-bs-toggle="tooltip" data-bs-placement="left" title="Switch to dark theme"><span
                      class="fas fa-moon fs-0"></span></label>
                </div>
              </li>
            </ul>
          </nav>   
    `)
}

// -------------------------
// ---------Table-----------
// -------------------------

type Tabel struct {
	Columns   []Column
	Employees []m.Employees
}

type Column struct {
	ColName  string
	DataName string
	Class    string
	Sort     bool
}

var (
	col = []Column{
		{ColName: `<div class="form-check mb-0">
                      <input class="form-check-input" type="checkbox" id="allCheckBox" data-bulk-select='{"body":"bulk-select-body","actions":"bulk-select-actions","replacedElement":"bulk-select-replace-element"}' />
               </div>`, DataName: "Checkbox", Sort: false, Class: `class="align-middle white-space-nowrap"`},
		{ColName: "Name", DataName: "Name", Sort: true},
		{ColName: "Email", DataName: "Email", Sort: false},
		{ColName: "Department", DataName: "Department", Sort: true},
		{ColName: "Designation", DataName: "Designation", Sort: true},
		{ColName: "Join Date", DataName: "Joindate", Sort: true},
		{ColName: "Birth Date", DataName: "Birthdate", Sort: true},
		{ColName: "Address", DataName: "Address", Sort: false},
		{ColName: "Gender", DataName: "Gender", Sort: false},
		{ColName: "Mobile", DataName: "Mobile", Sort: false},
		{ColName: "Salary", DataName: "Salary", Sort: true},
		{ColName: "Action", DataName: "Action", Sort: false},
	}
)

// Fetch all employee details
func (T *Tabel) BuildHead() (string, []string) {
	T.Columns = col

	var thHTML string
	var sortSlice []string
	for _, value := range col {
		var thClass string
		if value.Sort {
			thClass = fmt.Sprint(`class="sort" data-sort="`, value.DataName, `"`)
			sortSlice = append(sortSlice, value.DataName)
		}
		thHTML += fmt.Sprint(`<th `, thClass, ` `, value.Class, `>`, value.ColName, `</th>`)
	}
	return thHTML, sortSlice
}

func FetchEmployees() ([]m.Employees, error) {
	if strings.TrimSpace(webPort) == "" {
		webPort = "3200"
	}
	if strings.TrimSpace(webHost) == "" {
		webHost = "0.0.0.0"
	}
	url := fmt.Sprint("http://", webHost, ":", webPort, "/get-emps")
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var emp []m.Employees
	json.NewDecoder(resp.Body).Decode(&emp)
	return emp, nil
}

func mapEmployeeData(emp m.Employees) map[string]string {
	return map[string]string{
		"Checkbox": `
            
                <div class="form-check mb-0">
                  <input class="form-check-input" type="checkbox" data-bulk-select-row="data-bulk-select-row" />
                </div>
           `,
		"Name":        emp.Name,
		"Email":       emp.Email,
		"Department":  emp.Department,
		"Designation": emp.Designation,
		"Joindate":    emp.Joindate,
		"Birthdate":   emp.Birthdate,
		"Address":     emp.Address,
		"Gender":      emp.Gender,
		"Mobile":      emp.Mobile,
		"Salary":      emp.Salary,
		"Action": `
                <div>
                  <button class="btn btn-link p-0 edit-btn" data-id="` + emp.ID + `" type="button" data-bs-placement="top" title="Edit"><span class="text-500 fas fa-edit"></span></button>
                  <button class="btn btn-link p-0 ms-2 delete-btn" id="delete-emp" data-id="` + emp.ID + `" type="button" data-bs-toggle="tooltip" data-bs-placement="top" title="Delete"><span class="text-500 fas fa-trash-alt"></span></button>
                </div>
              `,
	}
}

func (T *Tabel) BuildBody() string {
	var trBody string
	for _, emp := range T.Employees {
		dataMap := mapEmployeeData(emp)

		trBody += `<tr>`
		for _, column := range T.Columns {
			value := dataMap[column.DataName]
			trBody += fmt.Sprint(`<td class = "`, column.DataName, `">`, value, `</td>`)
		}
		trBody += `</tr>`
	}
	return trBody
}

func (T *Tabel) Build() string {
	// Fetch data
	data, err := FetchEmployees()
	if err != nil {
		log.Print("Error -", err)
		return ""
	}
	T.Employees = data
	Tabledata, sortSlice := T.BuildHead()
	sort, _ := json.Marshal(sortSlice)

	// Check if data is empty
	if len(data) == 0 {
		return `
<div class="card p-5">
  <div class="tab-pane preview-tab-pane active" role="tabpanel" aria-labelledby="tab-dom-5a3c857c-8deb-4307-af64-a62f27089443" id="dom-5a3c857c-8deb-4307-af64-a62f27089443">
    <div>
      <div class="d-flex justify-content-between g-0">
          <div class="col-auto">
            <p style="user-select: none; font-size: 3vh;">Employee List:</p>
          </div>
          <div class="col-auto col-sm-3 mb-3">
            <form>
              <div class="input-group">
                <input class="form-control form-control-sm shadow-none search" type="search"
                placeholder="Search..." aria-label="search" />
                <div class="input-group-text bg-transparent">
                  <span class="fa fa-search fs--1 text-600"></span>
                </div>
              </div>
            </form>
          </div>
      </div>
    </div>
  </div>
</div>
    `
	}

	return u.JoinStr(`
  <div class="card p-5">
    <div class="tab-pane preview-tab-pane active" role="tabpanel"
      aria-labelledby="tab-dom-5a3c857c-8deb-4307-af64-a62f27089443"
      id="dom-5a3c857c-8deb-4307-af64-a62f27089443">
      <div id="tableExample3" data-list='{"valueNames":` + string(sort) + `,"page":10,"pagination":true}'>
        <div class="d-flex justify-content-between g-0">
          <div class="col-auto">
            <p style="user-select: none; font-size: 3vh;">Employee List:</p>
          </div>
          <div class="col-auto col-sm-3 mb-3">
            <form>
              <div class="input-group">
                <input class="form-control form-control-sm shadow-none search" type="search"
                placeholder="Search..." aria-label="search" />
                <div class="input-group-text bg-transparent">
                  <span class="fa fa-search fs--1 text-600"></span>
                </div>
              </div>
            </form>
          </div>
        </div>
        <div class="d-flex justify-content-end" id="allDeleteButton"></div>
        <div class="table-responsive scrollbar">
          <table class="table table-bordered table-hover table-striped fs--1 mb-0">
            <thead class="bg-primary">
              <tr class="text-light">
                ` + Tabledata + `
              </tr>
            </thead>
            <tbody class="list" id="bulk-select-body">
              ` + T.BuildBody() + `
            </tbody>
          </table>
        </div>
        <div class="d-flex justify-content-center mt-3">
        <button class="btn btn-sm btn-falcon-default me-1" type="button" title="Previous"
        data-list-pagination="prev"><span class="fas fa-chevron-left"></span></button>
        <ul class="pagination mb-0"></ul>
        <button class="btn btn-sm btn-falcon-default ms-1" type="button" title="Next"
        data-list-pagination="next"><span class="fas fa-chevron-right"> </span></button>
        </div>
        </div>
        </div>
        </div>
        </div>`)
}

// ------------------------------------------------
// -----Modal to edit or add Employee Details -----
// ------------------------------------------------

type Modals struct {
	ModalID string
	FormId  string
	Title   string
	Action  string
	EmpData m.Employees
	EmpId   string
	Html    string `json:"html"`
}

// Func to build model
func (m *Modals) Build() string {
	var DropDownOptions string
	var EmpIDHtml string

	if m.Action == "Add" {
		m.Title = "Add Employee Details"
		m.ModalID = "AddModal"
		m.FormId = "AddEmpForm"
		DropDownOptions = `<option value="" disabled selected>Select Gender</option>
                <option value="Male" >Male</option>
                <option value="Female">Female</option>
                <option value="Prefer not to say">Prefer not to say</option>`
		EmpIDHtml = ""
	}
	if m.Action == "Edit" {
		m.Title = "Edit Employee Details"
		m.ModalID = "EditModal"
		m.FormId = "EditEmpForm"
		genderOptions := []string{"Male", "Female", "Prefer not to say"}

		for _, gender := range genderOptions {
			selected := ""
			if gender == m.EmpData.Gender {
				selected = "selected"
			}
			DropDownOptions += fmt.Sprintf(`<option value="%s" %s>%s</option>`, gender, selected, gender)
		}

		EmpIDHtml = fmt.Sprint(`<div class="form-group d-flex" id="emp-id">
              <h5>Employee ID : <h5 id="id-for-emp">`, m.EmpData.ID, `</h5></h5>
            </div>`)

	}
	var dateOnlyDOB string
	var dateOnlyDOJ string
	// Split the date strings and extract the date parts
	if m.EmpData.Birthdate != "" && m.EmpData.Joindate != "" {

		dateParts := strings.Split(m.EmpData.Birthdate, "T")
		if len(dateParts) != 2 {
			log.Print("Error splitting DOB:", m.EmpData.Birthdate)
			// Handle error as needed
		}
		dateOnlyDOB = dateParts[0]

		dateParts = strings.Split(m.EmpData.Joindate, "T")
		if len(dateParts) != 2 {
			log.Print("Error splitting DOJ:", m.EmpData.Joindate)
			// Handle error as needed
		}
		dateOnlyDOJ = dateParts[0]
	}

	return u.JoinStr(`
          <!-- Employee Edit Modal -->
<div class="modal fade" id="`, m.ModalID, `" tabindex="-1" role="dialog" aria-hidden="true">
  <div class="modal-dialog modal-dialog-centered" role="document" style="width:50% !">
    <div class="modal-content position-relative">
      <div class="modal-body p-0">
        <div class="rounded-top-lg py-3 ps-4 pe-6 bg-light">
          <h4 class="mb-1" id="ModalHead">`, m.Title, `</h4>
        </div>
        <div class="p-4 pb-0">
          <form id="`, m.FormId, `">
            `, EmpIDHtml, `
            <div class="form-group">
              <label for="emp-name">Name:</label>
              <input type="text" id="emp-name" name="emp-name" class="form-control" value="`, m.EmpData.Name, `" required>
            </div>
            <div class="row gx-2">
              <div class="form-group mb-3 col-sm-6">
                <label for="emp-dep">Department:</label>
                <input type="text" id="emp-dep" name="emp-dep" class="form-control" value="`, m.EmpData.Department, `"required>
              </div>
              <div class="form-group mb-3 col-sm-6">
                <label for="emp-deg">Designation:</label>
                <input type="text" id="emp-deg" name="emp-deg" class="form-control" value="`, m.EmpData.Designation, `" required>
              </div>
            </div>
            <div class="row gx-2">
              <div class="form-group mb-3 col-sm-6">
                <label for="emp-mob">Mobile No.:</label>
                <input type="number" id="emp-mob" name="emp-mob" class="form-control" minlength="8" maxlength="15" value="`, m.EmpData.Mobile, `">
              </div>
              <div class="form-group mb-3 col-sm-6">
                <label for="emp-dob">Date of Birth:</label>
                <input type="date" id="emp-dob" name="emp-dob" class="form-control"  value="`, dateOnlyDOB, `"required>
              </div>
            </div>
            <div class="form-group">
              <label for="emp-email">Email:</label>
              <input type="email" id="emp-email" name="emp-email" class="form-control" value="`, m.EmpData.Email, `" required>
            </div>
            <div class="form-group">
              <label for="emp-add">Address:</label>
              <input type="text" id="emp-add" name="emp-add" class="form-control" value="`, m.EmpData.Address, `" required>
            </div>
            <div class="row gx-2">
              <div class="form-group form-group mb-3 col-sm-6">
                <label for="emp-salary">Salary:</label>
                <input type="number" id="emp-salary" name="emp-salary" class="form-control" value="`, m.EmpData.Salary, `" required>
              </div>
              <div class="form-group form-group mb-3 col-sm-6">
                <label for="emp-join">Date of Joining:</label>
                <input type="date" id="emp-join" name="emp-join" class="form-control" value="`, dateOnlyDOJ, `" required>
              </div>
            </div>
            <div class="form-group mb-3 col-sm-6">
              <label for="emp-gender">Gender:</label>
              <select id="emp-gender" name="emp-gender" class="form-control w-100" required>
                  `, DropDownOptions, `
              </select>
            </div>
          </form>
        </div>
      </div>
      <div class="modal-footer mt-1 d-flex justify-content-end">
        <button type="button" class="btn btn-secondary ml-2" id="closeBtn">Close</button>
        <button type="submit" id="submit-form" class="btn btn-primary">`, m.Action, `</button>
      </div>
    </div>
  </div>
</div>
  `)
}
