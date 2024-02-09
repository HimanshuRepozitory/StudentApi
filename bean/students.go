package bean

type StudentData struct {
	Roll_number      string `json:"roll_number"`
	Student_name     string `json:"student_name"`
	Email            string `json:"email"`
	Admission_number string `json:"admission_number"`
}

type UpdateStudentData struct {
	ID               int     `json:"id"`
	Roll_number      string `json:"roll_number"`
	Student_name     string `json:"student_name"`
	Email            string `json:"email"`
	Admission_number string `json:"admission_number"`
}


