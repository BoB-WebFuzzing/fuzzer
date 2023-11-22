package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func runTimer(fuzzingPath string, timeout int) {
	timerChan := make(chan os.Signal, 1)
	interval := 1 * time.Second
	var crashes int
	var paths int

	signal.Notify(timerChan, syscall.SIGINT, syscall.SIGTERM)

	time.Sleep(interval)

	for i := 0; i < timeout; i++ {
		select {
		case <-timerChan:
			fmt.Println("\nInterrupt signal received. Exiting...")
			return
		default:
			progress := float64(i) / float64(timeout) * 100
			bar := strings.Repeat("=", int(float64(i) / float64(timeout) * 30))
			spaces := strings.Repeat(" ", 30 - int(float64(i) / float64(timeout) * 30))

			files, err := ioutil.ReadDir(fuzzingPath + "/output/crashes")

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

			files, err = ioutil.ReadDir(fuzzingPath + "/output/queue")

			if err != nil {
				panic(err)
			}

			paths = 0
			for _, file := range files {
				if !file.IsDir() {
					if startsWith(file.Name(), "id:") {
						paths++
					}
				}
			}
			
			fmt.Print("\033[K")

			if progress < 33 {
				fmt.Printf("  [\033[31m%v>%v\033[0m][%ds/%ds %.2f%%] completed", bar, spaces, i, timeout, progress)
			} else if progress < 66 {
				fmt.Printf("  [\033[38;5;208m%v>%v\033[0m][%ds/%ds %.2f%%] completed", bar, spaces, i, timeout, progress)
			} else if progress < 100 {
				fmt.Printf("  [\033[33m%v>%v\033[0m][%ds/%ds %.2f%%] completed", bar, spaces, i, timeout, progress)
			}

			fmt.Printf(" found %d paths and total \033[32;5;3m%d crashes\033[0m\r", paths, crashes)

			time.Sleep(interval)
		}
	}

	termChan <- syscall.SIGTERM

	fmt.Printf("  [\033[32m==============================>\033[0m][%ds/%ds %.2f%%] completed\n", timeout, timeout, 100.0)
	fmt.Println("Task completed!\n")
}

func startsWith(str string, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}
