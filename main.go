package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract"
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
	//	Parent:     "parent1",
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
	//	Parent: "parent1",
	//	Timestamp: 1571590425,
	//}
	//
	//bytes, err = json.Marshal(agreement)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(bytes))

	if err := shim.Start(new(contract.MedicalContract)); err != nil {
		fmt.Printf("Error starting MedicalContract chaincode: %s", err)
	}
}
