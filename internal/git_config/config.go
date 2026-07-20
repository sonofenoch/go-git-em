package config

import "os"
import "io"
import "encoding/json"
import "fmt"

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type GitConfig struct {
	User                User   `json:"user"`
	Default_branch_name string `json:"default_branch_name"`
}

func Read() (*GitConfig, error) {
	var err error
	config, err := os.Open(".gogitconfig")
	if err != nil {
		return nil, fmt.Errorf("could not open .gogitconfig")
	}
	defer config.Close()
	bytes, err := io.ReadAll(config)

	if err != nil {
		return nil, fmt.Errorf("could not read contents of .gogitconfig")
	}

	var gitconfig GitConfig
	err = json.Unmarshal(bytes, &gitconfig)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal .gogitconfig")
	}

	return &gitconfig, nil
}
