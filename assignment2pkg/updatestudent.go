package assignment2pkg

// ============== Update Student ==========================================
func UpdateStudent(Data []Student, Id int,ChangedName string) (a[]Student ){
	for i:= range Data {
		if Data[i].Id == Id {
			Data[i].Name = ChangedName;
		}
	}
	return Data;
}
