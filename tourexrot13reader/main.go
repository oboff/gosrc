package main

import (
	"unicode"
	//"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	for i := 0; i < len(b); i++ {
		switch {
		case unicode.IsUpper(rune(b[i])):
			if b[i]+13 > 'Z' {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		case unicode.IsLower(rune(b[i])):
			if b[i]+13 > 'z' {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
