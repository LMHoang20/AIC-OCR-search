package router

import (
	"OCRsearch/controllers"

	"github.com/gorilla/mux"
)

func initFrame(r *mux.Router) {
	r.HandleFunc("/frame/{video}", controllers.FrameInstance().FrameHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/map-frame/{video}/{start}/{end}", controllers.FrameInstance().FrameHandler2).Methods("GET", "OPTIONS")
}
