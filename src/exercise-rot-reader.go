package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i, ch := range b {
		if ch >= 'a' && ch <= 'z' {
			b[i] = (ch - 'a' + 13) % 26 + 'a'
		}
		if ch >= 'A' && ch <= 'Z' {
			b[i] = (ch - 'A' + 13) % 26 + 'A'
		}
	}
	return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
