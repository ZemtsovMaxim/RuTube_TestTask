package notifications

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/ZemtsovMaxim/RuTube_TestTask/backend/models"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

var emailConfig = EmailConfig{
	SMTPHost:     "smtp.example.com",
	SMTPPort:     "587",
	SMTPUsername: "your-email@example.com",
	SMTPPassword: "your-email-password",
}

func SendBirthdayNotification(subscriber models.User, employee models.Employee) {
	from := emailConfig.SMTPUsername
	password := emailConfig.SMTPPassword
	to := subscriber.Email
	subject := fmt.Sprintf("Happy Birthday %s!", employee.Name)
	body := fmt.Sprintf("Dear %s,\n\nToday is %s's birthday! Don't forget to wish them a happy birthday.\n\nBest regards,\nBirthday Notifier", subscriber.Username, employee.Name)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, emailConfig.SMTPHost)

	err := smtp.SendMail(emailConfig.SMTPHost+":"+emailConfig.SMTPPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	} else {
		log.Printf("Email sent successfully to %s", to)
	}
}
