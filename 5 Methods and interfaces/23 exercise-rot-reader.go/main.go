package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Byte(c byte) byte {
    switch {
    case c >= 'A' && c <= 'Z':
        return 'A' + (c-'A'+13)%26
    case c >= 'a' && c <= 'z':
        return 'a' + (c-'a'+13)%26
    default:
        return c
    }
}

func (ro *rot13Reader) Read(p []byte) (int, error){
	n, err := ro.r.Read(p) 
	for i := 0; i < n; i++ {
		p[i] = rot13Byte(p[i]) // aplicamos ROT13 sólo a los bytes leídos
	}
	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
