package agreement

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/model"
	"poc/contract/service"
	"strconv"
)

func Create(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key, doctor, timestamp")
	}

	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}
	if !user.IsParent() {
		return "", errors.New("only parents can create agreements")
	}
	key, doctor, timestampString := args[0], args[1], args[2]
	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		return "", err
	}

	var agreement model.Agreement
	agreementService := service.NewAgreementService(stub)
	bytes := agreementService.Find(key)

	if bytes != nil {
		return "", fmt.Errorf("Agreement with the same key is already created", key)
	}

	agreement = agreementService.Create(key, doctor, user.Id, timestamp)
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