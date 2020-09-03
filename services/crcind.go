package services

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gonzalezlrjesus/search-api/models"
)

// SearchCrcind .
func SearchCrcind(query string) models.CrcindAPI {

	response, err := http.Get("http://www.crcind.com/csp/samples/SOAP.Demo.cls?soap_method=GetListByName&name=" + query)

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

	var responseObject models.CrcindAPI
	err = xml.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseObject
	}

	return responseObject
}
