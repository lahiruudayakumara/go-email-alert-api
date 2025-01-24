package email

import (
	"email-alert-api/config"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// EmailRequest represents the email payload
type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// SendEmail sends an email using EmailJS
func SendEmail(emailReq EmailRequest) error {
	client := resty.New()

	// Prepare the payload for the EmailJS API
	var payload = map[string]interface{}{
		"service_id":  config.GetServiceID(),
		"template_id": config.GetTemplateID(),
		"user_id":     config.GetEmailJSAPIKey(),
		"accessToken": config.GetEmailJSPrivateKey(),
		"template_params": map[string]string{
			"email":   emailReq.To,
			"subject": emailReq.Subject,
			"message": emailReq.Message,
		},
	} // Send the POST request
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(config.GetEmailJSURL())

	// Error handling: check if there's an issue with the request
	if err != nil {
		log.Printf("Email sending error: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	// Check the response status code
	if resp.StatusCode() != http.StatusOK {
		// Log the response for debugging purposes
		log.Printf("Failed to send email, status: %d, response: %s", resp.StatusCode(), resp.String())
		return errors.New("failed to send email, unexpected response from server")
	}

	// Success case, no error
	return nil
}
