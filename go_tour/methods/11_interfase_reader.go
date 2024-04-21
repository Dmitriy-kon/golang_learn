package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := range len(b) {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'A' + (((b[i] - 'A') + 13) % 26)
		}
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 'a' + (((b[i] - 'a') + 13) % 26)
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	fmt.Printf("%s\n", r)
	io.Copy(os.Stdout, &r)
}
