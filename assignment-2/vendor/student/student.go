package student

//=========== Add Student ===========================
func AddStudent(Data []Student, Name Student) (a[]Student ){
	a = append(Data,Name);
	return a;
}


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

// ============== Update Student ==========================================
func UpdateStudent(Data []Student, Id int,ChangedName string) (a[]Student ){
	for i:= range Data {
		if Data[i].Id == Id {
			Data[i].Name = ChangedName;
		}
	}
	return Data;
}
