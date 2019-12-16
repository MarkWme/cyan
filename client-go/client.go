package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	apiServer := os.Getenv("API_SERVER")
	var requestsPerSecond int = 1
	requestsPerSecond, _ = strconv.Atoi(os.Getenv("REQUESTS_PER_SECOND"))
	getVersionURL := apiServer + "/api/getVersion"
	log.Printf("Starting to poll %s at %d requests per second", getVersionURL, requestsPerSecond)
	for {
		time.Sleep(time.Duration(1000/requestsPerSecond) * time.Millisecond)
		response, err := http.Get(getVersionURL)
		if err != nil {
			log.Printf("HTTP request failed %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			log.Printf(string(data))
		}
	}
}
