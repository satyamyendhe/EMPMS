package models

type Employees struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Joindate    string `json:"joindate"`
	Birthdate   string `json:"birthdate"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Mobile      string `json:"mobile"` // Use string to check length==10
	Salary      string `json:"salary"`
}
