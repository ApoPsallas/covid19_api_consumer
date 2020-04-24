package app

//APIService interface
type APIService interface {
	GetAffectedCountries() ([]byte, error)
}
