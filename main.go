package main

import (
	"fmt"
)

var configData config
var requestData request

func main() {
	(*config).parseJSON(&configData, "json/config.json")
	printConfig()

	fmt.Println("------------------------------------------------------------")

	(*request).parseJSON(&requestData, "json/request_data.json")
	printRequest()
}