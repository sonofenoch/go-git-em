package files

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func GenerateHash(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	hasher := sha1.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", fmt.Errorf("could not copy bytes to hasher: %w", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func GenerateFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("could not open %s to generate hash: %w", path, err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("could not read %s to generate hash: %w", path, err)
	}
	return GenerateHash(bytes)
}
