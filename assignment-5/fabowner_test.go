package main

import (
	"fmt"
	"strings"
	"testing"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func TestInstancesCreation(test *testing.T) {
	stub := InitChaincode(test)
	Invoke(test, stub, "createOwner", "o1", "George", "BMW", "Male")
	Invoke(test, stub, "createOwner", "o2", "Mike", "Apple", "Male")
	Invoke(test, stub, "createOwner", "o3", "James", "Microsoft", "Male")
	Invoke(test, stub, "getOwner", "o2")
}

func InitChaincode(test *testing.T) *shimtest.MockStub {
	stub := shimtest.NewMockStub("testingStub", new(SmartContract))
	result := stub.MockInit("000", nil)

	if result.Status != shim.OK {
		test.FailNow()
	}
	return stub
}


func Invoke(test *testing.T, stub *shimtest.MockStub, function string, args ...string) {
	cc_args := make([][]byte, 1+len(args))
	cc_args[0] = []byte(function)
	for i, arg := range args {
		cc_args[i+1] = []byte(arg)
	}
	result := stub.MockInvoke("000", cc_args)
	fmt.Println("Call:    ", function, "(", strings.Join(args, ","), ")")
	fmt.Println("Code: ", result.Status)
	fmt.Println("Payload: ", string(result.Payload))

	if result.Status != shim.OK {
		test.FailNow()
	}
}

