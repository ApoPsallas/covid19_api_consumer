package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAffectedCountries(t *testing.T) {

	countries := []string{"gr", "it"}
	timestamp := "now"
	expected := &AffectedCountries{countries, timestamp}
	actual := NewAffectedCountries(countries, timestamp)

	assert.Equal(t, expected, actual)
}
