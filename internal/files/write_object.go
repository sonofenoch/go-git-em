package files

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/sonofenoch/go-git-em/internal/tree"
)

func WriteObject(object_type string, payload []byte) (string, error) {
	var object bytes.Buffer
	fmt.Fprintf(&object, "%s %d\x00", object_type, len(payload))
	object.Write(payload)
	object_bytes := object.Bytes()

	hash, err := GenerateHash(object_bytes)
	if err != nil {
		return "", err
	}

	// verify output path
	output_dir := fmt.Sprintf("%s/%s", ".gogit/objects", hash[0:2])
	err = CreatePathIfNotExists(output_dir)
	if err != nil {
		return "", fmt.Errorf("%s does not exist and could not be created: %w", output_dir, err)
	}
	output_filename := output_dir + "/" + hash[2:]

	// open output file and create writer
	if _, err := os.Stat(output_filename); err == nil {
		return hash, nil
	}
	output, err := os.OpenFile(output_filename, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("could not create and open compressed file %s: %w", output_filename, err)
	}
	defer output.Close()

	w := zlib.NewWriter(output)
	defer w.Close()

	// compress file and flush
	_, err = w.Write(object_bytes)
	if err != nil {
		return "", fmt.Errorf("could not write %s to object registry", object_type)
	}

	return hash, err

}

func WriteBlob(filename string) (string, error) {
	// open input file & read into bytes
	payload, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not read the input file into bytes")
	}
	hash, err := WriteObject("blob", payload)
	return hash, err
}

func WriteTree(tree *tree.Tree) (string, error) {
	var payload bytes.Buffer
	for _, entry := range tree.Entries {
		hash_bytes, err := hex.DecodeString(entry.Hash)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&payload, "%s ", entry.FileMode)
		payload.WriteString(entry.Filename)
		payload.WriteString("\x00")
		payload.Write(hash_bytes)
	}
	payload_bytes := payload.Bytes()

	hash, err := WriteObject("tree", payload_bytes)
	return hash, err
}

func WriteCommit(commit_message string, tree_hash string, parents []string, author string, committer string) (string, error) {
	var payload bytes.Buffer
	fmt.Fprintf(&payload, "tree %s\n", tree_hash)
	for _, parent := range parents {
		fmt.Fprintf(&payload, "parent %s\n", parent)
	}

	fmt.Fprintf(&payload, "author %s\n", author)
	fmt.Fprintf(&payload, "committer %s\n\n", committer)

	fmt.Fprint(&payload, commit_message)

	payload_bytes := payload.Bytes()

	hash, err := WriteObject("commit", payload_bytes)
	return hash, err
}

func CreatePathIfNotExists(path string) error {
	err := os.MkdirAll(path, 0755)
	return err
}
