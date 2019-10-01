package repository

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
)

type CardRepository struct {
	Stub shim.ChaincodeStubInterface
}

func (cr  *CardRepository) Find(key string, card *model.Card) error {
	cardBytes, err := cr.Stub.GetState(key)
	if err != nil {
		return  fmt.Errorf("Failed to get card: %s with error: %s", key, err)
	}
	if cardBytes == nil {
		return  fmt.Errorf("Card was not found: %s", key)
	}

	err = json.Unmarshal([]byte(cardBytes), card)
	if err != nil {
		return  fmt.Errorf("Failed to unmarshall card obj", key)
	}

	return nil
}