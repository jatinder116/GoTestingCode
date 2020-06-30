package student

//StudentStructure: structure for student data
type Student struct {
	Id	int 		`json:"id" validate:"required,min=1,max=200"`
	Name	string 		`json:"name" validate:"required"`
	Email   string      `json:"email" validate:"required,email"`
	City	string 		`json:"city"`
	Gender  string    `json:"gender" validate:"required"`
}
