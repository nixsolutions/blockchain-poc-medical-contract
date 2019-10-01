package model

const AGREEMENT_CREATED_STATUS  = "valid"
const AGREEMENT_ACTIVE_STATUS  = "invalid"

type Agreement struct {
	Status string `json:"status"`
	Doctor string `json:"doctor"`
	Parents []string `json:"parents"`
	SignedByParent bool `json:"signed_by_parent"`
	SignedByDoctor bool `json:"signed_by_doctor"`
}