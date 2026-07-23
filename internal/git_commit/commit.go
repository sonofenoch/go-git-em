package git_commit

import (
	"fmt"
	"time"

	"github.com/sonofenoch/go-git-em/internal/files"
	"github.com/sonofenoch/go-git-em/internal/git_config"
	"github.com/sonofenoch/go-git-em/internal/index"
	"github.com/sonofenoch/go-git-em/internal/repo"
	"github.com/sonofenoch/go-git-em/internal/tree"
)

func Commit(commit_message string) error {
	tree := tree.BuildTree(index.GetIndex())

	fmt.Println(tree)
	tree_hash, err := files.WriteTree(tree)
	if err != nil {
		return err
	}

	if repo.Repo_info != nil {
		if tree_hash != repo.Repo_info.Head {
			fmt.Println(tree_hash)
			fmt.Println(repo.Repo_info.Head)
		}
	} else {
		branch := git_config.Config.Init.DefaultBranch
		current_time := time.Now()
		_, offset := current_time.Zone()
		author := fmt.Sprintf("%s <%s> %d %d", git_config.Config.User.Name, git_config.Config.User.Email, current_time.Unix(), offset)

		// create commit. Parent nil, author and committer the same.
		commit_hash, err := files.WriteCommit(commit_message, tree_hash, nil, author, author)
		if err != nil {
			return err
		}
		fmt.Println(branch)
		fmt.Println(commit_hash)
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
