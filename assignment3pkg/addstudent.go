package assignment3pkg

import (
	"fmt"
	"gotestcode/assignment3pkg/structpkg"
	"encoding/json"
)
type Studentjson = structpkg.Student

//=========== Add Student ===========================
func AddStudent(Data []byte, Newdata Studentjson) (a string) {
	var studentData []Studentjson
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