package models

type GetLogin struct {
	// IsAdmin  bool   `json:"isadmin"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
