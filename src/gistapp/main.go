package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gist"
	"io/ioutil"
	"net/http"
	"os"
	"snippet"
	"urlshortener"
)

const url = "https://api.github.com/gists"

func main() {
	aToken := os.Getenv("GH_TOKEN") // Github Access Token
	//gKey := os.Getenv("GOOGL_API_KEY") // API key for Google Url Shortener
	b := snippet.Create("aciklama", false)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "error: No input file(s)\n")
		os.Exit(64)
	}

	for index := 0; index < len(args); index++ {
		err := b.AddFile(args[index])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(66)
		}
	}

	parsed, _ := json.Marshal(b)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(parsed))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+aToken)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var content []byte

	if resp.StatusCode != 201 {
		fail := gist.Fail{}
		content, _ = ioutil.ReadAll(resp.Body)
		json.Unmarshal(content, &fail)
		fmt.Fprint(os.Stderr, "Github: ", fail.Message, "\n")
		os.Exit(1)
	}

	success := gist.Gist{}
	content, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(content, &success)

	fmt.Println(string(urlshortener.Shorten(success.URL)))

	// data, err := ioutil.ReadAll(resp.Body) // just in case
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println(string(data))

}
