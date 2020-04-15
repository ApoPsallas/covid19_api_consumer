package app

//AffectedCountries is a struct to unmarshal the API response   
type AffectedCountries struct {
	AffectedCountries []string `json:"affected_countries"`
	Timestamp         string   `json:"statistic_taken_at"`
}

//NewAffectedCountries constructor
func NewAffectedCountries(countries []string, timestamp string) *AffectedCountries {
	return &AffectedCountries{countries, timestamp}
}
