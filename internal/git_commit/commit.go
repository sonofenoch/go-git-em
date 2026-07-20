package git_commit

import (
	"fmt"
	"reflect"

	"github.com/sonofenoch/go-git-em/internal/index"
	"github.com/sonofenoch/go-git-em/internal/tree"
)

func Commit() error {
	tree1 := tree.BuildTree(index.GetIndex())
	tree2 := tree.BuildTree(index.GetIndex())
	fmt.Println(reflect.DeepEqual(tree1, tree2))

	// get current index
	// build tree
	// hash tree
	// compare tree hash to HEAD tree hash
	// if different, write commit to object registry
	// move HEAD
	// move BRANCH

	return nil
}
