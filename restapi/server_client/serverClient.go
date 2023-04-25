package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal("Unable to connect to the url : ", err)
	}
	resBodyDataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unable to read the response : ", err)
	}
	fmt.Println(string(resBodyDataBytes), resp.Status, resp.StatusCode)
}

// func main() {
// 	client := http.Client{}
// 	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
// 	if err != nil {
// 		fmt.Println("Unable to read the response : ", err)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("Unable to connect to the url : ", err)
// 	}
// 	resBodyDataBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Unable to read the response : ", err)
// 	}
// 	fmt.Println(string(resBodyDataBytes), resp.Status, resp.StatusCode)
// }
