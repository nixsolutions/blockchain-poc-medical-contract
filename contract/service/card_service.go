package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
)

type CardService struct {
	basicRepo *BasicRepository
	keyPrefix string
}

func NewCardService(stub shim.ChaincodeStubInterface) *CardService {
	return &CardService{basicRepo: &BasicRepository{Stub: stub}, keyPrefix: "CARD"}
}

func (service *CardService) Find(key string) ([]byte, error) {
	return service.basicRepo.Find(service.keyPrefix + key)
}

func (service *CardService) FindAndUnmarshal(key string, card *model.Card) error {
	return service.basicRepo.FindAndUnmarshal(service.keyPrefix+key, card)
}

func (service *CardService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}

func (service *CardService) FindCardsByParent(parentId string) ([]model.Card, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"card\",\"parent\":\"%s\"}}", parentId)
	return service.FindCardsByQuery(queryString)
}

func (service *CardService) FindCardsByQuery(queryString string) ([]model.Card, error) {
	resultsIterator, err := service.basicRepo.Stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var medCards []model.Card

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var card model.Card
		err = json.Unmarshal(queryResponse.Value, &card)
		if err != nil {
			return nil, err
		}
		medCards = append(medCards, card)
	}
	return medCards, nil
}

func (service *CardService) HasAccessToCard(cardKey string) (bool, error) {
	var card model.Card
	err := service.FindAndUnmarshal(cardKey, &card)
	if err != nil {
		return false, err
	}

	user, err := NewAuthService(service.basicRepo.Stub).GetUser()
	if err != nil {
		return false, err
	}

	if user.IsParent() {
		return card.Parent == user.Id, nil
	}

	if user.IsHospitalWorker() {
		agreement, _ := NewAgreementService(service.basicRepo.Stub).FindAgreementByDoctorAndParent(user.Id, card.Parent)
		if agreement != nil {
			return agreement.Status == model.SIGNED_STATUS, nil
		}

		return false, errors.New("agreement was not found for card owner")
	}

	return false, nil
}
