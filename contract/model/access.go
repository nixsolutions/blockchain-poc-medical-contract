package model

const ACCESS_STATUS_VALID = "valid"
const ACCESS_STATUS_INVALID = "invalid"

type Access struct {
	Doctor string `json:"doctor"`
	GivenBy string `json:"given_by"`
	Status string `json:"status"`
	Fields []string `json:"fields"`
	Timestamp int64 `json:"timestamp"`
	Card string `json:"card"`
}