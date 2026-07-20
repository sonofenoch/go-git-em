package main

import (
	"fmt"
	"os"

	"github.com/sonofenoch/go-git-em/internal/git_add"
	"github.com/sonofenoch/go-git-em/internal/git_commit"
	"github.com/sonofenoch/go-git-em/internal/git_init"
	"github.com/sonofenoch/go-git-em/internal/git_status"
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
	branch := "master"

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
				fmt.Printf("On branch %s%v", branch, err)
			}
			panic(err)
		}
	case "commit":
		err := git_commit.Commit()
		if err != nil {
			panic(err)
		}
	case "push":
		fmt.Printf("TODO\n")
	}
}
