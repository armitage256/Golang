package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/md5"
	"fmt"
)

func main() {

	password := "hello world"

	SHA1   := fmt.Sprintf( "%x", sha1.Sum( []byte( password ) ) )
	SHA256 := fmt.Sprintf( "%x", sha256.Sum256( []byte( password ) ) )
	MD5    := fmt.Sprintf( "%x", md5.Sum( []byte( password ) ) )

	fmt.Println("SHA-1: ", SHA1)
	fmt.Println("SHA-256: ", SHA256)
	fmt.Println("MD5: ", MD5)

}