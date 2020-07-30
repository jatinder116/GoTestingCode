
/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 2 utility libraries for formatting and reading and writing JSON
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the Owner structure, with 4 properties.  Structure tags are used by encoding/json library
type Owner struct {
	Id   string `json:"id"`
	Name  string `json:"name"`
	Compnay string `json:"company"`
	Gender  string `json:"gender"`
}

/*
 * Init is called during chaincode instantiation to initialize any data
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	switch function {
	case "createOwner":
		fmt.Println("add owner")
		return s.createOwner(APIstub, args)
	case "getOwner":
		fmt.Println("Get owner")
		return s.getOwner(APIstub, args)
	default:
		fmt.Println("Invalid Smart Contract function name.")
		return shim.Error("Invalid Smart Contract function name.")
	}
}

// Get owner by id
func (s *SmartContract) getOwner(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ownerAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Owner does not exist with Id '" + args[0] + "'")
		}
	return shim.Success(ownerAsBytes)
}

// Create an owner
func (s *SmartContract) createOwner(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting at least 3")
	}
	var owner = Owner{Id: args[0], Name: args[1], Compnay: args[2], Gender: args[3]}
	// save the Owner
	ownerAsBytes, _ := json.Marshal(owner)
	err = APIstub.PutState(args[0], ownerAsBytes)
	if err != nil {
		fmt.Println("Could not save an Owner")
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
