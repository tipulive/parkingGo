package controller

import (
	"fmt"
	"net/http"
	"parking/helper"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {

	cfg := helper.SMTPConfig{
		Host:     "server162.web-hosting.com",
		Port:     465,
		Username: "appinfo@appdev.live",
		Password: "Ca![R$j7ndcd",
		From:     "appinfo@appdev.live",
	}

	email := helper.Email{
		To:      []string{"ericsoft123@gmail.com"},
		Subject: "Welcome to Our App!",
		Body:    "<h1>Hello!</h1><p>Thanks for signing up.</p>",
		HTML:    true,
	}

	if err := helper.SendEmail(cfg, email); err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
