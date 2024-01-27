package model

type NewUsers struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
