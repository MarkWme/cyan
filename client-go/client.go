package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	apiServer := os.Getenv("API_SERVER")
	getVersionURL := apiServer + "/api/getVersion"
	fmt.Println("Starting to poll " + getVersionURL)
	for {
		time.Sleep(1 * time.Second)
		response, err := http.Get(getVersionURL)
		if err != nil {
			fmt.Printf("HTTP request failed %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
}