package tree

import "github.com/sonofenoch/go-git-em/internal/object"

type TreeEntry struct {
	FileMode string
	Filename string
	Obj      object.Object
}
type Tree struct {
	Entries []TreeEntry
}

func NewTree() *Tree {
	return &Tree{Entries: []TreeEntry{}}
}
