package git

import ()

// Simple named struct that just has a name and points to a commit in the tree
type Branch struct {
	Name   string
	Commit *Commit
}

func NewBranch(name string, commit *Commit) *Branch {
	return &Branch{name, commit}
}

type Repo struct {
	Name         string
	Head         *Commit // head just points at a commit
	Commit_count int
	Branches     []Branch
	Path         string
	Config       *GitConfig
}

func NewRepo(name string, initial_commit *Commit) *Repo {
	return &Repo{Name: name, Head: initial_commit, Branches: []Branch{{Name: Config.Default_branch_name, Commit: initial_commit}}}
}
