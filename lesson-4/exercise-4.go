package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (read rot13Reader) Read(b []byte) (int, error) {
	n, err := read.r.Read(b)
	for i, _ := range b {
		if b[i] >= 'A' && b[i] <= 'M' {
			b[i] += 13
		} else if b[i] >= 'N' && b[i] <= 'Z' {
			b[i] -= 13
		} else if b[i] >= 'a' && b[i] <= 'm' {
			b[i] += 13
		} else if b[i] >= 'n' && b[i] <= 'z' {
			b[i] -= 13
		}
	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
