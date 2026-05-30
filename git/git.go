package git 

import (
	"encoding/json"
	"go-git-em/dag"
	"fmt"
)

// A datastructure containing commit objects. 
// Each commit object points to it's predecessor
// The first commit's previous is nil
type GitTree struct {
	head *dag.Commit
	conf Config
	path string
}

// effectively git init. Sets up directory structure
func newGitTree(path string) (*GitTree, error) {
	var err error	

	err = os.Mkdir(path, 0744) 
	if err != nil {
		return nil, err
	}

	err = os.Mkdir(path + "/staging", 0744) 
	if err != nil {
		return nil, err
	}


	conf, err := ReadConfig()
	if err != nil {
		return nil, err
	}

	gt := GitTree{
		head: nil 
		conf: conf
		path: path
	}	

	return &gt, nil
}


