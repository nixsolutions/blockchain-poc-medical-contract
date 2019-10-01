package service

import (
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
	return service.FindAndUnmarshal(service.keyPrefix+key, card)
}

func (service *CardService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}
