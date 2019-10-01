package vaccination

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/repository"
)

func DeleteVaccination(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a value")
	}

	var card model.Card
	cardRepository := &repository.CardRepository{Stub: stub}
	err := cardRepository.Find(args[0], &card)
	if err != nil {
		return "", err
	}

	card.Vaccination = removeVaccination(card.Vaccination, args[1])

	cardJsonBytes, err := json.Marshal(card)
	if err != nil {
		return "", fmt.Errorf("Failed to marshall card", args[0])
	}

	err = stub.PutState(args[0], cardJsonBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to set card: %s", args[0])
	}

	return string(cardJsonBytes), nil
}

func removeVaccination(vaccinationItems []model.VaccinationItem, vaccinationName string) []model.VaccinationItem {
	var newVaccinationItems []model.VaccinationItem
	for _, value := range vaccinationItems {
		if value.Name == vaccinationName {
			continue
		}
		newVaccinationItems = append(newVaccinationItems, value)
	}
	return newVaccinationItems
}
