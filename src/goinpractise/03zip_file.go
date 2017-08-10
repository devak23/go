package goinpractise

import (
	"os"
	"io"
	"compress/gzip"
	"fmt"
)

// ZipFileMain is the main program. It compresses all the files specified
// on the command line one by one. The program is non performant as it
// does not utilize the CPU cores effectively.
func ZipFileMain() {
	// check if appropriate arguments are passed
	if len(os.Args) <= 2 {
		fmt.Println("Usage: go run 03zip_file.go <file1> <file2> ...")
		os.Exit(1)
	}

	// read the files from the os.Args[1:]
	for _, file := range os.Args[1:] {
		compressFile(file)
	}
}

func compressFile(filename string) error {
	// open the input file name
	in, err := os.Open(filename)
	// return if there is an error
	if err != nil {
		return err
	}
	// if not, remember to close the input file
	defer in.Close()

	// open the output file
	out, err := os.Create(filename + ".gz")
	// if there is an error return it
	if err != nil {
		return err
	}
	// else remember to close the output file
	defer out.Close()

	// wrap the output file into a GZIP output file
	gzout := gzip.NewWriter(out)
	defer gzout.Close()

	// use the io.Copy(dest, source) to copy the contents of the source
	// into the destination
	_, err = io.Copy(gzout, in)

	return err
}
