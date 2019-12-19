package main

import (
	"encoding/json"
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
	log.Println("Starting HTTP server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getVersionValue() version {
	return version{"Cyan API Server", "1.0.13"}
}
