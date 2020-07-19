package main
import (
	"fmt"
	"student"
)
type Studentjson = student.Student

func main() {
	var jsonText = []byte(`[
        {"stuid":1,"stuname":"Alicia","stucity":"USA"},
        {"stuid":2,"stuname":"Mike","stucity":"Canada"},
        {"stuid":3,"stuname":"Milina","stucity":"Russia"}
	]`)
	//============ Add student ===========================
	addStu := student.AddStudent(jsonText,Studentjson{Id: 4,Name:"Diana",City:"Spain",});
	fmt.Println("New Student Data: ",addStu);

		//============ Find Student =============================
		getStu, status := student.FindStudent([]byte(addStu),7);
		if status==0 {
			fmt.Println("Value not found in slice")
		}else{
		fmt.Println("Find Student with id: ",getStu);
		}

	//============ Update Student =============================
	updateStu := student.UpdateStudent([]byte(addStu),3,"Emilie");
	fmt.Println("Update student: ",updateStu);

	//============ Delete Student =============================
	delStu := student.DeleteStudent([]byte(addStu),2);
	fmt.Println("Delete Student: ",delStu);
	
}