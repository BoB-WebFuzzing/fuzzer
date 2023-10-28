package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

type fuzzCampaignStatus struct {
	TrialStart		string			`json:"trial_start"`
	TrialComplete	bool			`json:"trial_complete"`
	Targets			[]fuzzTarget	`json:"targets"`
}

type fuzzTarget struct {
	TargetPath			string			`json:"target_path"`
	Requests			[]string		`json:"requests"`
	Methods				map[string]int	`json:"methods"`
	IsSoapAction		bool			`json:"is_soapaction"`
	LastCompletedTrial	int				`json:"last_completed_trial"`
	LastCompletedRefuzz	int				`json:"last_completed_refuzz"`
}

func runAFL(fuzzingPath string) {
	cmd := exec.Command("sh", fuzzingPath + "/run.sh")

	go exitAFL(cmd)

	output, _ := cmd.CombinedOutput()

	os.WriteFile(fuzzingPath + "/output/fuzzer.log", output, 0644)
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
	createFuzzStat(fuzzingDir)
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
	scriptContent := "#!/bin/sh\n\n"

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

func createFuzzStat(fuzzingPath string) {
	var fuzzStat fuzzCampaignStatus
	uniqCheck := make(map[string]int)
	targetIndex := 0
	
	fuzzStat.TrialStart = time.Now().Format("2006_01_02_15_04")
	fuzzStat.TrialComplete = false
	fuzzStat.Targets = []fuzzTarget{}

	for key, value := range requestData.RequestsFound {
		targetURL := strings.Split(value.URL, "?")[0]
		method := strings.Split(key, " ")[0]
		_, exist := uniqCheck[targetURL]

		if exist {
			fuzzStat.Targets[uniqCheck[targetURL]].Requests = append(fuzzStat.Targets[uniqCheck[targetURL]].Requests, key)
			fuzzStat.Targets[uniqCheck[targetURL]].Methods[method]++
		} else {
			uniqCheck[targetURL] = targetIndex
			targetIndex++

			tempFuzzTarget := fuzzTarget{
				Methods: make(map[string]int),
			}

			tempFuzzTarget.TargetPath = strings.Split(value.URL, "?")[0]
			tempFuzzTarget.Requests = append(tempFuzzTarget.Requests, key)
			tempFuzzTarget.Methods[method] = 1

			fuzzStat.Targets = append(fuzzStat.Targets, tempFuzzTarget)
		}
	}

	data, err := json.MarshalIndent(fuzzStat, "", "	")
	data = []byte(strings.ReplaceAll(string(data), "\\u0026", "&"))

	if err != nil {
		panic(err)
	}

	file, err := os.Create(fuzzingPath + "/output/fuzz_stat.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	
	_, err = file.Write(data)

	if err != nil {
		panic(err)
	}
}