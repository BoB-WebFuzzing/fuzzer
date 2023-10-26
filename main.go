package main

import (

)

var configData config
var requestData request

func main() {
	(*config).parseJSON(&configData, "json/config.json")
	printConfig()

	(*request).parseJSON(&requestData, "json/request_data.json")
	printRequest()
}