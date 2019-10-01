package model

type Card struct {
	Name string `json:"name"`
	BirthDate string `json:"birth_date"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Vaccination []VaccinationItem `json:"vaccination"`
	Parents []string `json:"parents"`
}