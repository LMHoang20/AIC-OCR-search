package controllers

import (
	"OCRsearch/helpers"
	"OCRsearch/searchers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SearchController is a struct that contains the SearchHandler function
type SearchController struct{}

var search *SearchController

// SearchInstance is a function that returns a pointer to a SearchController
func SearchInstance() *SearchController {
	if search == nil {
		search = &SearchController{}
	}
	return search
}

func (c *SearchController) SearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid limit", nil)
		return
	}
	candidates, err := searchers.NewExact("RAM").Search(query, limit)
	helpers.SendResponse(w, http.StatusOK, "OK", candidates)
}
