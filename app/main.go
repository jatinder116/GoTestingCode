package main

import (
	"fmt"
	"gotestcode/assignment3pkg"
	"gotestcode/assignment3pkg/structpkg"
)

type Studentjson = structpkg.Student


func main() {
	//============= assignment 2 file data =======================================
	student();
	fmt.Println("============== *************** ======================")
	fmt.Println("============== *************** ======================")

	var jsonText = []byte(`[
        {"stuid":1,"stuname":"Alicia","stucity":"USA"},
        {"stuid":2,"stuname":"Mike","stucity":"Canada"},
        {"stuid":3,"stuname":"Milina","stucity":"Russia"}
	]`)
	//============ Add student ===========================
	g := assignment3pkg.AddStudent(jsonText,Studentjson{Id: 4,Name:"Diana",City:"Spain",});
	fmt.Println("New Student Data: ",g);

		//============ Find Student =============================
		l, m := assignment3pkg.FindStudent([]byte(g),7);
		if m==0 {
			fmt.Println("Value not found in slice")
		}else{
		fmt.Println("Find Student with id: ",l);
		}

	//============ Update Student =============================
	k := assignment3pkg.UpdateStudent([]byte(g),3,"Emilie");
	fmt.Println("Update student: ",k);


	//============ Delete Student =============================
	h := assignment3pkg.DeleteStudent([]byte(g),2);
	fmt.Println("Delete Student: ",h);
	


	//======================== testing data==============================
		//============ Convert json data to struct data ====================
	// 	studentJsonData :=`{"student":[
	// 		{"stuid":1,"stuname":"Alicia","stucity":"USA"},
	// 		{"stuid":2,"stuname":"Mike","stucity":"Canada"},
	// 		{"stuid":3,"stuname":"Milina","stucity":"Russia"}
	// 		]}`
	
	// 	 structureData:= make(map[string][]Studentjson)
	// 	err := json.Unmarshal([]byte(studentJsonData), &structureData)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("convert json to struct data", structureData)
	
	// 	//==================================================================
	// 	students:= []Studentjson{
	// 		Studentjson{
	// 			Id: 1,
	// 			Name:"Alicia",
	// 			City:"USA",
	// 		},
	// 		Studentjson{
	// 			Id: 2,
	// 			Name:"Mike",
	// 			City:"Canada",
	// 		},
	// 		Studentjson{
	// 			Id: 3,
	// 			Name:"Milina",
	// 			City:"Russia",
	// 		},
	// 	}
	
	// 	student := &students
	// 	e, err := json.Marshal(student)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	
	// fmt.Println(string(e))
}