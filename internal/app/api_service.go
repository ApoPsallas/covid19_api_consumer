package app

//APIService interface
type APIService interface {
	GetAffectedCountries() (*AffectedCountries, error)
}
