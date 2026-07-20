package tree

import (
	"github.com/sonofenoch/go-git-em/internal/object"
)

type Tree struct {
	Files      []Blob
	Subfolders []Tree
}
