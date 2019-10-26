package main

import (
	"encoding/json"
	"fmt"
	"poc/contract/model"
)

// main function starts up the chaincode in the container during instantiate
func main() {
	//vac := []model.VaccinationItem{
	//	{
	//		Name:      "measles",
	//		Timestamp: 1571590625,
	//	},
	//}
	//card := model.Card{
	//	Name:        "John Doe",
	//	Type:        "card",
	//	BirthDate:   "2007-05-24",
	//	Height:      197,
	//	Weight:      85,
	//	Vaccination: vac,
	//	Parents:     []model.Parent{{Id: "parent-1-uid"}, {Id: "parent-2-uid"}},
	//}
	//bytes, err := json.Marshal(card)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(bytes))
	//agreement := model.Agreement{
	//	Type: "agreement",
	//	Status: "TO_SIGN",
	//	Doctor: "doctor-uniq-id",
	//	Parents: []string{ "parent-1-uid",  "parent-2-uid"},
	//	Timestamp: 1571590425,
	//}

	access := model.Access{
		Type: "access",
		Doctor: "neuro-uniq-id",
		GivenBy: "doctor-uniq-id",
		Status: "valid",
		Fields: []string{"name", "vaccination"},
		Timestamp: 1571590525,
		Card: "card-uniq-id",
	}

	bytes, err := json.Marshal(access)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	//if err := shim.Start(new(contract.MedicalContract)); err != nil {
	//	fmt.Printf("Error starting MedicalContract chaincode: %s", err)
	//}
}
