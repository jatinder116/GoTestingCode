package main

import (
	"fmt"
	"gotestcode/assignment2pkg"
	"gotestcode/assignment2pkg/structpkg"
	
)
type Student = structpkg.Student

func student() {
	//======= assignment1 data ======================
	sqrprime()

	students:= []Student{
		Student{
			Id: 1,
			Name:"Alicia",
			City:"USA",
		},
		Student{
			Id: 2,
			Name:"Mike",
			City:"Canada",
		},
		Student{
			Id: 3,
			Name:"Milina",
			City:"Russia",
		},
	}

	//============ Add New Student Value=============================
	g := assignment2pkg.AddStudent(students,Student{Id: 4,Name:"Diana",City:"Spain",});
	fmt.Println("New Student Data: ",g);

	//============ Add New Student Value with struct literals=============================
	u := assignment2pkg.AddStudent(g,Student{Id: 5,City:"Spain",});
	fmt.Println("New Student Data with struct literals: ",u);

	//============ Add New Student Value with struct literals=============================
	t := assignment2pkg.AddStudent(u,Student{Id:6});
	fmt.Println("New Student Data with struct literals: ",t);

	//============ Find Student =============================
	l, m := assignment2pkg.FindStudent(t,7);
	if !m {
        fmt.Println("Value not found in slice")
    }else{
	fmt.Println("Find Student with id: ",l);
	}

	//============ Update Student =============================
	k := assignment2pkg.UpdateStudent(t,3,"Kewin");
	fmt.Println("Update student: ",k);

	//============ Delete Student =============================
	h := assignment2pkg.DeleteStudent(k,4);
	fmt.Println("Delete Student: ",h);
}