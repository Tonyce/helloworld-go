// package main

// import (
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// )

// func main() {
// 	done := make(chan os.Signal, 1)
// 	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

// 	<-done
// 	log.Print("Server Stopped")
// }

////////////////////////////////////////////////////////////////////////////

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux" // need to use dep for package management
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	time.Sleep(time.Second * 5)
	w.Write([]byte("Test is what we usually do:" + strconv.Itoa(os.Getpid())))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestEndpoint).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	s := <-done
	log.Printf("Server Stopped %+v", s)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
