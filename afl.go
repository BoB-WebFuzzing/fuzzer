package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func runAFL(fuzzingPath string) {
	cmd := exec.Command("sh", fuzzingPath + "/run.sh")

	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))


	if err != nil {
		// panic(err)
	}

	fmt.Println(string(output))
}

func initDir(i int) {
	fuzzingDir := fmt.Sprintf("fuzzing-%d", i)
	inputDir := fuzzingDir + "/input"
	outputDir := fuzzingDir + "/output"

	mkdir(fuzzingDir)
	mkdir(inputDir)
	mkdir(outputDir)

	createScript(fuzzingDir)
	createDict(fuzzingDir)
}

func mkdir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, os.ModePerm)

		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Directory already exists:", dirName)
	}
}

func createScript(fuzzingPath string) {
	scriptPath := fuzzingPath + "/run.sh"
	file, err := os.Create(scriptPath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var targets []string
	scriptContent := `#!/bin/sh

`

	for key := range requestData.RequestsFound {
		targets = append(targets, strings.Split(key, " ")[1])
	}

	scriptContent += configData.AFLPath + "afl-fuzz"
	scriptContent += " -i " + fuzzingPath + "/input"
	scriptContent += " -o " + fuzzingPath + "/output"
	scriptContent += " -m " + configData.Memory
	scriptContent += " -x " + fuzzingPath + "/input/dict.txt -- "
	scriptContent += configData.TargetBinary
	scriptContent += targets[0]

	_, err = file.WriteString(scriptContent)
	if err != nil {
		panic(err)
	}

	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		panic(err)
	}
}

func createDict(fuzzingPath string) {
	dictPath := fuzzingPath + "/input/dict.txt"
	var dictContent string
	file, err := os.Create(dictPath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for i, param := range requestData.InputSet {
		dictContent += fmt.Sprintf("string_%d=\"%v\"\n", i, strings.ReplaceAll(url.QueryEscape(param), "%", "\\x"))
	}

	_, err = file.WriteString(dictContent)
	if err != nil {
		panic(err)
	}
}
