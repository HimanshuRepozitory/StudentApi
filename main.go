package main

import (
	"SampleAPI/controller"
	"SampleAPI/db"
	"SampleAPI/helper"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

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

	var wg sync.WaitGroup
	//Gracefully shutdown the server

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	wg.Add(1)
	go helper.StartServer(server, &wg)

	//making a channel which recieve  interrupt signal from OS
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	//waiting for interrupt signal
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(1)
	go helper.ShutDownServerGracefully(ctx, &wg, server)

	wg.Wait() //blocks until all WaitGroup counter is zero and waits

}
