package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

func main() {
	name := "email" // email

	hash := md5.New()
	const s = "3x8)*&$"

	_, err := hash.Write([]byte(s + name))
	if err != nil {
		log.Fatalln(err)
	}

	key := hash.Sum(nil)
	fmt.Printf("%X-%X", key[:2], key[2:4])
}
