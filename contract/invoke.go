package contract

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"poc/contract/action/access"
	"poc/contract/action/agreement"
	"poc/contract/action/card"
	"poc/contract/action/vaccination"
)

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *MedicalContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	switch fn {
	case "setCard":
		result, err = card.SetCard(stub, args)
	case "getCard":
		result, err = card.GetCard(stub, args)

	case "createAccess":
		result, err = access.CreateAccess(stub, args)
	case "getAccess":
		result, err = access.GetAccess(stub, args)

	case "getAgreement":
		result, err = agreement.GetAgreement(stub, args)
	case "createAgreement":
		result, err = agreement.CreateAgreement(stub, args)
	case "signAgreement":
		result, err = agreement.SignAgreement(stub, args)

	case "addVaccination":
		result, err = vaccination.AddVaccination(stub, args)
	case "updateVaccinationTimestamp":
		result, err = vaccination.UpdateVaccinationTimestamp(stub, args)
	case "deleteVaccination":
		result, err = vaccination.DeleteVaccination(stub, args)

	default:
		result, err = card.GetCard(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}
