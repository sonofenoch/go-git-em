package git 


// gets the commit hash (aggregate hash of all files)
func (gt *GitTree) getCommitHash( string) (string, error) {
	return "", nil
}

// commit function. Adds new commit, resets head to point at new commit 
func (gt *GitTree) Commit(message: string) *dag.Commit {
	hash := getHash()
	config := readConfig()
	cm := commit{
		hash: hash,
		username: gt.Config.username, 
		email: gt.email,
		message: message,
		previous: gt.head 
	}

	gt.head = cm 
	return &cm
}


