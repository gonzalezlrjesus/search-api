package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gonzalezlrjesus/search-api/models"
)

// SearchApple .
func SearchApple(query string, mediaType string) models.AppleAPI {

	query = removeAndReplaceWhiteSpaces(query)

	response, err := http.Get("https://itunes.apple.com/search?term=" + query + "&entity=" + mediaType + "&limit=5")

	if err != nil {
		log.Println("GET")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ReadAll")

		log.Fatal(err)
	}

	var responseObject models.AppleAPI
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Println("Unmarshal")
		log.Fatal(err)
	}

	return responseObject
}

func removeAndReplaceWhiteSpaces(query string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(strings.TrimSpace(query), "+")
}
