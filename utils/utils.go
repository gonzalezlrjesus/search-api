package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gonzalezlrjesus/search-api/models"
)

// Respond function to responde request
func Respond(w http.ResponseWriter, data map[string]interface{}, status uint) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("X-Frame-Options", "Deny")
	w.Header().Add("Content-Security-Policy", "script-src 'self'")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("Referrer-Policy", "no-referrer")
	w.Header().Add("Feature-Policy", "vibrate 'none'; geolocation 'none'")

	switch status {
	case 200: // break 200 Accept Request
		w.WriteHeader(http.StatusOK)
		break
	case 201: // break 201 created POST
		w.WriteHeader(http.StatusCreated)
		break
	case 204: // break 204 No Content (Just Delete Http)
		w.WriteHeader(http.StatusNoContent)
		break
	case 301: // break 301 Moved Permanently
		w.WriteHeader(http.StatusMovedPermanently)
		break
	case 400: // break 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		break
	case 401: // break 401 Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		break
	case 403: // break 403 Forbidden
		w.WriteHeader(http.StatusForbidden)
		break
	case 404: // break 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		break
	case 500: // break 500 Internal Server Error
		w.WriteHeader(http.StatusInternalServerError)
		break
	}

	data["status"] = status
	json.NewEncoder(w).Encode(data)
}

// AppleStructToResponseStruct .
func AppleStructToResponseStruct(appleList models.AppleAPI) []models.APIResponse {
	var responseList []models.APIResponse

	for _, b := range appleList.Results {
		newAPIResponse := models.APIResponse{Name: b.TrackName, Kind: b.Kind, Date: b.ReleaseDate, Originapi: "APPLE"} // pulled out for clarity
		responseList = append(responseList, newAPIResponse)
	}
	// log.Println(responseList)
	return responseList
}

// TvMazeStructToResponseStruct .
func TvMazeStructToResponseStruct(showsList models.TvMazeAPI) []models.APIResponse {
	var responseList []models.APIResponse

	for _, b := range showsList {
		newAPIResponse := models.APIResponse{Name: b.Show.Name, Kind: "show", Date: b.Show.Premiered, Originapi: "TVMAZE"} // pulled out for clarity
		responseList = append(responseList, newAPIResponse)
	}

	return responseList
}

// CrcindStructToResponseStruct .
func CrcindStructToResponseStruct(personList models.CrcindAPI) []models.APIResponse {
	var responseList []models.APIResponse

	for _, b := range personList.Body.GetListByNameResponse.GetListByNameResult.PersonIdentification {
		newAPIResponse := models.APIResponse{Name: b.Name, Kind: "person", Date: b.DOB, Originapi: "CRCIND"} // pulled out for clarity
		responseList = append(responseList, newAPIResponse)
	}

	return responseList
}

// QuickSortToSortbyName .
func QuickSortToSortbyName(responseList []models.APIResponse) []models.APIResponse {

	if len(responseList) > 1 {
		pivotIndex := len(responseList) / 2
		var smallerItems = []models.APIResponse{}
		var largerItems = []models.APIResponse{}

		for i := range responseList {
			val := responseList[i].Name
			if i != pivotIndex {
				if val < responseList[pivotIndex].Name {
					smallerItems = append(smallerItems, responseList[i])
				} else {
					largerItems = append(largerItems, responseList[i])
				}
			}
		}

		QuickSortToSortbyName(smallerItems)
		QuickSortToSortbyName(largerItems)

		var merged []models.APIResponse = append(append(append([]models.APIResponse{}, smallerItems...), []models.APIResponse{responseList[pivotIndex]}...), largerItems...)

		for j := 0; j < len(responseList); j++ {
			responseList[j] = merged[j]
		}

	}
	return responseList
}
