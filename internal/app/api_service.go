package app

//GetAffectedCountries will send a HTTP request
func GetAffectedCountries() (*AffectedCountries, error) {
	return NewAffectedCountries([]string{"gr", "uk", "es", "it"}, "now"), nil
}
