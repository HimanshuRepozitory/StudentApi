package usecases

import (
	"SampleAPI/bean"
	"SampleAPI/helper"
	models "SampleAPI/my_models"
	"SampleAPI/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context, data *bean.StudentData) (*models.Student, error) {
	fmt.Println("data : ", data)
	insertData := &models.Student{
		RollNumber:      data.Roll_number,
		StudentName:     data.Student_name,
		AdmissionNumber: data.Admission_number,
		Email:           data.Email,
	}

	fmt.Printf("insertData in usecases=====================> %+v \n: ", insertData)

	err := repositories.CreateStudentData(c, insertData)
	if err != nil {
		return nil, err
	}
	return insertData, nil

}

func UpdateStudentData(c *gin.Context, data *bean.UpdateStudentData) (int, error) {
	fmt.Println("the data is :", data)

	student, err := repositories.FindById(c, data.ID)
	if err != nil {
		fmt.Println("error in finding student by id : ", err)
		return 0, err
	}
	fmt.Println("the student data is  : ", student)

	helper.CheckFields(data, student)

	rowsAffected, err := repositories.UpdateStudentData(c, student)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteStudentById(c *gin.Context, userId int) (int, error) {
	student, err := repositories.FindById(c, userId)
	if err != nil {
		return 0, err
	}

	rowAffected, err := repositories.DeleteStudentById(c, student)
	if err != nil {
		return 0, err
	}

	return rowAffected, nil
}

func GetStudentById(c *gin.Context, id int) (*models.Student, error) {
	student, err := repositories.FindById(c, id)
	if err != nil {
		return nil, err
	}

	return student, err

}

func GetAllStudentData(c *gin.Context) (models.StudentSlice, error) {
	student, err := repositories.GetAllStudentData(c)
	if err != nil {
		return nil, err
	}

	return student, err
}
