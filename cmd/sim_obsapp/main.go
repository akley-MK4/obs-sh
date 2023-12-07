package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	listenAddr   = "127.0.0.1:8080"
	readTimeout  = 10 * time.Second
	writeTimeout = 15 * time.Second
)

func main() {
	httpSvr := &http.Server{Addr: listenAddr,
		Handler:      &Handler{},
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	go func() {
		time.Sleep(time.Minute * 1)
		log.Println("Exit the SimObsApp process")
		os.Exit(0)
	}()

	if err := httpSvr.ListenAndServe(); err != nil {
		log.Println("Failed to listen http server, ", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

type Handler struct {
}

func (t *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	log.Printf("Access path: %v\n", path)

	w.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(w, fmt.Sprintf(`{"code": %d}`, 1))
	if err != nil {
		log.Println("io.WriteString failed, ", err.Error())
	}

}
