package app

//API_service interface
type API_service interface {
	GetAffectedCountriesClient() ([]byte, error)
}
