package router

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	initHealth(r)
	initSearch(r)
	return r
}
