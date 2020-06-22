package assignment3pkg
import (
	"fmt"
	"encoding/json"
)


// ==================== Find Student ===================================
func FindStudent(Data []byte, Id int) (a string, b int ){
	var studentData []Studentjson
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
