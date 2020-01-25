package contract

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"poc/contract/action/agreement"
	"poc/contract/action/card"
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
	case "createCard":
		result, err = card.Create(stub, args)
	case "updateCard":
		result, err = card.Update(stub, args)
	case "getCard":
		result, err = card.Get(stub, args)
	case "getCards":
		result, err = card.GetAll(stub, args)

	case "getAgreement":
		result, err = agreement.Get(stub, args)
	case "createAgreement":
		result, err = agreement.Create(stub, args)
	case "signAgreement":
		result, err = agreement.Sign(stub, args)

	default:
		result, err = card.Get(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}
