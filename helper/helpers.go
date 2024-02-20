package helper

import (
	"SampleAPI/bean"
	models "SampleAPI/my_models"
)

func CheckFields(data *bean.UpdateStudentData, student *models.Student) {
	if data.Student_name != "" {
		student.StudentName = data.Student_name
	}
	if data.Admission_number != "" {
		student.AdmissionNumber = data.Admission_number
	}
	if data.Email != "" {
		student.Email = data.Email
	}
	if data.Roll_number != "" {
		student.RollNumber = data.Roll_number
	}
}
