package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {

	s, err := ReadFileUTF16(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s]\n", s)
	return 0
}

func printHex(bs []byte) {
	for _, b := range bs {
		fmt.Printf("%02X,", b)
	}
	fmt.Println("")
}
