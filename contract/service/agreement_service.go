package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"time"
)

type AgreementService struct {
	basicRepo *BasicRepository
	keyPrefix string
}

func NewAgreementService(stub shim.ChaincodeStubInterface) *AgreementService {
	return &AgreementService{basicRepo: &BasicRepository{Stub: stub}, keyPrefix: "AGR"}
}

func (service *AgreementService) Find(key string) ([]byte, error) {
	return service.basicRepo.Find(service.keyPrefix + key)
}

func (service *AgreementService) FindAndUnmarshal(key string, agreement *model.Agreement) error {
	return service.basicRepo.FindAndUnmarshal(service.keyPrefix+key, agreement)
}

func (service *AgreementService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}

func (service *AgreementService) Create(doctor string, parents []string) model.Agreement {
	return model.Agreement{
		Status:    model.TO_SIGN_STATUS,
		Doctor:    doctor,
		Parents:   parents,
		Timestamp: time.Now().Unix(),
		Type: "agreement",
	}
}

func (service *AgreementService) Put(key string, agreement []byte) error {
	return service.basicRepo.Stub.PutState(service.keyPrefix+key, agreement)
}

func (service *AgreementService) Sign(agreement *model.Agreement) {
	agreement.Status = model.SIGNED_STATUS
}

func (service *AgreementService) FindAgreementByDoctorAndParent(doctorId string, parentId string) (*model.Agreement, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"agreement\",\"doctor\":\"%s\",\"parents\": {\"$elemMatch\": {\"id\": \"%s\"}}}}", doctorId, parentId)

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
		var agreement model.Agreement
		err = json.Unmarshal(queryResponse.Value, &agreement)
		if err != nil {
			return nil, err
		}
		return &agreement, nil
	}
	return nil, errors.New("agreement was not found")
}

func (service *AgreementService) FindAgreementsByDoctor(doctorId string) ([]model.Agreement, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"agreement\",\"doctor\":\"%s\"}}", doctorId)

	resultsIterator, err := service.basicRepo.Stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var agreements []model.Agreement
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var agreement model.Agreement
		err = json.Unmarshal(queryResponse.Value, &agreement)
		if err != nil {
			return nil, err
		}
		agreements = append(agreements, agreement)
	}

	return agreements, nil
}
