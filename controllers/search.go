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

func extractParamsSearchRequest(r *http.Request) (string, string, int, error) {
	vars := mux.Vars(r)
	query := helpers.NormalizeUnicode(vars["query"])
	method := vars["method"]
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		return "", "", 0, err
	} else if limit <= 0 {
		return "", "", 0, fmt.Errorf("Invalid limit")
	} else if query == "" {
		return "", "", 0, fmt.Errorf("Invalid query")
	} else if method == "" {
		return "", "", 0, fmt.Errorf("Invalid method")
	}
	return query, method, limit, nil
}

func (c *SearchController) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query, method, limit, err := extractParamsSearchRequest(r)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	searcher := searchers.NewSearcher(method, constants.DBType)
	candidates, err := searcher.Search(query, limit)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	helpers.SendResponse(w, http.StatusOK, "OK", candidates)
}
