package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract"
)

// main function starts up the chaincode in the container during instantiate
func main() {
	//vac := []model.VaccinationItem{
	//	model.VaccinationItem{
	//		Name:      "test1",
	//		Timestamp: 1111,
	//	},
	//	model.VaccinationItem{
	//		Name:      "test2",
	//		Timestamp: 2222,
	//	},
	//}
	//card := model.Card{
	//	Name:        "Max Pechenin",
	//	Type:        "card",
	//	BirthDate:   "1999-05-24",
	//	Height:      197,
	//	Weight:      85,
	//	Vaccination: vac,
	//	Parents:     []model.Parent{{Id: "parent-max-1"}, {Id: "parent-max-2"}},
	//}
	//bytes, err := json.Marshal(card)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(bytes))

	if err := shim.Start(new(contract.MedicalContract)); err != nil {
		fmt.Printf("Error starting MedicalContract chaincode: %s", err)
	}
}
