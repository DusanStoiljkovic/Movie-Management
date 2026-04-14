package dto

type ResponseLoginUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
