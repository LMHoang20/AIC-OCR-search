package controllers

import (
	"OCRsearch/constants"
	"OCRsearch/helpers"
	"OCRsearch/searchers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SearchController struct{}

var search *SearchController

func SearchInstance() *SearchController {
	if search == nil {
		search = &SearchController{}
	}
	return search
}

func (c *SearchController) SearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := helpers.NormalizeUnicode(vars["query"])
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid limit", nil)
		return
	} else if limit <= 0 {
		helpers.SendResponse(w, http.StatusBadRequest, "Limit must be greater than 0", nil)
		return
	} else if query == "" {
		helpers.SendResponse(w, http.StatusBadRequest, "Invalid query", nil)
		return
	}
	fmt.Println("Query: ")
	for _, c := range query {
		fmt.Printf("%c\n", c)
	}
	fmt.Println("Limit: ", limit)
	candidates, err := searchers.NewExact(constants.DBType).Search(query, limit)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "OK", candidates)
}
