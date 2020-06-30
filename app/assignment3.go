package main

import (
	"fmt"
	"gotestcode/assignment3pkg"
	"gotestcode/assignment3pkg/structpkg"
)

type Studentjson = structpkg.Student


func main() {
	//============= assignment 2 file data =======================================
	// student();
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
	
}