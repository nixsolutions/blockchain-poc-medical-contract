package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
)

type AccessService struct {
	basicRepo *BasicRepository
	keyPrefix string
}

func NewAccessService(stub shim.ChaincodeStubInterface) *AccessService {
	return &AccessService{basicRepo: &BasicRepository{Stub: stub}, keyPrefix: "ACS"}
}

func (service *AccessService) Find(key string) ([]byte, error) {
	return service.basicRepo.Find(service.keyPrefix + key)
}

func (service *AccessService) FindAndUnmarshal(key string, access *model.Access) error {
	return service.basicRepo.FindAndUnmarshal(service.keyPrefix+key, access)
}

func (service *AccessService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}

func (service *AccessService) Create(doctor string, givenBy string, card string, fields []string, timestamp int64) model.Access {
	return model.Access{
		Doctor:    doctor,
		GivenBy:   givenBy,
		Fields:    fields,
		Status:    model.ACCESS_STATUS_VALID,
		Timestamp: timestamp,
		Card:      card,
		Type: "access",
	}
}

func (service *AccessService) MakeInvalid(access *model.Access) {
	access.Status = model.ACCESS_STATUS_INVALID
}

func (service *AccessService) Put(key string, agreement []byte) error {
	return service.basicRepo.Stub.PutState(service.keyPrefix+key, agreement)
}

func (service *AccessService) FindAccessByDoctor(doctorId string, cardId string) (*model.Access, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"access\",\"status\":\"valid\",\"doctor\":\"%s\",\"card\":\"%s\"}}", doctorId, cardId)

	resultsIterator, err := service.basicRepo.Stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var access model.Access
		err = json.Unmarshal(queryResponse.Value, &access)
		if err != nil {
			return nil, err
		}
		return &access, nil
	}
	return nil, errors.New("access was not found")
}

func (service *AccessService) FindAccessesByDoctor(doctorId string) ([]model.Access, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"access\",\"doctor\":\"%s\"}}", doctorId)

	resultsIterator, err := service.basicRepo.Stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var accesses []model.Access
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var access model.Access
		err = json.Unmarshal(queryResponse.Value, &access)
		if err != nil {
			return nil, err
		}
		accesses = append(accesses, access)
	}
	return accesses, errors.New("access was not found")
}
