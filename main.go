package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	name := "test" // email

	s := "3x8)*&$"
	h := md5.New()
	io.WriteString(h, s+name)
	s = fmt.Sprintf("%X", h.Sum(nil))
	s = s[:4] + "-" + s[4:]
	fmt.Println(s[0:9])
}
