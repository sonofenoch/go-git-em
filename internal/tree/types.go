package tree

type Tree struct {
	Files      map[string]string
	Subfolders map[string]*Tree
}

func NewTree() *Tree {
	return &Tree{Files: map[string]string{}, Subfolders: map[string]*Tree{}}
}
