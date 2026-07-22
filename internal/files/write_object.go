package files

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/sonofenoch/go-git-em/internal/tree"
)

func WriteBlob(filename string) (string, error) {
	// open input file & read into bytes
	input, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("could not open %s to compress: %w", filename, err)
	}
	defer input.Close()
	file_bytes, err := io.ReadAll(input)
	if err != nil {
		return "", fmt.Errorf("could not read the input file into bytes")
	}

	// build object content
	var b bytes.Buffer

	fmt.Fprintf(&b, "blob %d\x00", len(file_bytes))
	b.Write(file_bytes)
	b_bytes := b.Bytes()

	hash, err := GenerateHash(b_bytes)
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
	_, err = os.Stat(output_filename)
	if !os.IsNotExist(err) {
		// if the file already exists, we can exit early.
		return "", nil
	} else if err != nil && os.IsExist(err) {
		return "", fmt.Errorf("could not stat %s: %w", filename, err)
	}

	// open output file and create writer
	output, err := os.OpenFile(output_filename, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("could not create and open compressed file %s: %w", output_filename, err)
	}
	defer output.Close()

	w := zlib.NewWriter(output)
	defer w.Close()

	// compress file and flush
	_, err = io.Copy(w, bytes.NewReader(b_bytes))

	return hash, nil
}

func WriteTree(tree *tree.Tree) error {
	var pb bytes.Buffer
	for _, entry := range tree.Entries {
		hash_bytes, err := hex.DecodeString(entry.Hash)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(&pb, "%s ", entry.FileMode)
		pb.WriteString(entry.Filename)
		pb.WriteString("\x00")
		pb.Write(hash_bytes)
	}
	buffer_len := pb.Len()
	var b bytes.Buffer
	fmt.Fprintf(&b, "tree %d\x00", buffer_len)
	b.Write(pb.Bytes())
	hash, err := GenerateHash(b.Bytes())
	fmt.Println(hash)
	if err != nil {
		return err
	}
	// verify output path
	output_dir := fmt.Sprintf("%s/%s", ".gogit/objects", hash[0:2])
	err = CreatePathIfNotExists(output_dir)
	if err != nil {
		return fmt.Errorf("%s does not exist and could not be created: %w", output_dir, err)
	}
	output_filename := output_dir + "/" + hash[2:]

	// open output file and create writer
	output, err := os.OpenFile(output_filename, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("could not create and open compressed file %s: %w", output_filename, err)
	}
	defer output.Close()

	w := zlib.NewWriter(output)
	defer w.Close()

	// compress file and flush
	_, err = io.Copy(w, bytes.NewReader(b.Bytes()))

	return nil

}

func CreatePathIfNotExists(path string) error {
	err := os.MkdirAll(path, 0755)
	return err
}
