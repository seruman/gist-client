package Gist

import "io/ioutil"
import "path/filepath"

type File struct {
	Content string `json:"content"`
}

type Snippet struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

func Create(desc string, public bool) Snippet {
	s := Snippet{Description: desc, Public: public, Files: map[string]File{}}
	return s
}

func (s *Snippet) AddFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	d := filepath.Base(filename)
	check(err)
	s.Files[d] = File{Content: string(content)}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
