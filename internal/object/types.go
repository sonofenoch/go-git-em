package object

type Object struct {
	Type string // blob | commit | tree | etc
	Hash string
}

type Blob struct {
	Filename string
	Object
}

type Commit struct {
	CommitMessage string
	Object
}
