package git_init

import (
	"github.com/sonofenoch/go-git-em/internal/git_config"
)

func Init(repo_name string) error {
	config, err := git_config.GetConfig()
	if err != nil {
		return err
	}
	err = Create_gogit(config.Init.DefaultBranch)
	if err != nil {
		return err
	}
	return nil
}
