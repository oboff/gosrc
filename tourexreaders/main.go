//Stimulus: Implement a Reader type that emits an
//			infinite stream of the ASCII character 'A'

package main

import (
	//"io"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	for i, _ := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
