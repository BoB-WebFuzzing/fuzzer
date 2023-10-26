package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	Testname		string	`json:"testname"`
	AFLPath			string	`json:"afl_path"`
	TargetBinary    string	`json:"target_binary"`
	BaseURL         string	`json:"base_url"`
	BasePort        int		`json:"base_port"`
	Timeout         int		`json:"timeout"`
	LdLibraryPath   string	`json:"ld_library_path"`
	LdPreload       string	`json:"ld_preload"`
	Memory          string	`json:"memory"`
	FirstCrash      bool	`json:"first_crash"`
	Cores           int		`json:"cores"`
	Direct          struct {
		URL              string				`json:"url"`
		Port             int				`json:"port"`
		PostData         string				`json:"postData"`
		GetData          string				`json:"getData"`
		PositiveHeaders  map[string]string	`json:"positiveHeaders"`
		PositiveBody     string				`json:"positiveBody"`
		Method           string				`json:"method"`
		LoginSessionCookie string			`json:"loginSessionCookie"`
	}	`json:"direct"`
}

func (c config) parseJSON(fileName string) {
	data, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()
	
	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &configData)
}

func printConfig() {
	fmt.Printf("%-20v %v\n", "Testname:", configData.Testname)
	fmt.Printf("%-20v %v\n", "AFLPath:", configData.AFLPath)
	fmt.Printf("%-20v %v\n", "TargetBinary:", configData.TargetBinary)
	fmt.Printf("%-20v %v\n", "BaseURL:", configData.BaseURL)
	fmt.Printf("%-20v %v\n", "BasePort:", configData.BasePort)
	fmt.Printf("%-20v %v\n", "Timeout:", configData.Timeout)
	fmt.Printf("%-20v %v\n", "LdLibraryPath:", configData.LdLibraryPath)
	fmt.Printf("%-20v %v\n", "LdPreload:", configData.LdPreload)
	fmt.Printf("%-20v %v\n", "Memory:", configData.Memory)
	fmt.Printf("%-20v %v\n", "FirstCrash:", configData.FirstCrash)
	fmt.Printf("%-20v %v\n", "Cores:", configData.Cores)
	fmt.Printf("%-20v\n", "Direct:")
	fmt.Printf("	%-20v %v\n", "URL:", configData.Direct.URL)
	fmt.Printf("	%-20v %v\n", "Port:", configData.Direct.Port)
	fmt.Printf("	%-20v %v\n", "PostData:", configData.Direct.PostData)
	fmt.Printf("	%-20v %v\n", "GetData:", configData.Direct.GetData)
	fmt.Printf("	%-20v %v\n", "PositiveHeaders:", configData.Direct.PositiveHeaders)
	fmt.Printf("	%-20v %v\n", "PositiveBody:", configData.Direct.PositiveBody)
	fmt.Printf("	%-20v %v\n", "Method:", configData.Direct.Method)
	fmt.Printf("	%-20v %v\n", "LoginSessionCookie:", configData.Direct.LoginSessionCookie)
}