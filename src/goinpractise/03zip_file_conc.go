package goinpractise

import (
	"os"
	"io"
	"compress/gzip"
	"fmt"
	"goreceipes/concurrency/syncutils"
	"time"
)

// ZipFileConcMain is the main program. It compresses all the files specified
// on the command line one by one. The program is non performant as it
// does not utilize the CPU cores effectively.

func ZipFileConcMain() {
	// check if appropriate arguments are passed
	if len(os.Args) <= 2 {
		fmt.Println("Usage: go run 03zip_file.go <file1> <file2> ...")
		os.Exit(1)
	}

	var wg  = syncutils.Wg
	var startTime = time.Now()

	// read the files from the os.Args[1:]
	var count int = 0
	for _, file := range os.Args[1:] {
		count++
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)

	}

	wg.Wait()
	var elapsedTime = time.Since(startTime)
	fmt.Printf("Total time taken to compress %d files = %f seconds\n", count, (elapsedTime/1000))
}

func compress(filename string) error {
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
