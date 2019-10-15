package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type BasicRepository struct {
	Stub shim.ChaincodeStubInterface
}

func (repo  *BasicRepository) Find(key string) ([]byte, error) {
	bytes, err := repo.Stub.GetState(key)
	if err != nil {
		return  nil, fmt.Errorf("Failed to get obj: %s with error: %s", key, err)
	}

	if bytes == nil {
		return  nil, fmt.Errorf("Failed to get obj: %s with error: %s", key, err)
	}

	return bytes, nil
}

func (repo *BasicRepository) FindAndUnmarshal(key string, dest interface{}) error {
	bytes, err := repo.Find(key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(bytes), dest)
	if err != nil {
		return  fmt.Errorf("Failed to unmarshall obj", key)
	}

	return nil
}

func (repo *BasicRepository) Exists(key string) (bool, error) {
	bytes, err := repo.Find(key)
	if err != nil {
		return false, err
	}
	return bytes != nil, nil
}
