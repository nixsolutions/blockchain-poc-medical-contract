package model

const TO_SIGN_STATUS  = "TO_SIGH"
const SIGNED_STATUS  = "SIGNED"

type Agreement struct {
	Status string `json:"status"`
	Doctor string `json:"doctor"`
	Parents []string `json:"parents"`
	Timestamp int64 `json:"timestamp"`
}