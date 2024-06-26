package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ZemtsovMaxim/RuTube_TestTask/backend/models"
	"github.com/ZemtsovMaxim/RuTube_TestTask/backend/notifications"

	"github.com/ZemtsovMaxim/RuTube_TestTask/backend/handlers"

	"github.com/gorilla/mux"
)

var users = []models.User{
	{ID: 1, Username: "user1", Password: "password", Email: "user1@example.com"},
	{ID: 2, Username: "user2", Password: "password", Email: "user2@example.com"},
}

var subscriptions = []models.Subscription{
	{UserID: 1, EmployeeID: 1},
	{UserID: 2, EmployeeID: 2},
}

var employees = []models.Employee{
	{ID: 1, Name: "John Doe", BirthDate: "1990-03-25", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", BirthDate: "1985-06-10", Email: "jane@example.com"},
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")
	router.HandleFunc("/subscribe", handlers.Subscribe).Methods("POST")
	router.HandleFunc("/unsubscribe", handlers.Unsubscribe).Methods("POST")
	router.HandleFunc("/users/{userId}/notification-time", handlers.SetNotificationTime).Methods("POST")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/public/"))))

	// Запуск сервера
	go func() {
		log.Println("Server started at :8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// Периодическая проверка дней рождения
	ticker := time.NewTicker(24 * time.Hour)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				checkBirthdays()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Блокировка основной горутины
	select {}
}

func checkBirthdays() {
	today := time.Now().Format("2006-01-02")
	for _, employee := range employees {
		if employee.BirthDate == today {
			for _, subscription := range subscriptions {
				if subscription.EmployeeID == employee.ID {
					for _, user := range users {
						if user.ID == subscription.UserID {
							notifications.SendBirthdayNotification(user, employee)
						}
					}
				}
			}
		}
	}
}
