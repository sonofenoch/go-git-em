package tree

import (
	"slices"

	"github.com/sonofenoch/go-git-em/internal/index"
)

func BuildTree(i *index.Index) *Tree {
	tree := NewTree()

	for _, entry := range i.Entries {
		tree.Entries = append(tree.Entries, TreeEntry{FileMode: entry.Mode, Filename: entry.Path, Hash: entry.Hash})
	}

	slices.SortFunc(tree.Entries, func(a, b TreeEntry) int {
		if a.Filename < b.Filename {
			return -1
		} else if a.Filename > b.Filename {
			return 1
		}
		return 0
	}) // by default slices does lexicographic sorting on the bytes of a string
	return tree
}
