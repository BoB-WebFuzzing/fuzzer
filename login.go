package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var sessid string

func Login() {
	loginData := configData.Login
	var url string

	if loginData.Port < 0 {
		url = loginData.URL
	} else {
		url = loginData.URL+":"+strconv.Itoa(loginData.Port)
	}

	req, err := http.NewRequest(loginData.Method, url, bytes.NewBufferString(loginData.PostData))
	
	if err != nil {
		panic(err)
	}

	for key, value := range loginData.PositiveHeaders {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	cookies := res.Cookies()
	for _, cookie := range cookies {
		fmt.Println(cookie.Name, ":", cookie.Value)

		if matched, _ := regexp.MatchString(`(?i)sess`, cookie.Name); matched {
			fmt.Println("	Cookie Name:", cookie.Name)
			fmt.Println("	Cookie Value:", cookie.Value)

			os.Setenv("LOGIN_COOKIE", cookie.Name + ":" + cookie.Value)

			sessid = cookie.Value
		}
	}
}

func testLogin() {
	req, err := http.NewRequest("GET", "https://dreamhack.io/", nil)
	
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", "sessionid="+sessid)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}