package repository

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
)

type AgreementRepository struct {
	Stub shim.ChaincodeStubInterface
}

func (ar  *AgreementRepository) Find(key string, agreement *model.Agreement) error {
	agreementBytes, err := ar.Stub.GetState(key)
	if err != nil {
		return  fmt.Errorf("Failed to get agreement: %s with error: %s", key, err)
	}
	if agreementBytes == nil {
		return  fmt.Errorf("Agreement was not found: %s", key)
	}

	err = json.Unmarshal([]byte(agreementBytes), agreement)
	if err != nil {
		return  fmt.Errorf("Failed to unmarshall agreement obj", key)
	}

	return nil
}