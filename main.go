package main

import (
	"fmt"
	"log"
	"net/http"
)

func getDataFromThirdPartyHTTPAPI(url string) (*http.Response, int, error) {
	var resp, _ = http.Get(url)

	if resp.StatusCode == http.StatusNotFound {
		return nil, http.StatusNotFound, fmt.Errorf("third party HTTP API responded with 404 StatusNotFound")
	}
	return resp, http.StatusOK, nil
}

func main() {
	var productionURL = "https://www.realproductionURL.com"

	var _, statusCode, err = getDataFromThirdPartyHTTPAPI(productionURL)

	if err != nil {
		log.Printf("Error occurred: %v. Status code: %d", err, statusCode)
	}else {
		log.Printf("Call to third party HTTP API was successful")
	}
}
