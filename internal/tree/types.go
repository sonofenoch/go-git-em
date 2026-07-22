package tree

type TreeEntry struct {
	FileMode string
	Filename string
	Hash     string
}
type Tree struct {
	Entries []TreeEntry
}

func NewTree() *Tree {
	return &Tree{Entries: []TreeEntry{}}
}
