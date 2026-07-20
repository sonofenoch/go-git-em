package git_init

import (
	"fmt"
	"io"
	"os"
)

func Create_gogit(branch string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current dir: %w", err)
	}

	err = os.Mkdir(".gogit", 0755)
	if os.IsExist(err) {
		fmt.Printf("reinitializing existing repository in %s/.gogit\n", wd)
	}
	err = os.Chdir(".gogit")
	if err != nil {
		return fmt.Errorf("could not cd into .gogit")
	}

	file_contents := map[string]string{
		"HEAD":        fmt.Sprintf("ref: refs/heads/%s", branch),
		"description": "Unnamed repository; edit this file 'description' to name the repository.",
	}

	gogit_dirs := []string{"hooks", "info", "objects", "objects/info", "objects/pack", "refs", "refs/heads", "refs/tags"}

	for _, dir := range gogit_dirs {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("could not create %s: %w", dir, err)
		}
	}

	for filename, content := range file_contents {
		err := init_file(filename, content)
		if err != nil {
			return fmt.Errorf("could not create %s: %w", filename, err)
		}
	}

	return nil
}

func init_file(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("could not open %s", filename)
	}
	_, err = io.WriteString(file, content)
	if err != nil {
		return fmt.Errorf("could not initialize %s", filename)
	}
	return nil
}
