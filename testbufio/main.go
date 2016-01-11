package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

// Process our args
func main() {
	if len(os.Args) < 3 {
		printUsage()
	} else {
		bufferFiles(os.Args[1], os.Args[2])
	}
}

// Print usage menu
func printUsage() {
	log.Println("Usage:")
	log.Println("bufio inputFile outputFile")
}

// borrowed from SO article.
// http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file
// Why fix what ain't broke?
func bufferFiles(input, output string) {
	// open input file
	fi, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}

	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}
}
