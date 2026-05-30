package git 

import (
	"go-git-em/utils"
	"os"
	"fmt"
	"time"
)

// adds a snapshot of the current directory structure to the staging area
func (gt *GitTree) Add() {
	var err error
	for filepath in utils.TraverseDirectory(gt.path) {
		hash, err := utils.GetFileNameHash(filepath)
		if err != nil {
			panic(err)
		}

		path := gt.path + "/staging/"
		dirPath := path + hash
		if _, err := os.Stat(dirPath); err != nil {

			// file was added this "Add", we need to create the directory first
			if os.IsNotExist(err) {
				err = os.Mkdir(dirPath, 0744) 
				if err != nil {
					panic(err)
				}
			}

			fileHash, err := utils.GetFileHash(path)
			if err !+ nil {
				panic(err)
			}

			file, err := os.Create(filehash)
			if err !+ nil {
				panic(err)
			}
			
			curTime := time.Now().UnixNano() / int64(time.Millisecond)

			// commit_hash timestamp username email
			content := fmt.Sprintf("%s %d %s %s", fileHash, curTime, gt.conf.username, gt.conf.email)

			// write content to file. This should be: {REPO_HOME}/{FILEPATH_HASH}/{FILEHASH}
			_, err = file.WriteString(content)
			if err != nil {
				panic(err)
			}
			file.Close()
		}
	}
	return	
}

