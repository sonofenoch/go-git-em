package git_commit

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/sonofenoch/go-git-em/internal/files"
	"github.com/sonofenoch/go-git-em/internal/git_branch"
	"github.com/sonofenoch/go-git-em/internal/git_config"
	"github.com/sonofenoch/go-git-em/internal/index"
	"github.com/sonofenoch/go-git-em/internal/repo"
	"github.com/sonofenoch/go-git-em/internal/tree"
)

func Commit(commit_message string) error {
	tree := tree.BuildTree(index.GetIndex())

	tree_hash, err := files.WriteTree(tree)
	if err != nil {
		return err
	}

	var branch string
	if repo.Repo_info != nil {
		if tree_hash != repo.Repo_info.Head {
			fmt.Println(tree_hash)
			fmt.Println(repo.Repo_info.Head)
			branch = repo.Repo_info.Branch
		} else {
			return nil
		}
	} else {
		branch = git_config.Config.Init.DefaultBranch
	}
	current_time := time.Now()
	_, offset := current_time.Zone()
	author := fmt.Sprintf("%s <%s> %d %d", git_config.Config.User.Name, git_config.Config.User.Email, current_time.Unix(), offset)

	// create commit. Parent nil, author and committer the same.
	// TODO: There a few cases when the author and the committer would not be the same... might not care for this project
	commit_hash, err := files.WriteCommit(commit_message, tree_hash, []string{}, author, author)
	if err != nil {
		return err
	}
	fmt.Println(branch)
	fmt.Println(commit_hash)

	err = os.WriteFile(".gogit/refs/heads/"+branch, []byte(commit_hash), 0644)
	if err != nil {
		return fmt.Errorf("could not change head")
	}

	err = repo.Get_repo_info()

	if err == nil {
		branches, err := git_branch.List_branches()
		if err != nil {
			return err
		}
		if !slices.Contains(branches, branch) {
			git_branch.Add_branch(branch)
		}
	} else {
		return fmt.Errorf("could not refresh repo info")
	}
	// get current index
	// build tree
	// hash tree
	// compare tree hash to HEAD tree hash
	// if different, write commit to object registry
	// move HEAD
	// move BRANCH

	return nil
}
