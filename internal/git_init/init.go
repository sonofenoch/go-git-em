package git_init

import (
	"github.com/sonofenoch/go-git-em/internal/git_config"
)

func Init(repo_name string) error {
	err := Create_gogit(git_config.Config.Init.DefaultBranch)
	if err != nil {
		return err
	}
	return nil
}
