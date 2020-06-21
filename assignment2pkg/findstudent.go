package assignment2pkg

// ==================== Find Student ===================================
func FindStudent(Data []Student, Id int) (a Student, b bool ){
	for i:= range Data {
		if Data[i].Id == Id {
			a =  Data[i];
			return a, true
		}
	}
    return a, false
}
