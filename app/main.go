package main

import "fmt"

type Student struct {
    Id    int
    Name  string
    City string
}

// =============== Add Student ==========================================
func AddStudent(Dummy []Student, Name Student) (a[]Student ){
	a = append(Dummy,Name);
	return a;
}

// ==================== Find Student ===================================
func FindStudent(Dummy []Student, Id int) (a Student ){
	for i := 0; i < len(Dummy); i++ {
		if Dummy[i].Id == Id {
			a = Dummy[i];
		}
	}
	return a;
}

// ============== Update Student ==========================================
func UpdateStudent(Dummy []Student, Id int,ChangedName string) (a[]Student ){
	for i := 0; i < len(Dummy); i++ {
		if Dummy[i].Id == Id {
			Dummy[i].Name = ChangedName;
		}
	}
	return Dummy;
}

// ==================== Delete Student =======================================
func DeleteStudent(Dummy []Student, Id int) (a[]Student ){
	for i := 0; i < len(Dummy); i++ {
		if Dummy[i].Id == Id {
			a = append(Dummy[:i], Dummy[i+1:]...)
			break
		}
	}
	return a;
}


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
	g := AddStudent(students,Student{Id: 4,Name:"Diana",City:"Spain",});
	fmt.Println("New Student Data: ",g);
	l := FindStudent(g,4);
	fmt.Println("Find Student with id: ",l);
	k := UpdateStudent(g,3,"Kewin");
	fmt.Println("Update student: ",k);
	h := DeleteStudent(k,4);
	fmt.Println("Delete Student: ",h);
}