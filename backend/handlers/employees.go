package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ZemtsovMaxim/RuTube_TestTask/backend/models"
)

var employees = []models.Employee{
	{ID: 1, Name: "John Doe", BirthDate: "1990-03-25", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", BirthDate: "1985-06-10", Email: "jane@example.com"},
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
