package vaccination

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

func DeleteVaccination(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a value")
	}

	var card model.Card
	cardService := service.NewCardService(stub)
	err := cardService.FindAndUnmarshal(args[0], &card)
	if err != nil {
		return "", err
	}
	accessService := service.NewAccessService(stub)
	user := service.NewAuthService(stub).GetUser()
	if user.IsNeuropathologist() {
		return "", errors.New("only Neuropathologist can do vaccination")
	}

	access, err := accessService.FindAccessByDoctor(user.Id, args[0])
	if access == nil {
		return  "", errors.New("access was not found")
	}

	if access.Invalid() {
		return  "", errors.New("access is invalid")
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
