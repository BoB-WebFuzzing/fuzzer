package main

import (
)

var configData config

func main() {
	parseJSON("config.json")
	printConfig()
}