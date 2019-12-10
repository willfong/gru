package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func errCheck(e error, errmsg string) {
	if e != nil {
		fmt.Println(errmsg)
		panic(e)
	}
}

func main() {
	if len(os.Args) != 3 {
		cmdName := os.Args[0]
		fmt.Println("Please run using: ", cmdName, " <github username> <path to authorized keys>")
		os.Exit(1)
	}
	githubUsername := os.Args[1]
	authorizedKeysPath := os.Args[2]

	githubURL := fmt.Sprintf("https://github.com/%s.keys", githubUsername)
	fmt.Println("Using GitHub: ", githubURL)
	resp, err := http.Get(githubURL)
	errCheck(err, "Error fetching GitHub Keys")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errCheck(err, "Error reading GitHub Keys")

	sshKeys := string(body)
	if sshKeys == "Not Found" {
		fmt.Println("No keys found: ", githubURL)
		os.Exit(1)
	}
	fmt.Println(sshKeys)

	err = ioutil.WriteFile(authorizedKeysPath, body, 0644)
	errCheck(err, "Error writing to authorized_keys path")
}
