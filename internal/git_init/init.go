package git_init

import ()

func Init(repo_name string, branch string) error {
	err := Create_gogit(branch)
	if err != nil {
		return err
	}
	return nil
}
