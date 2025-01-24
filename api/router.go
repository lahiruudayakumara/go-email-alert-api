package api

import (
	"email-alert-api/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/send-email", handlers.EmailHandler).Methods("POST")

	// Apply middlewares
	//router.Use(middleware.LoggerMiddleware)
	//router.Use(middleware.CORSMiddleware)
	//router.Use(middleware.RateLimitMiddleware)

	return router
}
