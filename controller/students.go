package controller

import (
	"SampleAPI/bean"
	"SampleAPI/server"
	"SampleAPI/usecases"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func StopAcceptingReqDuringShutdown(c *gin.Context) {

	// fmt.Println("Entering the middleware ============>>>>>>")
	// fmt.Println("Accept Req -==================-> ", server.AcceptReq)
	if !server.AcceptReq {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("can't accept new requests"))
		return
	}
	c.Next()
}

func NewStudentsRoutes(router *gin.Engine) {
	router.Use(StopAcceptingReqDuringShutdown)
	router.POST("/students", Create)
	router.GET("/students", Students)
	router.GET("/students/:id", Student)
	router.PATCH("/students", Update)
	router.DELETE("/students/:id", Delete)
	router.GET("/", Default)
}

func Default(c *gin.Context) {

	server.IncrementRunningReq()
	// fmt.Println("active req --------->>>>>>", server.RunningReq())
	time.Sleep(time.Second * 20)
	server.DecrementRunningReq()
	c.JSON(http.StatusOK, gin.H{"data": "the server is runnung on port number 8080 ..."})
}

func Create(c *gin.Context) {
    server.IncrementRunningReq()
	data := bean.StudentData{}
	if err := c.BindJSON(&data); err != nil {
		fmt.Println("Error in binding json : ", err)
		c.JSON(http.StatusConflict, gin.H{"message": "error in binding data in controller ", "error": err})
		return
	}

	student, err := usecases.CreateStudent(c, &data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in creating student INTERNAL SERVER ERROR!!", "error": err.Error()})
		return
	}
    server.DecrementRunningReq()
	c.JSON(http.StatusAccepted, gin.H{"message": "Student Admiteed successfully!!!", "success": true, "data": student})
}

func Update(c *gin.Context) {

	data := bean.UpdateStudentData{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "error in marhalling!!", "error": err})
		fmt.Println("error is marshalling", err)
		return
	}

	rowsAffected, err := usecases.UpdateStudentData(c, &data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in updating student INTERNAL SERVER ERROR!!", "error": err.Error(), "rowsAffected": rowsAffected})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Student Updated successfully!!!", "success": true, "data": data, "rowsAffected": rowsAffected})
}

func Delete(c *gin.Context) {

	userId := c.Param("id")
	userid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("type conversion error")
		return
	}

	rowAffected, err := usecases.DeleteStudentById(c, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in deleting student INTERNAL SERVER ERROR!!", "error": err.Error(), "rowsAffected": rowAffected})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Student deleted successfully!!!", "success": true, "rowsAffected": rowAffected})

}

func Student(c *gin.Context) {

	userId := c.Param("id")
	userid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("type conversion error")
		return
	}

	student, err := usecases.GetStudentById(c, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in getting student data INTERNAL SERVER ERROR!!", "error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "successfully get student Data", "success": true, "data": student})
}

func Students(c *gin.Context) {

	student, err := usecases.GetAllStudentData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in getting all student data INTERNAL SERVER ERROR!!", "error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "successfully get student Data", "success": true, "data": student})
}
