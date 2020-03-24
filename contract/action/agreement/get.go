package agreement

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract/service"
)

func Get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}

	agreementService := service.NewAgreementService(stub)
	var response interface{}

	if user.IsParent() {
		response, err = agreementService.FindAgreementsByParent(user.Id)
	}

	if user.IsHospitalWorker() {
		response, err = agreementService.FindAgreementsByDoctor(user.Id)
	}

	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", fmt.Errorf("Failed to marshall agreement obj", args[0])
	}

	return string(jsonBytes), nil
}
