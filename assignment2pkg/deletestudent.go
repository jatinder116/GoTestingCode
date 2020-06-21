package assignment2pkg

// ==================== Delete Student =======================================
func DeleteStudent(Data []Student, Id int) (a[]Student ){
	for i:= range Data {
		if Data[i].Id == Id {
			a = append(Data[:i], Data[i+1:]...)
			break
		}
	}
	return a;
}
