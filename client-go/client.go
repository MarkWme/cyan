package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		response, err := http.Get("http://localhost:3000/api/getVersion")
		if err != nil {
			fmt.Printf("HTTP request failed %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
}