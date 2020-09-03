package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gonzalezlrjesus/search-api/models"
)

// SearchTvMaze .
func SearchTvMaze(query string) models.TvMazeAPI {

	response, err := http.Get("http://api.tvmaze.com/search/shows?q=" + query)

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

	var responseObject models.TvMazeAPI
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Println("Unmarshal")
		log.Fatal(err)
	}

	return responseObject
}
