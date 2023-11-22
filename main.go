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
		(*ConfigData).parseJSON(&configData, os.Args[1])
	}

	if checkFile(os.Args[2]) {
		(*RequestData).parseJSON(&requestData, os.Args[2])
	}

	termChan = make(chan os.Signal, 1)
	intChan = make(chan os.Signal, 1)

	printName()
	// runTimer("tmp", configData.Timeout)

	// printConfig()

	// fmt.Println("------------------------------------------------------------")

	// printRequest()

	// fmt.Println("------------------------------------------------------------")


	os.Setenv("LOGIN_COOKIE", "checkFuzz=true")
	// Login()

	// fmt.Println("------------------------------------------------------------")

	// testLogin()

	// test
	runAFL(initDir(0), 0)
	// fmt.Println(targetPoints)
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

func printName() {
	fmt.Println("\033[38;5;129m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m_\033[39m\033[38;5;93m_\033[39m\033[38;5;93m \033[39m\033[38;5;99m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m_\033[39m\033[38;5;63m_\033[39m\033[38;5;63m_\033[39m\033[38;5;63m_\033[39m\033[38;5;63m_\033[39m\033[38;5;69m_\033[39m\033[38;5;33m_\033[39m\033[38;5;33m \033[39m\033[38;5;33m_\033[39m\033[38;5;33m_\033[39m\033[38;5;33m_\033[39m\033[38;5;33m_\033[39m\033[38;5;33m_\033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;38m \033[39m\033[38;5;38m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;43m\033[39m	")
	fmt.Println("\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;99m\\\033[39m\033[38;5;63m \033[39m\033[38;5;63m\\\033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m/\033[39m\033[38;5;63m \033[39m\033[38;5;63m/\033[39m\033[38;5;69m_\033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m_\033[39m\033[38;5;33m|\033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m_\033[39m\033[38;5;33m_\033[39m\033[38;5;39m_\033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m_\033[39m\033[38;5;39m \033[39m\033[38;5;39m_\033[39m\033[38;5;39m_\033[39m\033[38;5;38m_\033[39m\033[38;5;38m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m \033[39m\033[38;5;44m_\033[39m\033[38;5;44m \033[39m\033[38;5;43m_\033[39m\033[38;5;49m_\033[39m\033[38;5;49m \033[39m\033[38;5;49m\033[39m	")
	fmt.Println("\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;93m \033[39m\033[38;5;99m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m\\\033[39m\033[38;5;63m \033[39m\033[38;5;63m\\\033[39m\033[38;5;63m \033[39m\033[38;5;63m/\033[39m\033[38;5;63m\\\033[39m\033[38;5;63m \033[39m\033[38;5;63m/\033[39m\033[38;5;69m \033[39m\033[38;5;33m/\033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m|\033[39m\033[38;5;33m \033[39m\033[38;5;33m|\033[39m\033[38;5;33m \033[39m\033[38;5;33m|\033[39m\033[38;5;33m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m_\033[39m\033[38;5;39m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;38m|\033[39m\033[38;5;38m_\033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m/\033[39m\033[38;5;44m_\033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m/\033[39m\033[38;5;44m \033[39m\033[38;5;44m_\033[39m\033[38;5;44m \033[39m\033[38;5;43m\\\033[39m\033[38;5;49m \033[39m\033[38;5;49m'\033[39m\033[38;5;49m_\033[39m\033[38;5;49m_\033[39m\033[38;5;49m|\033[39m\033[38;5;49m\033[39m	")
	fmt.Println("\033[38;5;93m \033[39m\033[38;5;99m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m\\\033[39m\033[38;5;63m \033[39m\033[38;5;63mV\033[39m\033[38;5;63m \033[39m\033[38;5;69m \033[39m\033[38;5;33mV\033[39m\033[38;5;33m \033[39m\033[38;5;33m/\033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m|\033[39m\033[38;5;33m \033[39m\033[38;5;33m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m_\033[39m\033[38;5;39m|\033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m|\033[39m\033[38;5;38m_\033[39m\033[38;5;38m|\033[39m\033[38;5;44m \033[39m\033[38;5;44m|\033[39m\033[38;5;44m/\033[39m\033[38;5;44m \033[39m\033[38;5;44m/\033[39m\033[38;5;44m \033[39m\033[38;5;44m/\033[39m\033[38;5;44m \033[39m\033[38;5;44m/\033[39m\033[38;5;44m \033[39m\033[38;5;43m \033[39m\033[38;5;49m_\033[39m\033[38;5;49m_\033[39m\033[38;5;49m/\033[39m\033[38;5;49m \033[39m\033[38;5;49m|\033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m\033[39m	")
	fmt.Println("\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;69m\\\033[39m\033[38;5;33m_\033[39m\033[38;5;33m/\033[39m\033[38;5;33m\\\033[39m\033[38;5;33m_\033[39m\033[38;5;33m/\033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m_\033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m|\033[39m\033[38;5;39m_\033[39m\033[38;5;39m|\033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;38m \033[39m\033[38;5;38m\\\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m,\033[39m\033[38;5;44m_\033[39m\033[38;5;44m/\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m_\033[39m\033[38;5;44m/\033[39m\033[38;5;44m_\033[39m\033[38;5;43m_\033[39m\033[38;5;49m_\033[39m\033[38;5;49m\\\033[39m\033[38;5;49m_\033[39m\033[38;5;49m_\033[39m\033[38;5;49m_\033[39m\033[38;5;49m|\033[39m\033[38;5;49m_\033[39m\033[38;5;49m|\033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;48m \033[39m\033[38;5;48m\033[39m	")
	fmt.Println("\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;63m \033[39m\033[38;5;69m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;33m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;39m \033[39m\033[38;5;38m \033[39m\033[38;5;38m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;44m \033[39m\033[38;5;43m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;49m \033[39m\033[38;5;48m \033[39m\033[38;5;48m \033[39m\033[38;5;48m \033[39m\033[38;5;48m \033[39m\033[38;5;48m\033[39m	")
}
