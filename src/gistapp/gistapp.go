package main

import (
	"Gist"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg)
	//data, err := ioutil.ReadFile(argsWithProg[1])

	b := Gist.Create("aciklama", false)
	b.AddFile("testfile.txt", "content")
	b.AddFile("testfile2.txt", "content2")

	parsed, _ := json.Marshal(b)
	fmt.Println(string(parsed))
}
