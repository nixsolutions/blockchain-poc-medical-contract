package service

import (
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
	return service.FindAndUnmarshal(service.keyPrefix+key, agreement)
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
	}
}

func (service *AgreementService) Put(key string, agreement []byte) error {
	return service.basicRepo.Stub.PutState(service.keyPrefix+key, agreement)
}

func (service *AgreementService) Sign(agreement *model.Agreement) {
	agreement.Status = model.SIGNED_STATUS
}
