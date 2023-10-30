package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func runTimer(fuzzingPath string, targetPoint string, timeout int) {
	interval := 1 * time.Second
	crashes := 0
	
	for i := 0; i < timeout; i++ {
		progress := float64(i) / float64(timeout) * 100

		files, err := ioutil.ReadDir(fuzzingPath + "/input/seeds/" + targetPoint)

		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				if startsWith(file.Name(), "id:") { 
					crashes++
				}
			}
		}
		
		fmt.Print("\033[K")
		fmt.Printf("[%ds/%ds %.2f%%] completed", i, timeout, progress)
		fmt.Printf(" found total \033[32;5;3m%d crashes\033[0m\r", crashes)

		time.Sleep(interval)
	}

	fmt.Printf("%ds/%ds %.2f%% completed\n", timeout, timeout, 100.0)
	fmt.Println("Task completed!")
}


func startsWith(str string, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}