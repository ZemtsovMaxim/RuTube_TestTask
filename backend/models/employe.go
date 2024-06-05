package models

type Employee struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Email     string `json:"email"`
}
