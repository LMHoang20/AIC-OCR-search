package router

import (
	"OCRsearch/controllers"

	"github.com/gorilla/mux"
)

func initSearch(r *mux.Router) {
	r.HandleFunc("/search/{method}", controllers.SearchInstance().SearchHandler).Methods("POST")
}
