package assignment2pkg

import (
	"gotestcode/assignment2pkg/structpkg"
)
type Student = structpkg.Student

//=========== Add Student ===========================
func AddStudent(Data []Student, Name Student) (a[]Student ){
	a = append(Data,Name);
	return a;
}
