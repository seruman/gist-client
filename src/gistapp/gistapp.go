package main

import (
	"Gist"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://api.github.com/gists"

func main() {
	aToken := os.Getenv("GH_TOKEN")
	b := Gist.Create("aciklama", false)

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

	if resp.StatusCode != 201 {
		fail := Gist.Fail{}
		content, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(content, &fail)

		fmt.Fprint(os.Stderr, "Github: ", fail.Message, "\n")
	}

	// _, err := ioutil.ReadAll(resp.Body) // just in case
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// 	os.Exit(1)
	// }

}
