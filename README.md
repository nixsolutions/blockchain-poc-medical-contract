# MedicalContract chaincode(smart contract)

## Where to start

Every program in Golang starts with a function main(). You can find main() function in  main.go.
In main() function we start the chaincode. We pass a reference of MedicalContract object to shim.Start() method to start the chaincode.

## Contract methods

You can find MedicalContract struct and implementation of 2 required methods `Invoke` and `Init` in `contract/MedicalContract.go`.
`Init` is called during chaincode instantiation to initialize any data. Note that chaincode upgrade also calls this function to reset or to migrate data.
`Invoke` is called per transaction on the chaincode. Each transaction is either a 'get' or a 'set' on the asset created by `Init` function. The Set method may create a new asset by specifying a new key-value pair.


## Action handlers
You can find actions in `contract/action` folder. Just like controllers in web apps with MVC pattern, actions are connectors between business logic layer of smart contract and transport layer.
There are actions for agreement:
- createAgreement
- getAgreement
- signAgreement

And there are actions for cards:
- getCard
- getCards
- createCard
- updateCard


## Models
You can find models in `contract/model` folder. This folder contains entities of smart contract.
There are 3 models:
- Agreement
- Card
- User

## Services
You can find services in `contract/service` folder. This  folder contains business logic of smart contract
There are `AgreementService` to work with agreements and `CardService`  to work with cards. 
AuthService is a wrapper to define who invokes the chaincode.
 