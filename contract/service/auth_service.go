package service

import (
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
func (service *AuthService) GetUser() *model.User {
	return model.NewUser()
}
