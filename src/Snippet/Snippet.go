package Snippet

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

type Fail struct {
	Message string
}

type Success struct {
}

func Create(desc string, public bool) Snippet {
	s := Snippet{Description: desc, Public: public, Files: map[string]File{}}
	return s
}

func (s *Snippet) AddFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	d := filepath.Base(filename)
	s.Files[d] = File{Content: string(content)}

	return nil
}
