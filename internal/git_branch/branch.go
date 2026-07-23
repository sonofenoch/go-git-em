package git_branch

import (
	"fmt"
	"os"
	"path"

	"github.com/sonofenoch/go-git-em/internal/git_config"
	"github.com/sonofenoch/go-git-em/internal/repo"
)

func List_branches() ([]string, error) {
	dir_entries, err := os.ReadDir(".gogit/refs/heads")
	if err != nil {
		return []string{}, fmt.Errorf("could not read .gogit/refs/heads")
	}

	var branches []string = []string{}
	for _, dir_ent := range dir_entries {
		if dir_ent.IsDir() {
			continue
		}
		branches = append(branches, path.Base(dir_ent.Name()))
	}

	return branches, err
}

func Add_branch(branch_name string) error {
	if repo.Repo_info != nil {
		head_hash := repo.Repo_info.Head
		err := os.WriteFile(".gogit/refs/heads/"+branch_name, []byte(head_hash), 0655)
		if err != nil {
			return fmt.Errorf("could not write new branch file")
		}
		return err
	} else {
		return fmt.Errorf("fatal: not a valid object name: '%s'", git_config.Config.Init.DefaultBranch)
	}
}
