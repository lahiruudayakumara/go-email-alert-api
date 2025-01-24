package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}

// GetServiceID retrieves the EmailJS service ID from the environment
func GetServiceID() string {
	return os.Getenv("SERVICE_ID") // Retrieve the service ID from the .env file or environment variables
}

// GetTemplateID retrieves the EmailJS template ID from the environment
func GetTemplateID() string {
	return os.Getenv("TEMPLATE_ID")
}

// GetEmailJSAPIKey retrieves the EmailJS API key from the environment
func GetEmailJSAPIKey() string {
	return os.Getenv("EMAILJS_API_KEY")
}

func GetEmailJSPrivateKey() string {
	return os.Getenv("EMAILJS_PRIVATE_KEY")
}

// GetEmailJSURL retrieves the EmailJS API URL from the environment
func GetEmailJSURL() string {
	return os.Getenv("EMAILJS_URL")
}
