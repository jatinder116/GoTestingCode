package main

import (
	"fmt"
	"gotestcode/assignment3pkg"
	"gotestcode/assignment3pkg/structpkg"
	"encoding/json"
)

type Studentjson = structpkg.Student


func main() {
	//============= assignment 2 file data =======================================
	student();
	fmt.Println("=======================")

	//============ Convert json data to struct data ====================
	studentJsonData :=`{"student":[
		{"stuid":1,"stuname":"Alicia","stucity":"USA"},
		{"stuid":2,"stuname":"Mike","stucity":"Canada"},
		{"stuid":3,"stuname":"Milina","stucity":"Russia"}
		]}`

	 structureData:= make(map[string][]Studentjson)
	err := json.Unmarshal([]byte(studentJsonData), &structureData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("convert json to struct data", structureData)

	//==================================================================
	students:= []Studentjson{
		Studentjson{
			Id: 1,
			Name:"Alicia",
			City:"USA",
		},
		Studentjson{
			Id: 2,
			Name:"Mike",
			City:"Canada",
		},
		Studentjson{
			Id: 3,
			Name:"Milina",
			City:"Russia",
		},
	}

	emp := &students
    e, err := json.Marshal(emp)
    if err != nil {
        fmt.Println(err)
        return
	}
	
	g := assignment3pkg.AddStudent(students,Studentjson{Id: 4,Name:"Diana",City:"Spain",});
	fmt.Println("New Student Data: ",g);
    fmt.Println(string(e))
}