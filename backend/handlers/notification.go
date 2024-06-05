package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NotificationTime struct {
	Time string `json:"time"`
}

func SetNotificationTime(w http.ResponseWriter, r *http.Request) {
	var notificationTime NotificationTime
	_ = json.NewDecoder(r.Body).Decode(&notificationTime)
	vars := mux.Vars(r)
	userID := vars["userId"]
	log.Printf("User %s set notification time to %s", userID, notificationTime.Time)
	w.WriteHeader(http.StatusOK)
}
