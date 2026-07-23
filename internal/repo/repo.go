package repo

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Repo struct {
	Head   string // The current hash head is pointing to
	Branch string // the current branch
}

var Repo_info *Repo

func Get_repo_info() error {
	head, err := os.ReadFile(".gogit/HEAD")
	if err != nil {
		return fmt.Errorf("could not read .gogit/HEAD")
	}
	head_str := string(head)
	ref, found := strings.CutPrefix(head_str, "ref: ")
	if found {
		head, err = os.ReadFile(".gogit/" + ref)
		if err != nil {
			return fmt.Errorf("could not read head ref")
		}
		head_hash := string(head)
		branch := path.Base(ref)
		Repo_info = &Repo{Head: head_hash, Branch: branch}
	} else {
		return fmt.Errorf("could not find HEAD")
	}
	return nil
}
