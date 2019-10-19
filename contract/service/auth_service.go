package service

import (
	"errors"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
)

type AuthService struct {
	basicRepo *BasicRepository
}

func NewAuthService(stub shim.ChaincodeStubInterface) *AuthService {
	return &AuthService{basicRepo: &BasicRepository{Stub: stub}}
}

//TODO: change to cid attr data
func (service *AuthService) GetUser() (*model.User, error) {
	// Get the client ID object
	id, err := cid.New(service.basicRepo.Stub)
	if err != nil {
		return nil, err
	}
	mspid, err := id.GetMSPID()
	if err != nil {
		return nil, err
	}
	role, found, err := id.GetAttributeValue("role")
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("role attr was not found")
	}

	personId, found, err := id.GetAttributeValue("id")
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("role attr was not found")
	}

	return model.NewUser(personId, role, mspid), nil
}
