package helper

import (
	"SampleAPI/bean"
	models "SampleAPI/my_models"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
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

func StartServer(server *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func ShutDownServerGracefully(ctx context.Context, wg *sync.WaitGroup, server *http.Server) {
	defer wg.Done()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	fmt.Println("Server gracefully stopped")
}
