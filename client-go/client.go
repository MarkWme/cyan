package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	apiServer := os.Getenv("API_SERVER")
	getVersionURL := apiServer + "/api/getVersion"
	log.Println("Starting to poll " + getVersionURL)
	for {
		time.Sleep(1 * time.Second)
		response, err := http.Get(getVersionURL)
		if err != nil {
			log.Printf("HTTP request failed %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			log.Println(string(data))
		}
	}
}
