package student

import (
	"fmt"
	"encoding/json"
)

//=========== Add Student ===========================
func AddStudent(Data []byte, Newdata Student) (a string) {
	var studentData []Student
	//======== unmarshel ==============================
	if err := json.Unmarshal([]byte(Data), &studentData); err != nil {
		 fmt.Println(err)
	}
	studentData = append(studentData, Newdata)
	// ====================== Marshel============
	result, err := json.Marshal(studentData)
	if err != nil {
		fmt.Println(err)
	}
	return string(result);
}

// ==================== Find Student ===================================
func FindStudent(Data []byte, Id int) (a string, b int ){
	var studentData []Student
	if err := json.Unmarshal([]byte(Data), &studentData); err != nil {
		 fmt.Println(err)
	}
	for i,v := range studentData {
		if studentData[i].Id == Id {
			result, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
			}
			return string(result), 1
		}
	}

    return string(1), 0
}

// ============== Update Student ==========================================
func UpdateStudent(Data []byte, Id int,ChangedName string) (a string ){
	var studentData []Student
	if err := json.Unmarshal([]byte(Data), &studentData); err != nil {
		 fmt.Println(err)
	}
	for i:= range studentData {
		if studentData[i].Id == Id {
			studentData[i].Name =ChangedName
		}
	}
	result, err := json.Marshal(studentData)
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}

//==================== Delete Student =======================================
func DeleteStudent(Data []byte, Id int) (a string){
	var studentData []Student
	if err := json.Unmarshal([]byte(Data), &studentData); err != nil {
		 fmt.Println(err)
	}
	for i:= range studentData {
		if studentData[i].Id == Id {
			studentData=append(studentData[:i], studentData[i+1:]...)
			break
		}
	}
	result, err := json.Marshal(studentData)
	if err != nil {
		fmt.Println(err)
	}
	return string(result);
}
