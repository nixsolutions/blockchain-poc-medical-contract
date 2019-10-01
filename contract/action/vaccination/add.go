package vaccination

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/repository"
)

func AddVaccination(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a value")
	}


	var card model.Card
	cardRepository := &repository.CardRepository{ Stub: stub}
	err := cardRepository.Find(args[0], &card)
	if err != nil {
		return "", err
	}

	var vaccItem model.VaccinationItem
	err = json.Unmarshal([]byte(args[1]), &vaccItem)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshall vacc item", args[1])
	}

	card.Vaccination = append(card.Vaccination, vaccItem)

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