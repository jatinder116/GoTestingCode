package assignment3pkg
import (
	"fmt"
	"encoding/json"
)


//==================== Delete Student =======================================
func DeleteStudent(Data []byte, Id int) (a string){
	var studentData []Studentjson
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
