package git_commit

import (
	"fmt"

	"github.com/sonofenoch/go-git-em/internal/files"
	"github.com/sonofenoch/go-git-em/internal/index"
	"github.com/sonofenoch/go-git-em/internal/tree"
)

func Commit() error {
	tree := tree.BuildTree(index.GetIndex())

	fmt.Println(tree)
	files.WriteTree(tree)

	// get current index
	// build tree
	// hash tree
	// compare tree hash to HEAD tree hash
	// if different, write commit to object registry
	// move HEAD
	// move BRANCH

	return nil
}
