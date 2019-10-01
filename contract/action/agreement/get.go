package agreement

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

// Get returns the value of the specified asset key
func GetAgreement(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	var agreement model.Agreement
	agreementService := service.NewAgreementService(stub)
	err := agreementService.FindAndUnmarshal(args[0], &agreement)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(agreement)
	if err != nil {
		return  "", fmt.Errorf("Failed to marshall agreement obj", args[0])
	}

	return string(jsonBytes), nil
}
