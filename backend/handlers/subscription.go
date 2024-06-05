package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Subscription struct {
	UserID     int `json:"userId"`
	EmployeeID int `json:"employeeId"`
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	var subscription Subscription
	_ = json.NewDecoder(r.Body).Decode(&subscription)
	log.Printf("User %d subscribed to employee %d", subscription.UserID, subscription.EmployeeID)
	w.WriteHeader(http.StatusOK)
}

func Unsubscribe(w http.ResponseWriter, r *http.Request) {
	var subscription Subscription
	_ = json.NewDecoder(r.Body).Decode(&subscription)
	log.Printf("User %d unsubscribed from employee %d", subscription.UserID, subscription.EmployeeID)
	w.WriteHeader(http.StatusOK)
}
