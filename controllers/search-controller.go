package controllers

import (
	"net/http"
	"net/url"

	"github.com/gonzalezlrjesus/search-api/models"
	"github.com/gonzalezlrjesus/search-api/services"
	"github.com/gonzalezlrjesus/search-api/utils"
)

// SearchController .
var SearchController = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()
	var responseList []models.APIResponse

	responseList = append(responseList, utils.AppleStructToResponseStruct(services.SearchApple(urlValues.Get("query"), "song"))...)
	responseList = append(responseList, utils.AppleStructToResponseStruct(services.SearchApple(urlValues.Get("query"), "movie"))...)
	responseList = append(responseList, utils.AppleStructToResponseStruct(services.SearchApple(urlValues.Get("query"), "ebook"))...)

	response :=
		map[string]interface{}{
			"query": responseList,
		}
	utils.Respond(w, response, 200)
}
