package main

import (
	"fmt"
	"student"
)
type Student = student.Student

func main() {

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
	addStu := student.AddStudent(students,Student{Id: 4,Name:"Diana",City:"Spain",});
	fmt.Println("New Student Data: ",addStu);

	//============ Add New Student Value with struct literals=============================
	addStuStru:= student.AddStudent(addStu,Student{Id: 5,City:"Spain",});
	fmt.Println("New Student Data with struct literals: ",addStuStru);

	//============ Add New Student Value with struct literals=============================
	addStuStruLit := student.AddStudent(addStuStru,Student{Id:6});
	fmt.Println("New Student Data with struct literals: ",addStuStruLit);

	//============ Find Student =============================
	getStu, err := student.FindStudent(addStuStruLit,7);
	if !err {
        fmt.Println("Value not found in slice")
	}else{
	fmt.Println("Find Student with id: ",getStu);
	}

	//============ Update Student =============================
	updateStu := student.UpdateStudent(addStuStruLit,3,"Kewin");
	fmt.Println("Update student: ",updateStu);

	//============ Delete Student =============================
	delStu := student.DeleteStudent(updateStu,4);
	fmt.Println("Delete Student: ",delStu);
}