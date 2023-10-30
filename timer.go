package main

import (
	"fmt"
	"time"
)

func runTimer(timeout int) {
	interval := 1 * time.Second

	for i := 0; i < timeout; i++ {
		progress := float64(i) / float64(timeout) * 100

		fmt.Print("\033[K")
		fmt.Printf("%ds/%ds %.2f%% completed", i, timeout, progress)
		fmt.Printf(" found total \033[32;5;3m%d crashes\033[0m\r", 1)

		time.Sleep(interval)
	}

	fmt.Printf("%ds/%ds %.2f%% completed\n", timeout, timeout, 100.0)
	fmt.Println("Task completed!")
}