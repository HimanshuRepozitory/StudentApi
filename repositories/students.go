package repositories

import (
	models "SampleAPI/my_models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateStudentData(c *gin.Context, insertData *models.Student) error {
	fmt.Println("insertData in repositories : ", insertData)

	err := insertData.Insert(c, boil.GetContextDB(), boil.Infer())
	fmt.Println("the error in repositories is  : ", err)
	if err != nil {
		return err
	}
	return nil
}

func FindById(c *gin.Context, id int) (*models.Student, error) {
	student, err := models.Students(qm.Where("id = ?", id)).One(c, boil.GetContextDB())
	if err != nil {
		return nil, err
	}
	return student, nil
}

func UpdateStudentData(c *gin.Context, student *models.Student) (int, error) {
	rowsAffected, err := student.Update(c, boil.GetContextDB(), boil.Infer())
	if err != nil {
		fmt.Println("error in updating the data in database in repository : ", err)
		return 0, err
	}
	fmt.Println("updating data in database in repository : ", rowsAffected)
	return int(rowsAffected), nil
}

func DeleteStudentById(c *gin.Context, student *models.Student) (int, error) {
	rowAffected, err := student.Delete(c, boil.GetContextDB())

	if err != nil {
		fmt.Println("error in deleting student data : ", err)
		return 0, err
	}

	fmt.Println("Student data deleted successfully... : ", rowAffected)
	return int(rowAffected), nil

}

func GetAllStudentData(c *gin.Context) (models.StudentSlice, error) {
	studentData, err := models.Students().All(c, boil.GetContextDB())
	if err != nil {
		return nil, err
	}

	return studentData, nil
}
