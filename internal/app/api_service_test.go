package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAffectedCountries(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := API{Client: server.Client()}
	resp, err := api.GetAffectedCountriesClient()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
