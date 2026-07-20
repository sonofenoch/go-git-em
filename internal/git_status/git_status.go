package git_status

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sonofenoch/go-git-em/internal/index"
)

var NothingToCommit = errors.New("nothing to commit, working tree clean")

func Status() error {
	i, err := index.Read()
	if err != nil {
		if os.IsNotExist(err) {
			return NothingToCommit
		}
		fmt.Println(os.IsNotExist(err))
		return err
	}

	// path: status
	// i.e README.md: modified
	var staged map[string]string = map[string]string{}
	var unstaged map[string]string = map[string]string{}
	var untracked []string = []string{}

	// ignore := files.ReadIgnoreList()
	// TODO: This needs a rework after commits and trees have been implimented
	// TODO: This is because in order to detect things like new and deleted files,
	// TODO: It needs to have the commit history
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.Contains(path, ".git/") || strings.Contains(path, ".gogit/") {
				return nil
			}
			if !index.PathInIndex(i, path) {
				untracked = append(untracked, path)
			} else {
				entry := index.GetEntry(i, path)
				if index.IsChanged(entry, path) {
					unstaged[path] = "modified"
				} else {
					staged[path] = "modified"
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Printf("on branch {BRANCH_NAME}\n")                     // TODO: make global branch
	fmt.Printf("your branch is {?UPTODATE?} with {REMOTE}\n\n") // TODO: make remote

	fmt.Printf("Changes to be committed:\n\n")
	for path, status := range staged {
		fmt.Printf("\t%s:\t%s\n", status, path)
	}

	fmt.Printf("\nChanges not staged for commit:\n\n")
	for path, status := range unstaged {
		fmt.Printf("\t%s:\t%s\n", status, path)
	}

	fmt.Printf("\nUntracked files:\n\n")
	for _, path := range untracked {
		fmt.Printf("\t%s\n", path)
	}
	return nil
}
