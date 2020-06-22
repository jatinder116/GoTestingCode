package assignment3pkg
import (
	"fmt"
	"encoding/json"
)


// ============== Update Student ==========================================
func UpdateStudent(Data []byte, Id int,ChangedName string) (a string ){
	var studentData []Studentjson
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
