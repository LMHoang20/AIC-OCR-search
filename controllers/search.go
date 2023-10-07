package controllers

import (
	"OCRsearch/constants"
	"OCRsearch/helpers"
	"OCRsearch/searchers"
	"encoding/json"
	"net/http"

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
	type Body struct {
		QueryText string `json:"query_text"`
		Topk      int    `json:"topk"`
	}
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return "", "", 0, err
	}
	method := vars["method"]
	query := body.QueryText
	limit := body.Topk
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
