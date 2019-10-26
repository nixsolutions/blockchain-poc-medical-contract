package card

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

// Get returns the value of the specified asset key
func GetAll(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	cardService := service.NewCardService(stub)
	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}

	var response interface{}
	if user.IsParent() {
		response, err = cardService.FindCardsByParent(user.Id)
	}

	if user.IsHospitalWorker() {
		response, err = GetAllForHospitalWorker(user, stub)
	}

	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func GetAllForHospitalWorker(user *model.User, stub shim.ChaincodeStubInterface) ([]model.Card, error) {
	agreements, err := service.NewAgreementService(stub).FindAgreementsByDoctor(user.Id)
	cardService := service.NewCardService(stub)

	if err != nil {
		return nil, err
	}

	var cards []model.Card

	for _, agreement := range agreements {
		parentCards, err := cardService.FindCardsByParent(agreement.Parent)
		if err != nil {
			return nil, err
		}
		cards = append(cards, parentCards...)
	}

	return cards, nil
}
