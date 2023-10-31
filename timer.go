package main

import (
	"fmt"
	"io/ioutil"
	// "syscall"
	"time"
)

func runTimer(fuzzingPath string, timeout int) {
	interval := 1 * time.Second
	var crashes int

	// time.Sleep(interval)

	for i := 0; i < timeout; i++ {
		progress := float64(i) / float64(timeout) * 100

		files, err := ioutil.ReadDir(fuzzingPath + "/output/default/crashes")

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			crashes = 0
			if !file.IsDir() {
				if startsWith(file.Name(), "id:") { 
					crashes++
				}
			}
		}
		
		fmt.Print("\033[K")
		fmt.Printf("  [%ds/%ds %.2f%%] completed", i, timeout, progress)
		fmt.Printf(" found total \033[32;5;3m%d crashes\033[0m\r", crashes)

		time.Sleep(interval)
	}

	// interrupt <- syscall.SIGINT

	fmt.Printf("%ds/%ds %.2f%% completed\n", timeout, timeout, 100.0)
	fmt.Println("Task completed!")
}

func startsWith(str string, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}