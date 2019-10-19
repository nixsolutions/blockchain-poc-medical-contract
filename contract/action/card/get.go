package card

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

// Get returns the value of the specified asset key
func GetCard(stub shim.ChaincodeStubInterface, args []string) (string, error) {
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
	if user.IsParent() || user.IsPediatrician() {
		bytes, err := json.Marshal(card)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	}
	access, err := service.NewAccessService(stub).FindAccessByDoctor(user.Id, args[0])
	if err != nil {
		return "", err
	}

	return scopedResponse(card, access.Fields)
}

func scopedResponse(card model.Card, fields []string) (string, error) {
	responseMap := make(map[string]interface{})
	cardMap := card.ToMap()
	for _, value := range fields {
		cardValue, ok := cardMap[value]
		if ok {
			responseMap[value] = cardValue
		}
	}

	bytes, err := json.Marshal(responseMap)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
