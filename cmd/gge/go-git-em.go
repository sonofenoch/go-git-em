package main

import (
	"fmt"
	"os"

	"github.com/sonofenoch/go-git-em/internal/git_add"
	"github.com/sonofenoch/go-git-em/internal/git_branch"
	"github.com/sonofenoch/go-git-em/internal/git_commit"
	"github.com/sonofenoch/go-git-em/internal/git_config"
	"github.com/sonofenoch/go-git-em/internal/git_init"
	"github.com/sonofenoch/go-git-em/internal/git_status"
	"github.com/sonofenoch/go-git-em/internal/repo"
)

func printHelp() {
	fmt.Print(`Welcome to enoch's Git implementation in Go!

Usage:
    go-git-em <command> [flags]

Commands:
    init
    add 
    status
    commit
    push

Flags:
    -v, --verbose
`)

}

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(-1)
	}
	err := repo.Get_repo_info()
	if err != nil {
		fmt.Printf("warning! no head or branch set\n")
	}
	err = git_config.GetConfig()
	if err != nil {
		panic(fmt.Errorf("could not read config: %w", err))
	}

	switch os.Args[1] {
	case "init":
		err := git_init.Init("go-git-em")
		if err != nil {
			panic(err)
		}
	case "add":
		if len(os.Args) <= 2 {
			fmt.Printf("nothing specified, nothing added\n")
			os.Exit(1)
		}
		err := git_add.Add(os.Args[2:])
		if err != nil {
			panic(err)
		}
	case "status":
		err := git_status.Status()
		if err != nil {
			if err == git_status.NothingToCommit {
				fmt.Printf("On branch %s%v", repo.Repo_info.Branch, err)
			}
			panic(err)
		}
	case "commit":
		if len(os.Args) < 3 {
			panic("no commit message passed")
		}
		err := git_commit.Commit(os.Args[2])
		if err != nil {
			panic(err)
		}
	case "branch":
		if len(os.Args) == 2 {
			branches, err := git_branch.List_branches()
			if err != nil {
				panic(err)
			}
			for _, branch := range branches {
				fmt.Println(branch)
			}
		} else if len(os.Args) == 3 {
			err := git_branch.Add_branch(os.Args[2])
			if err != nil {
				panic(err)
			}
		} else {
			panic("too many arguments passed to \"gge branch\"")
		}

	case "push":
		fmt.Printf("TODO\n")
	}
}
