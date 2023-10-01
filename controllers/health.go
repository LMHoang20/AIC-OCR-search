package controllers

import (
	"OCRsearch/helpers"
	"net/http"
)

type Health struct{}

var health *Health

func HealthInstance() *Health {
	if health == nil {
		health = &Health{}
	}
	return health
}

func (c *Health) HealthHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendResponse(w, http.StatusOK, "OK", nil)
}
