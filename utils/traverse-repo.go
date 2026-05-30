package utils

import (
	"fmt"
	"os"
	"io"
	"iter"
)


// traverses directory to add files to staging
func TraverseDirectory(dir string) iter.Seq[string] {
	var err error
	return func(yield func(string) bool) {
		entries, _ := os.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, entry := range entries {
			// recursively call if entry is a directory
			if entry.IsDir() {
				err = traverseDirectory(entry.Name())
				if err != nil {
					return fmt.Errorf("Could not recursively traverse directory") 
				}
			} else if entry.IsFile() {
				if !yield(entry.Name())	{
					return
				}
			}
		}
	}
}

