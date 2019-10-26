package card

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

// Get returns the value of the specified asset key
func Get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) < 1 {
		return "", errors.New("Incorrect number of arguments. Expecting 1")
	}
	cardService := service.NewCardService(stub)
	hasAccess, err := cardService.HasAccessToCard(args[0])
	if !hasAccess {
		return "", errors.New("no access to card")
	}
	if err != nil {
		return "", err
	}

	var card model.Card
	err = cardService.FindAndUnmarshal(args[0], &card)
	if err != nil {
		return "", err
	}

	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}

	if !user.IsParent() || !user.IsHospitalWorker() {
		return "", errors.New("you are not allowed to get the card")
	}

	bytes, err := json.Marshal(card)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
