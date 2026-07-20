package tree

import (
	"strings"

	"github.com/sonofenoch/go-git-em/internal/index"
)

func BuildTree(i *index.Index) *Tree {
	tree := NewTree()

	for _, entry := range i.Entries {
		AddFile(tree, entry.Path, entry.Hash)
	}

	return tree
}

func AddSubFolder(tree *Tree, path, hash string) {
	split := strings.SplitN(path, "/", 2)
	dir, filename := split[0], split[1]
	if _, ok := tree.Subfolders[dir]; !ok {
		tree.Subfolders[dir] = NewTree()
	}
	AddFile(tree.Subfolders[dir], filename, hash)
}

func AddFile(tree *Tree, path, hash string) {
	if strings.Contains(path, "/") {
		AddSubFolder(tree, path, hash)
	} else {
		tree.Files[path] = hash
	}
}
