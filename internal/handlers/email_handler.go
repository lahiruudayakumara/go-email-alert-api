package handlers

import (
	"email-alert-api/internal/email"
	"encoding/json"
	"net/http"
)

// EmailHandler handles sending emails
func EmailHandler(w http.ResponseWriter, r *http.Request) {
	var emailReq email.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&emailReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := email.SendEmail(emailReq)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully!"))
}
