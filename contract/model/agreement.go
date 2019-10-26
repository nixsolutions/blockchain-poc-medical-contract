package model

const TO_SIGN_STATUS = "TO_SIGH"
const SIGNED_STATUS = "SIGNED"

type Agreement struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Doctor    string `json:"doctor"`
	Parent    string `json:"parent"`
	Timestamp int64  `json:"timestamp"`
}
