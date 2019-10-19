package access

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/service"
	"strconv"
	"strings"
)

func CreateAccess(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 4 {
		return "", fmt.Errorf("Incorrect arguments. 1 - med card key, 2 - access key,3 - doctor, 4 - fields(comma separated), 5 - timestamp unix")
	}

	medKey, accessKey, doctor, fieldsString, timestampString := args[0], args[1], args[2], args[3], args[4]
	fields := strings.Split(fieldsString, ",")

	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		return "", err
	}
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
	//TODO: add agreement check

	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}
	if !user.IsPediatrician() {
		return "", errors.New("user is not a Pediatrician")
	}

	access := accessService.Create(doctor, user.Id, medKey, fields, timestamp)
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