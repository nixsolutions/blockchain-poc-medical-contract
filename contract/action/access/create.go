package access

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/service"
	"strings"
)

func CreateAccess(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 4 {
		return "", fmt.Errorf("Incorrect arguments. 1 - med card key, 2 - access key,3 - doctor, 4 - fields(comma separated)")
	}
	//TODO: add org check and given by value
	medKey, accessKey, doctor, fieldsString := args[0], args[1], args[2], args[3]
	fields := strings.Split(fieldsString, ",")

	//var card model.Card

	//agreementService := service.NewAgreementService(stub)
	accessService := service.NewAccessService(stub)
	cardService := service.NewCardService(stub)

	//Check med card
	medCardExists, err := cardService.Exists(medKey)
	if !medCardExists {
		return  "", fmt.Errorf("Med card %s does not exist", medKey)
	}
	if err != nil {
		return  "", err
	}

	//TODO: find agreement for doctor and medcard parents and check it

	//TODO: change "doctor1" to cid data(from  cert)
	access := accessService.Create(doctor, "doctor1", medKey, fields)

	jsonBytes, err := json.Marshal(access)
	if err != nil {
		return  "", fmt.Errorf("Failed to marshall access obj", accessKey)
	}

	err = accessService.Put(accessKey, jsonBytes)
	if err != nil {
		return  "", err
	}

	return string(jsonBytes), nil
}