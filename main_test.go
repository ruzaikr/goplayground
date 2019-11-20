package main

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func respondWith404(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func TestGetDataFromThirdPartyHTTPAPI(t *testing.T) {

	var testServer = httptest.NewServer(http.HandlerFunc(respondWith404))

	var _, statusCode, err =  getDataFromThirdPartyHTTPAPI(testServer.URL)

	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, fmt.Errorf("third party HTTP API responded with 404 StatusNotFound"), err)
}