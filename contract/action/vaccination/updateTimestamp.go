package vaccination

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
	"strconv"
)

func UpdateVaccinationTimestamp(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a value")
	}
	cardId, vaccinationName, timestampString := args[0], args[1], args[2]

	var card model.Card
	cardService := service.NewCardService(stub)
	err := cardService.FindAndUnmarshal(cardId, &card)
	if err != nil {
		return "", err
	}

	for key, item := range card.Vaccination {
		if item.Name == vaccinationName {
			timestamp, err := strconv.ParseInt(timestampString, 10, 64)
			if err != nil {
				return "", err
			}
			card.Vaccination[key].Timestamp = timestamp
			break
		}
	}

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