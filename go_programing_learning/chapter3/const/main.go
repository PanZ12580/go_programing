package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	SHA256 = iota
	SHA384
	SHA512
)

var s = flag.String("f", "SHA256", "select SHA256 or SHA384 or SHA512")

func main() {
/*	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	differCount := utils.Differ(&s1, &s2)
	fmt.Printf("%x\n%x\nDiffernt Count: %d\n", s1, s2, differCount)*/
	flag.Parse()
	switch *s {
	case "SHA256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(flag.Arg(0))))
	case "SHA384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(flag.Arg(0))))
	case "SHA512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(flag.Arg(0))))
	}
}
