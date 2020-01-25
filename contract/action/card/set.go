package card

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func Create(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a value")
	}
	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}
	if !user.IsParent() {
		return "", errors.New("user is not a parent")
	}

	var card model.Card
	err = json.Unmarshal([]byte(args[1]), &card)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshall arg", args[0])
	}

	cardJsonBytes, err := json.Marshal(card)
	if err != nil {
		return "", fmt.Errorf("Failed to marshall card", args[0])
	}

	err = stub.PutState("CARD" + args[0], cardJsonBytes)
	if err != nil {
		return "", fmt.Errorf("Failed to set card: %s", args[0])
	}

	return args[1], nil
}