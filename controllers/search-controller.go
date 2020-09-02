package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gonzalezlrjesus/search-api/utils"
)

// SearchController .
var SearchController = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()

	log.Println(urlValues.Get("query"))
	response :=
		map[string]interface{}{
			"query": urlValues.Get("query"),
		}
	utils.Respond(w, response, 200)
}
