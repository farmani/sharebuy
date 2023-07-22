package dto

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Activated bool   `json:"activated"`
	Version   int    `json:"-"`
}
