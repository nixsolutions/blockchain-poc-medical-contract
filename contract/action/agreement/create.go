package agreement

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
	"strings"
)

// Get returns the value of the specified asset key
func CreateAgreement(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	//TODO: add org check
	key, doctor, parentsString := args[0], args[1], args[2]
	parents := strings.Split(parentsString, ",")

	var agreement model.Agreement
	agreementService := service.NewAgreementService(stub)
	bytes, err := agreementService.Find(key)

	if bytes != nil {
		return "", fmt.Errorf("Agreement with the same key is already created", key)
	}
	if err != nil {
		return  "", err
	}

	agreement = agreementService.Create(doctor, parents)

	jsonBytes, err := json.Marshal(agreement)
	if err != nil {
		return  "", fmt.Errorf("Failed to marshall agreement obj", key)
	}

	err = agreementService.Put(key, jsonBytes)
	if err != nil {
		return  "", err
	}

	return string(jsonBytes), nil
}