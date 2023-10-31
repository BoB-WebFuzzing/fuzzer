package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var termChan chan os.Signal
var intChan chan os.Signal
var resetChan chan struct{}

func exitAFL(c *exec.Cmd) {
	signal.Notify(termChan, syscall.SIGTERM)

	<-termChan

	process := c.Process
	err := process.Signal(syscall.SIGINT)

	if err != nil {
		fmt.Println("\nFailed to send SIGINT:\n\t", err)
	}

	fmt.Println("\nSIGTERM received. Exiting...")

	close(resetChan)
}

func exitFuzzer(c *exec.Cmd) {
	signal.Notify(intChan, os.Interrupt, syscall.SIGINT)

	<-intChan

	process := c.Process
	err := process.Signal(syscall.SIGINT)

	if err != nil {
		fmt.Println("\nFailed to send SIGINT:\n\t", err)
	}

	fmt.Println("\nSIGINT received. Exiting...")

	close(resetChan)

	os.Exit(-1)
}