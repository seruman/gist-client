package Gist

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

func (s *Snippet) AddFile(filename string, content string) {
	s.Files[filename] = File{Content: content}
}
