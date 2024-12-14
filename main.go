package main

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587                     // Use 465 if 587 doesn't work
	username := "akmobile.tm@gmail.com" // Your Gmail address
	password := "whclvwobghfdrmqm"      // Your Gmail app password (16 characters)

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", username)                                                               // Sender email address
	m.SetHeader("To", "multimillioner2030@gmail.com")                                           // Receiver email address
	m.SetHeader("Subject", "Gomail Test Email")                                                 // Email subject
	m.SetBody("text/plain", "This is a test email sent using Gomail in plain text.")            // Email body
	m.AddAlternative("text/html", "<h1>Hello!</h1><p>This is a test email in HTML format.</p>") // Optional HTML body

	// Set up the SMTP dialer
	d := gomail.NewDialer(smtpHost, smtpPort, username, password)

	// Optional: In case of certificate issues, skip verification (use only in testing)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email and handle errors
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
