package dag

import (
	"fmt"
	"ioutil"
)

type Commit struct {
	hash string 
	username string 
	email string 
	timestamp int64 
	message string
	fileTree []File
	previous *commit
}

type File struct {
	directory string
	filename string
	hash string 
	lastModifiedCommit *Commit 
}

func newFile(path) {
	file = File{path: path, hash: getFileHash(path)}
}


