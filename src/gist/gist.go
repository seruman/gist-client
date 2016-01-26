package gist

type Fail struct {
	Message string
}

type Gist struct {
	URL   string          `json:"html_url"`
	Files map[string]File `json:"files"`
}

type File struct {
	Filename string
	Type     string
	Language string
	RawURL   string
}
