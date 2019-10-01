package model

type Access struct {
	Doctor string `json:"doctor"`
	GivenBy string `json:"given_by"`
	Status string `json:"status"`
	Fields []string
}