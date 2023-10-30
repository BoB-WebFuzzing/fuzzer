package main

import (
	"fmt"
	"os"

	// "github.com/fatih/color"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	if checkFile(os.Args[1]) {
		(*config).parseJSON(&configData, "json/config.json")
	}

	if checkFile(os.Args[2]) {
		(*request).parseJSON(&requestData, "json/request_data.json")
	}

	// printConfig()

	fmt.Println("------------------------------------------------------------")

	// printRequest()

	fmt.Println("------------------------------------------------------------")

	// Login()

	// fmt.Println("------------------------------------------------------------")

	// testLogin()


	// runTimer(120)


	// for i := 0; i < configData.Cores; i++ {
	// 	initDir(i)
	// }

	// test
	runAFL(initDir(0), 0)
}

func usage() {
	fmt.Println("Usage : fuzzer <path of config file> <path of request data file>")
	fmt.Println("Example : fuzzer config.json request_data.json")

	os.Exit(-1)
}

func checkFile(fileName string) bool {
    _, err := os.Stat(fileName)

    if os.IsNotExist(err) {
        fmt.Printf("%v File does not exist. Please check your path.\n", fileName)
		panic(err)
	}

	return true
}