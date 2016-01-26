package gist

type Fail struct {
	Message string
}

type Gist struct {
	URL   string
	Files map[string]File
}

type File struct {
	Filename string
	Type     string
	Language string
	RawURL   string
}
