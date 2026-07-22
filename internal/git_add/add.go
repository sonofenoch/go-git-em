package git_add

import (
	"github.com/sonofenoch/go-git-em/internal/files"
	"github.com/sonofenoch/go-git-em/internal/index"
)

func Add(paths []string) error {
	i := index.GetIndex()
	for _, path := range paths {
		file_hash, err := files.WriteBlob(path)
		if err != nil {
			return err
		}
		index.AddFile(i, path, file_hash)
	}

	return index.Write(i)
}
