package models

type Log struct {
	ID        string `json:"id"`
	Created   string `json:"created"`
	EmpName   string `json:"empname"`
	EmpEmail  string `json:"empemail"`
	EmpDeg    string `json:"empdeg"`
	Operation string `json:"operation"`
	UpdatedBy string `json:"updatedby"`
}
