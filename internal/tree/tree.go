package tree

import (
	"github.com/sonofenoch/go-git-em/internal/object"
)

type Blob struct {
	Filename string
	Object   object.Object
}

type Tree struct {
	Files     []Blob
	Subfoldes []Tree
}
