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
	requestsPerSecond, _ := strconv.Atoi(os.Getenv("REQUESTS_PER_SECOND"))
	getVersionURL := apiServer + "/api/getVersion"
	log.Printf("Starting to poll " + getVersionURL)
	for {
		time.Sleep(time.Duration(1/requestsPerSecond) * time.Second)
		response, err := http.Get(getVersionURL)
		if err != nil {
			log.Printf("HTTP request failed %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			log.Printf(string(data))
		}
	}
}
