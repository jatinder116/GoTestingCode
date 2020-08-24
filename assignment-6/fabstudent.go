package main

import (
	"log"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Student struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	City     string `json:"city"`
}

// InitLedger adds a base set of students to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	students := []Student{
		{ID: "Stu1", Name: "Jatinder", Gender: "Male", City: "India"},
		{ID: "Stu2", Name: "James", Gender: "Male", City: "Canada"},
		{ID: "Stu3", Name: "Chang", Gender: "Male", City: "Japan"},
		{ID: "Stu4", Name: "Alicia", Gender: "Female", City: "Russia"},
	}

	for _, student := range students {
		studentJSON, err := json.Marshal(student)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(student.ID, studentJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

//  create a new student with given details
func (s *SmartContract) CreateStu(ctx contractapi.TransactionContextInterface, id string, name string, gender string, city string) error {

	exists, err := s.CheckStu(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the student %s already exists", id)
	}

	student := Student{
		ID:             id,
		Name:          name,
		Gender:         gender,
		City:          city,
	}
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}
	
	fmt.Println("Endpoint Hit: homePage",id)
	fmt.Println("Endpoint Hit: homePage=============",string(studentJSON))
	return ctx.GetStub().PutState(id, studentJSON)
}

// returns the student with given id
func (s *SmartContract) GetStu(ctx contractapi.TransactionContextInterface, id string) (*Student, error) {
	studentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if studentJSON == nil {
		return nil, fmt.Errorf("the student %s does not exist", id)
	}
	fmt.Println("Endpoint Hit: homePagesdvsdf=============",string(studentJSON))
	var student Student
	err = json.Unmarshal(studentJSON, &student)
	if err != nil {
		return nil, err
	}
	fmt.Println("Endpoint Hit: homePage",id)
	// fmt.Println("Endpoint Hit: homePage=============",student)
	fmt.Printf("%+v", student)
	return &student, nil
}

//updates an existing student
func (s *SmartContract) UpdateStu(ctx contractapi.TransactionContextInterface, id string, name string, gender string, city string) error {
	exists, err := s.CheckStu(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the student %s does not exist", id)
	}

	student := Student{
		ID:             id,
		Name:          name,
		Gender:         gender,
		City:          city,
	}
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, studentJSON)
}

//deletes a given student
func (s *SmartContract) DeleteStu(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.CheckStu(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the student %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// returns true when student with given ID exists
func (s *SmartContract) CheckStu(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	studentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return studentJSON != nil, nil
}

// returns all students
func (s *SmartContract) GetAllStu(ctx contractapi.TransactionContextInterface) ([]*Student, error) {
	// range query with empty string for startKey and endKey does an
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var students []*Student
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var student *Student
		err = json.Unmarshal(queryResponse.Value, &student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}



func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
