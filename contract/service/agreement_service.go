package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
)

type AgreementService struct {
	basicRepo *BasicRepository
	keyPrefix string
}

func NewAgreementService(stub shim.ChaincodeStubInterface) *AgreementService {
	return &AgreementService{basicRepo: &BasicRepository{Stub: stub}, keyPrefix: "AGR"}
}

func (service *AgreementService) Find(key string) []byte {
	return service.basicRepo.Find(service.keyPrefix + key)
}

func (service *AgreementService) FindAndUnmarshal(key string, agreement *model.Agreement) error {
	return service.basicRepo.FindAndUnmarshal(service.keyPrefix+key, agreement)
}

func (service *AgreementService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}

func (service *AgreementService) Create(id string, doctor string, parent string, timestamp int64) model.Agreement {
	return model.Agreement{
		Id:        id,
		Status:    model.TO_SIGN_STATUS,
		Doctor:    doctor,
		Parent:    parent,
		Timestamp: timestamp,
		Type:      "agreement",
	}
}

func (service *AgreementService) Put(key string, agreement []byte) error {
	return service.basicRepo.Stub.PutState(service.keyPrefix+key, agreement)
}

func (service *AgreementService) Sign(agreement *model.Agreement) {
	agreement.Status = model.SIGNED_STATUS
}

func (service *AgreementService) FindAgreementByDoctorAndParent(doctorId string, parentId string) (*model.Agreement, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"agreement\",\"doctor\":\"%s\",\"parent\":\"%s\"}}", doctorId, parentId)
	agreements, err := service.FindAgreementsByQuery(queryString)
	if err != nil {
		return nil, err
	}
	if len(agreements) == 0 {
		return nil, errors.New("agreements was not found")
	}

	return &agreements[0], nil
}

func (service *AgreementService) FindAgreementsByDoctor(doctorId string) ([]model.Agreement, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"agreement\",\"doctor\":\"%s\"}}", doctorId)
	return service.FindAgreementsByQuery(queryString)
}

func (service *AgreementService) FindAgreementsByParent(parentId string) ([]model.Agreement, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"type\":\"agreement\",\"parent\":\"%s\"}}", parentId)
	return service.FindAgreementsByQuery(queryString)
}

func (service *AgreementService) FindAgreementsByQuery(query string) ([]model.Agreement, error) {
	resultsIterator, err := service.basicRepo.Stub.GetQueryResult(query)
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
