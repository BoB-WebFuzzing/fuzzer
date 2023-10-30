package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func exitAFL(c *exec.Cmd) {
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt

	process := c.Process
	err := process.Signal(syscall.SIGINT)

	if err != nil {
		fmt.Println("\nFailed to send SIGINT:", err)
	}

	fmt.Println("\nSIGINT received. Exiting...")
}