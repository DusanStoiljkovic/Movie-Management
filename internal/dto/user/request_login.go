package dto

type RequestLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
