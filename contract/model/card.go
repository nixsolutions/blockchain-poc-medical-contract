package model

type Card struct {
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	BirthDate   string            `json:"birth_date"`
	Height      int               `json:"height"`
	Weight      int               `json:"weight"`
	Vaccination []VaccinationItem `json:"vaccination"`
	Parents     []Parent          `json:"parents"`
}

type VaccinationItem struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}

func (card Card) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["name"] = card.Name
	m["birth_date"] = card.BirthDate
	m["height"] = card.Height
	m["weight"] = card.Weight
	m["vaccination"] = card.Vaccination
	m["parents"] = card.Parents

	return m
}
