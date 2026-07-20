package git_add

import (
	"github.com/sonofenoch/go-git-em/internal/files"
	"github.com/sonofenoch/go-git-em/internal/index"
	"github.com/sonofenoch/go-git-em/internal/object"
)

func Add(paths []string) error {
	i := index.GetIndex()
	for _, path := range paths {
		file_hash, err := files.GenerateFileHash(path)
		if err != nil {
			return nil
		}
		object.WriteBlob(object.Blob{Filename: path, Object: object.Object{Type: "blob", Hash: file_hash}})
		index.AddFile(i, path, file_hash)
	}

	return index.Write(i)
}
