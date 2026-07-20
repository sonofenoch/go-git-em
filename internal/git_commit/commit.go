package git

import (
	// "bytes"
	// "io"
	// "os"
	// "strings"
	"fmt"
	"time"
	// "github.com/sonofenoch/go-git-em/git/files"
)

// Struct for a commit in the git tree
type Commit struct {
	Tree      string
	Hash      string
	Parents   []*Commit // we can have multiple parents, but only one "child"
	Message   string
	Author    string
	Committer string
}

func NewCommit(tree string, hash string, parents []*Commit, message string, author string, committer string) *Commit {
	now := time.Now()
	epoch_time := now.UTC().Unix()
	_, time_offset := now.Local().Zone()
	return &Commit{Tree: tree, Hash: hash, Parents: parents, Message: message, Author: fmt.Sprintf("%s %d %d", author, epoch_time, int(time_offset/3600)), Committer: committer}
}

func (r *Repo) WriteCommitObjects(object_dir string, tree string, parents []*Commit, message string, author string, committer string) error {
	return nil
	// var commit_object strings.Builder
	// fmt.Fprintf(&commit_object, "commit %d\000tree %s\n", r.Commit_count, tree)
	// for i := range len(parents) {
	// 	fmt.Fprintf(&commit_object, "parent %s\n", parents[i].Hash)
	// }
	//
	// fmt.Fprintf(&commit_object, "author %s\n", author)
	// fmt.Fprintf(&commit_object, "committer %s\n\n", committer)
	// fmt.Fprintf(&commit_object, "%s", message)
	//
	// commit_bytes := []byte(commit_object.String())
	// hash, err := files.GenerateHash(commit_bytes)
	// if err != nil {
	// 	return fmt.Errorf("could not generate commit hash: %w", err)
	// }
	//
	// file, err := os.OpenFile(fmt.Sprintf("%s/%s/%s", object_dir, hash[0:2], hash[2:]), os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	// if err != nil {
	// 	return fmt.Errorf("could not create and open commit object file")
	// }
	// defer file.Close()
	//
	// _, err = io.Copy(file, bytes.NewReader(commit_bytes))
	// if err != nil {
	// 	return fmt.Errorf("could not write commit object: %w", err)
	// }
	//
	// return nil
}
