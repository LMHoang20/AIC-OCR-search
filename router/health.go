package router

import (
	"OCRsearch/controllers"

	"github.com/gorilla/mux"
)

func initHealth(r *mux.Router) {
	r.HandleFunc("/health", controllers.HealthInstance().HealthHandler).Methods("GET")
}
