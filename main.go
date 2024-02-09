package main

import (
	"SampleAPI/controller"
	"SampleAPI/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	//DATA BASE CONNECTION
	DB, err := db.DbConnection()
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}

	boil.SetDB(DB)
	fmt.Println("Database connected successfully")

	// MIGRATIONS TEST
	err = db.RunMigrations()
	if err != nil {
		fmt.Println("Failed to migrate up")
		log.Fatal(err)

	}

	//SERVER AND ROUTES
	router := gin.Default()
	router.POST("/admission", controller.CreateStudent)
	router.GET("/getStudents", controller.GetAllStudentData)
	router.GET("/getbyID/:id", controller.GetStudentById)
	router.PUT("/updateStudentData", controller.UpdateStudentData)
	router.DELETE("/deleteStudent/:id", controller.DeleteStudentById)

	router.GET("/", controller.Default)

	router.Run(":8080")

}
