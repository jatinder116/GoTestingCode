package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode/mocks"
	"github.com/stretchr/testify/require"
)

//go:generate counterfeiter -o mocks/transaction.go -fake-name TransactionContext . transactionContext
type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate counterfeiter -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go:generate counterfeiter -o mocks/statequeryiterator.go -fake-name StateQueryIterator . stateQueryIterator
type stateQueryIterator interface {
	shim.StateQueryIteratorInterface
}

func TestInitLedger(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	studentChaincode := SmartContract{}
	err := studentChaincode.InitLedger(transactionContext)
	require.NoError(t, err)

	chaincodeStub.PutStateReturns(fmt.Errorf("failed inserting key"))
	err = studentChaincode.InitLedger(transactionContext)
	require.EqualError(t, err, "failed to put to world state. failed inserting key")
}

func TestCreateStu(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	studentChaincode := SmartContract{}
	err := studentChaincode.CreateStu(transactionContext, "Stu8", "james", "kkk", "jjjj")
	fmt.Println("Endpoint Hit: homePage",err)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns([]byte{}, nil)
	// err = studentChaincode.CreateStu(transactionContext, "Stu9", "james", "kkk", "jjjj")
	// require.EqualError(t, err, "the student stu4 already exists")

	// chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve student"))
	// err = studentChaincode.CreateStu(transactionContext, "Stu10", "","", "")
	// require.EqualError(t, err, "failed to read from world state: unable to retrieve student")
}

func TestReadStu(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedStudent := &Student{ID: "Stu1"}
	bytes, err := json.Marshal(expectedStudent)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	studentChaincode := SmartContract{}
	student, err := studentChaincode.GetStu(transactionContext, "")
	require.NoError(t, err)
	require.Equal(t, expectedStudent, student)

	// chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve student"))
	// _, err = studentChaincode.GetStu(transactionContext, "")
	// require.EqualError(t, err, "failed to read from world state: unable to retrieve student")

	// chaincodeStub.GetStateReturns(nil, nil)
	// student, err = studentChaincode.GetStu(transactionContext, "Stu1")
	// require.EqualError(t, err, "the Student stu1 does not exist")
	// require.Nil(t, student)
}

func TestUpdateStu(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedStudent := &Student{ID: "Stu4"}
	bytes, err := json.Marshal(expectedStudent)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	studentChaincode := SmartContract{}
	err = studentChaincode.UpdateStu(transactionContext, "Stu4", "Happy", "Male", "France")
	require.NoError(t, err)

	// chaincodeStub.GetStateReturns(nil, nil)
	// err = studentChaincode.UpdateStu(transactionContext, "Stu4", "", "", "")
	// require.EqualError(t, err, "the Student Stu4 does not exist")

	// chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve Student"))
	// err = studentChaincode.UpdateStu(transactionContext, "Stu4", "", "", "")
	// require.EqualError(t, err, "failed to read from world state: unable to retrieve Student")
}

func TestDeleteStu(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	student := &Student{ID: "Stu3"}
	bytes, err := json.Marshal(student)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	chaincodeStub.DelStateReturns(nil)
	studentChaincode := SmartContract{}
	err = studentChaincode.DeleteStu(transactionContext, "Stu3")
	require.NoError(t, err)

	// chaincodeStub.GetStateReturns(nil, nil)
	// err = studentChaincode.DeleteStu(transactionContext, "Stu3")
	// require.EqualError(t, err, "the Student Stu3 does not exist")

	// chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve Student"))
	// err = studentChaincode.DeleteStu(transactionContext, "")
	// require.EqualError(t, err, "failed to read from world state: unable to retrieve Student")
}


func TestGetAllStu(t *testing.T) {
	student := &Student{ID: "Stu1"}
	bytes, err := json.Marshal(student)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)

	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	chaincodeStub.GetStateByRangeReturns(iterator, nil)
	studentChaincode := SmartContract{}
	students, err := studentChaincode.GetAllStu(transactionContext)
	require.NoError(t, err)
	require.Equal(t, []*Student{student}, students)

	iterator.HasNextReturns(true)
	iterator.NextReturns(nil, fmt.Errorf("failed retrieving next item"))
	students, err = studentChaincode.GetAllStu(transactionContext)
	require.EqualError(t, err, "failed retrieving next item")
	require.Nil(t, students)

	chaincodeStub.GetStateByRangeReturns(nil, fmt.Errorf("failed retrieving all Students"))
	students, err = studentChaincode.GetAllStu(transactionContext)
	require.EqualError(t, err, "failed retrieving all Students")
	require.Nil(t, students)
}
