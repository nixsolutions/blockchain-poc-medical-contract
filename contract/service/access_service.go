package service

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"time"
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
	return service.FindAndUnmarshal(service.keyPrefix+key, access)
}

func (service *AccessService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}

func (service *AccessService) Create(doctor string, givenBy string, card string, fields []string) model.Access {
	return model.Access{
		Doctor:    doctor,
		GivenBy:   givenBy,
		Fields:    fields,
		Status:    model.ACCESS_STATUS_VALID,
		Timestamp: time.Now().Unix(),
	}
}

func (service *AccessService) MakeInvalid(access *model.Access) {
	access.Status = model.ACCESS_STATUS_INVALID
}

func (service *AccessService) Put(key string, agreement []byte) error {
	return service.basicRepo.Stub.PutState(service.keyPrefix+key, agreement)
}
