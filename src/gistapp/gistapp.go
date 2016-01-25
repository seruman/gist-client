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
const username = "seruman"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	aToken := os.Getenv("token")

	argsWithProg := os.Args
	fmt.Println(argsWithProg)
	//data, err := ioutil.ReadFile(argsWithProg[1])

	b := Gist.Create("aciklama", false)
	b.AddFile(argsWithProg[1])
	//b.AddFile(argsWithProg[2])

	parsed, _ := json.Marshal(b)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(parsed))

	check(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+aToken)

	resp, err := client.Do(req)

	check(err)

	contents, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Println(string(contents))

}
