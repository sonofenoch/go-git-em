package object

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func WriteBlob(blob Blob) error {
	// verify output path
	output_dir := fmt.Sprintf("%s/%s", ".gogit/objects", blob.Hash[0:2])
	err := CreatePathIfNotExists(output_dir)
	if err != nil {
		return fmt.Errorf("%s does not exist and could not be created: %w", blob.Filename, err)
	}
	output_filename := output_dir + "/" + blob.Hash[2:]
	_, err = os.Stat(output_filename)
	if !os.IsNotExist(err) {
		// if the file already exists, we can exit early.
		return nil
	} else if err != nil && os.IsExist(err) {
		return fmt.Errorf("could not stat %s: %w", blob.Filename, err)
	}

	// open input file & read into bytes
	stats, err := os.Stat(blob.Filename)
	if err != nil {
		return fmt.Errorf("could not stat %s, %w", blob.Filename, err)
	}
	input, err := os.Open(blob.Filename)
	if err != nil {
		return fmt.Errorf("could not open %s to compress: %w", blob.Filename, err)
	}
	defer input.Close()

	// open output file and create writer
	output, err := os.OpenFile(output_filename, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("could not create and open compressed file %s: %w", output_filename, err)
	}
	defer output.Close()

	w := zlib.NewWriter(output)
	defer w.Close()

	// compress file and flush
	io.WriteString(w, fmt.Sprintf("blob %d", stats.Size()))
	_, err = io.Copy(w, input)

	return nil
}

func CreatePathIfNotExists(path string) error {
	err := os.MkdirAll(path, 0755)
	return err
}
