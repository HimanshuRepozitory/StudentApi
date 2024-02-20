package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func StartServer(server *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()
	AcceptReq = true
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

var RunningRequests int32
var RunningMutex sync.Mutex

func RunningReq() int32 {

	RunningMutex.Lock()
	defer RunningMutex.Unlock()
	return RunningRequests
}

func DecrementRunningReq() {
	RunningMutex.Lock()
	defer RunningMutex.Unlock()
	RunningRequests--
}

func IncrementRunningReq() {
	RunningMutex.Lock()
	defer RunningMutex.Unlock()
	RunningRequests++
}

func StartAndShutDownServer(server *http.Server, wg *sync.WaitGroup) {

	wg.Add(1)
	go StartServer(server, wg)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		<-quit
		// SetAcceptReq()
		AcceptReq = false
		// Stop accepting new requests
		fmt.Println("Quitting the server...")
		defer wg.Done()
		// Wait for active requests to complete
		fmt.Println("Waiting for the Active requests...")
		// fmt.Println("active Req : ",RunningRequests)
		for {

			if RunningReq() == 0 {
				fmt.Println("Server Shutdown Successfully!!")
				server.Close()
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}

var AcceptReq bool

