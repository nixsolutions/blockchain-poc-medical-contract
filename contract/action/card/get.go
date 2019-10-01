package card

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/repository"
)

// Get returns the value of the specified asset key
func GetCard(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	var card model.Card
	cardRepository := &repository.CardRepository{Stub: stub}
	err := cardRepository.Find("CARD" + args[0], &card)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(card)
	if err != nil {
		return  "", fmt.Errorf("Failed to marshall card obj", args[0])
	}

	return string(jsonBytes), nil
}
