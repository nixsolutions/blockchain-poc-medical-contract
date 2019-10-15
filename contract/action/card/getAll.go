package card

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

// Get returns the value of the specified asset key
func GetCards(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	cardService := service.NewCardService(stub)
	user := service.NewAuthService(stub).GetUser()

	if user.IsParent() {
		queryString := fmt.Sprintf("{\"selector\":{\"type\":\"card\",\"parents\":{\"$elemMatch\":{\"id\":\"%s\"}}}}", user.Id)
		cards, err := cardService.FindCardsByQuery(queryString)
		if err != nil {
			return "", err
		}
		bytes, err := json.Marshal(cards)
		if err != nil {
			return "", err
		}

		return string(bytes), nil
	}

	if user.IsPediatrician() {
		cards, err := GetCardsForPediatrician(user, stub)
		if err != nil {
			return "", err
		}
		bytes, err := json.Marshal(cards)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	if user.IsNeuropathologist() {
		cards, err := GetCardsForNeuropathologist(user, stub)
		if err != nil {
			return "", nil
		}
		bytes, err := json.Marshal(cards)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	return "", nil
}

func GetCardsForPediatrician(user *model.User, stub shim.ChaincodeStubInterface) ([]model.Card, error) {
	agreements, err := service.NewAgreementService(stub).FindAgreementsByDoctor(user.Id)
	cardService := service.NewCardService(stub)
	if  err !=  nil {
		return nil, err
	}
	var cards []model.Card
	for _, agreement := range agreements {
		for _, parent := range agreement.Parents {
			queryString := fmt.Sprintf("{\"selector\":{\"type\":\"card\",\"parents\": {\"$elemMatch\": {\"id\": \"%s\"}}}}", parent)
			parentCards, err := cardService.FindCardsByQuery(queryString)
			if err != nil {
				return nil, err
			}
			cards = append(cards, parentCards...)
			break
		}
	}
	return cards, nil
}


func GetCardsForNeuropathologist(user *model.User, stub shim.ChaincodeStubInterface) ([]map[string]interface{}, error) {
	accesses, err := service.NewAccessService(stub).FindAccessesByDoctor(user.Id)
	if err != nil {
		return nil, err
	}

	var cards []map[string]interface{}
	cardService := service.NewCardService(stub)

	for _, access := range accesses {
		var card model.Card
		err = cardService.FindAndUnmarshal(access.Card, &card)
		if err != nil {
			return nil, err
		}

		responseMap := make(map[string]interface{})
		cardMap := card.ToMap()

		for _, value := range access.Fields {
			cardValue, ok := cardMap[value]
			if ok {
				responseMap[value] = cardValue
			}
		}

		cards = append(cards, responseMap)
	}

	return cards, nil
}
