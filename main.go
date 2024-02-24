package main

import (
	"SampleAPI/controller"
	"SampleAPI/db"
	"SampleAPI/server"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	DbInstance, err := db.DbConnection()
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}
	boil.SetDB(DbInstance)
	fmt.Println("Database connected successfully")

	err = db.RunMigrations()
	if err != nil {
		fmt.Println("Failed to migrate up")
		log.Fatal(err)

	}

	router := gin.Default()

	controller.NewStudentsRoutes(router)

	

	//Gracefully shutdown the server
	var wg sync.WaitGroup
	ser := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.StartAndShutDownServer(ser, &wg)

}
