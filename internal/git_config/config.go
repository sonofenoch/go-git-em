package git_config

import "os"
import "io"
import "encoding/json"
import "fmt"

// IMPORTANT!
// Git uses ini config files. Thats annoying and I already have to recreate its behavior.
// I am using JSON as it is easier and I frankly can't be bothered. Might revisit but idk

// Ignoring 99% of core features as they are overkill for this project
type Core struct {
	Editor string `json:"editor"` // ignoring for now, but core is a placeholder
}

// We need these for commit messages etc
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Init struct {
	DefaultBranch string `json:"defaultBranch"`
}

type Remote struct {
	Name  string `json:"name"`
	Url   string `json:"url"`   // ignoring for now
	Fetch string `json:"fetch"` // ignoring for now
}

type Branch struct {
	Name   string `json:"name"`
	Remote string `json:"remote"`
	Merge  string `json:"merge"`
}

type GitConfig struct {
	Core   Core     `json:"core"`
	User   User     `json:"user"`
	Init   Init     `json:"init"`
	Remote []Remote `json:"remote"`
	Branch []Branch `json:"branch"`
}

var Config *GitConfig = nil

func read() (*GitConfig, error) {
	var err error
	config, err := os.Open(".gogit/config")
	if err != nil {
		return nil, fmt.Errorf("could not open .gogit/config: %w", err)
	}
	defer config.Close()
	bytes, err := io.ReadAll(config)

	if err != nil {
		return nil, fmt.Errorf("could not read contents of .gogit/config: %w", err)
	}

	var gitconfig GitConfig
	err = json.Unmarshal(bytes, &gitconfig)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal .gogit/config: %w", err)
	}

	return &gitconfig, nil
}

func GetConfig() error {
	var err error
	if Config == nil {
		Config, err = read()
		if err != nil {
			return err
		}
	}
	return nil
}
