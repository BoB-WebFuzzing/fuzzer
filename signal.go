package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var interrupt chan os.Signal

func exitAFL(c *exec.Cmd) {
	interrupt = make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	process := c.Process
	err := process.Signal(syscall.SIGINT)

	if err != nil {
		fmt.Println("\nFailed to send SIGINT:\n\t", err)
	}

	fmt.Println("\nSIGTERM received. Exiting...")
}

func exitFuzzer(c *exec.Cmd) {
	interrupt = make(chan os.Signal, 1)

	signal.Notify(interrupt, syscall.SIGINT)

	<-interrupt

	process := c.Process
	err := process.Signal(syscall.SIGINT)

	if err != nil {
		fmt.Println("\nFailed to send SIGINT:\n\t", err)
	}

	fmt.Println("\nSIGINT received. Exiting...")

	os.Exit(-1)
}