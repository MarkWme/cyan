package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type version struct {
	Application string
	Version     string
}

var ready bool = true

func getVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getVersionValue())
	log.Println("getVersion endpoint:", getVersionValue())
}

func podTerminate(w http.ResponseWriter, r *http.Request) {
	log.Println("podTerminate endpoint: Starting 30 second waiting period ...")
	ready = false
	time.Sleep(30 * time.Second)
	log.Println("Waiting period complete")
}

func podReady(w http.ResponseWriter, r *http.Request) {
	log.Println("podReady endpoint:", ready)
	if ready {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/api/getVersion", getVersion)
	http.HandleFunc("/api/podTerminate", podTerminate)
	http.HandleFunc("/api/podReady", podReady)
	srv := &http.Server{
		Addr:              ":3000",
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	log.Println("Starting HTTP server on port 3000")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func getVersionValue() version {
	return version{"Simple API Server", "1.1.0"}
}
