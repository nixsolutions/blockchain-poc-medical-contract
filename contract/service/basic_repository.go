package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type BasicRepository struct {
	Stub shim.ChaincodeStubInterface
}

func (repo *BasicRepository) Find(key string) []byte {
	bytes, _ := repo.Stub.GetState(key)

	return bytes
}

func (repo *BasicRepository) FindAndUnmarshal(key string, dest interface{}) error {
	bytes := repo.Find(key)
	if bytes == nil {
		return nil
	}
	err := json.Unmarshal([]byte(bytes), dest)
	if err != nil {
		return fmt.Errorf("Failed to unmarshall obj", key)
	}

	return nil
}

func (repo *BasicRepository) Exists(key string) (bool, error) {
	bytes := repo.Find(key)
	return bytes != nil, nil
}
