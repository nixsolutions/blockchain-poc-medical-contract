package access

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
)

func GetAccess(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	var access model.Access
	accessService := service.NewAccessService(stub)
	err := accessService.FindAndUnmarshal(args[0], &access)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(access)
	if err != nil {
		return  "", fmt.Errorf("Failed to marshall access obj", args[0])
	}

	return string(jsonBytes), nil
}
