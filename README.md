# MedicalContract chaincode(smart contract)

## Actors

There are 2 actors:
 - Hospital organization
 - Parent organization

## Purpose

The main purpose of MedicalContract is to store medical records of children in the blockchain system. There are some rules related to access and update actions for cards:
 - only parent can create a card
 - only hospital worker can update card data
 - parent can read cards of their children
 - hospital can read cards of examined children (hospital and parent must have an agreement)
 - only parent can create agreement
 - only hospital can sign an agreement

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
- createAgreement - Parent can create an agreement with the hospital, where their child is examined
- getAgreement - Parent and a hospital worker can see agreement from the world state database
- signAgreement - hospital can sign an agreement with a parent

And there are actions for cards:
- getCard - Parent and the hospital can get card of the child
- getCards - Parent can get cards of their children and hospital worker can get cards of examined children
- createCard - Parent can create a card
- updateCard - Hospital worker can update card


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

## License
The project is developed by [NIX Solutions](http://nixsolutions.com) Go team and distributed under [MIT LICENSE](https://github.com/nixsolutions/blockchain-poc-medical-contract/master/LICENSE)
